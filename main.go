package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/kardianos/osext"

	"github.com/Clever/workflow-manager/executor"
	"github.com/Clever/workflow-manager/executor/sfncache"
	"github.com/Clever/workflow-manager/gen-go/server"
	dynamodbstore "github.com/Clever/workflow-manager/store/dynamodb"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

// Config contains the configuration for the workflow-manager app
type Config struct {
	DynamoPrefixStateResources      string
	DynamoPrefixWorkflowDefinitions string
	DynamoPrefixWorkflows           string
	DynamoRegion                    string
	SFNRegion                       string
	SFNAccountID                    string
	SFNRoleARN                      string
}

func setupRouting() {
	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	err = logger.SetGlobalRouting(path.Join(dir, "kvconfig.yml"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	addr := flag.String("addr", ":8080", "Address to listen at")
	flag.Parse()

	c := loadConfig()
	setupRouting()

	svc := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(c.DynamoRegion)},
	})))
	db := dynamodbstore.New(svc, dynamodbstore.TableConfig{
		PrefixStateResources:      c.DynamoPrefixStateResources,
		PrefixWorkflowDefinitions: c.DynamoPrefixWorkflowDefinitions,
		PrefixWorkflows:           c.DynamoPrefixWorkflows,
	})
	sfnapi := sfn.New(session.New(), aws.NewConfig().WithRegion(c.SFNRegion))
	cachedSFNAPI, err := sfncache.New(sfnapi)
	if err != nil {
		log.Fatal(err)
	}
	wfmSFN := executor.NewSFNWorkflowManager(cachedSFNAPI, db, c.SFNRoleARN, c.SFNRegion, c.SFNAccountID)
	h := Handler{
		store:   db,
		manager: wfmSFN,
	}
	s := server.New(h, *addr)

	go executor.PollForPendingWorkflowsAndUpdateStore(context.Background(), wfmSFN, db)

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}

	log.Println("workflow-manager exited without error")
}

func awsSession(c Config) *session.Session {
	options := session.Options{
		Config:            aws.Config{Region: aws.String("us-east-1")},
		SharedConfigState: session.SharedConfigEnable,
	}

	return session.Must(session.NewSessionWithOptions(options))
}

func loadConfig() Config {
	return Config{
		DynamoPrefixStateResources: getEnvVarOrDefault(
			"AWS_DYNAMO_PREFIX_STATE_RESOURCES",
			"workflow-manager-test",
		),
		DynamoPrefixWorkflowDefinitions: getEnvVarOrDefault(
			"AWS_DYNAMO_PREFIX_WORKFLOW_DEFINITIONS",
			"workflow-manager-test",
		),
		DynamoPrefixWorkflows: getEnvVarOrDefault(
			"AWS_DYNAMO_PREFIX_WORKFLOWS",
			"workflow-manager-test",
		),
		DynamoRegion: os.Getenv("AWS_DYNAMO_REGION"),
		SFNRegion:    os.Getenv("AWS_SFN_REGION"),
		SFNAccountID: os.Getenv("AWS_SFN_ACCOUNT_ID"),
		SFNRoleARN:   os.Getenv("AWS_SFN_ROLE_ARN"),
	}
}

func getEnvVarOrDefault(envVarName, defaultIfEmpty string) string {
	value := os.Getenv(envVarName)
	if value == "" {
		value = defaultIfEmpty
	}

	return value
}
