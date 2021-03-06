package dynamodb

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/Clever/workflow-manager/gen-go/server/db"
	"github.com/Clever/workflow-manager/gen-go/server/db/tests"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestDynamoDBStore(t *testing.T) {
	// spin up dynamodb local
	testCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// restart test db if it gets killed before tests are completed.
	go func(doneC <-chan struct{}, t *testing.T) {
		cmd := exec.CommandContext(testCtx, "./dynamodb-local.sh")
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		for {
			select {
			case <-doneC:
				return
			default:
				if err := cmd.Wait(); err != nil {
					fmt.Printf("Test DB crashed: %v\n", err.Error())
					cmd = exec.CommandContext(testCtx, "./dynamodb-local.sh")
					if err := cmd.Start(); err != nil && t != nil {
						if err.Error() != "context canceled" {
							t.Fatal(err)
						}
					}
				}
			}
		}
	}(testCtx.Done(), t)

	// loop for 60s trying to establish a connection
	connected := false
	for start := time.Now(); start.Before(start.Add(60 * time.Second)); time.Sleep(1 * time.Second) {
		if c, err := net.Dial("tcp", "localhost:8002"); err == nil {
			c.Close()
			connected = true
			break
		}
	}
	if connected == false {
		t.Fatal("failed to connect within 60 seconds")
	}

	dynamoDBAPI := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String("doesntmatter"),
			Endpoint:    aws.String("http://localhost:8002" /* default dynamodb-local port */),
			Credentials: credentials.NewStaticCredentials("id", "secret", "token"),
		},
	})))

	tests.RunDBTests(t, func() db.Interface {
		prefix := "automated-testing"
		listTablesOutput, err := dynamoDBAPI.ListTablesWithContext(testCtx, &dynamodb.ListTablesInput{})
		if err != nil {
			t.Fatal(err)
		}
		for _, tableName := range listTablesOutput.TableNames {
			if strings.HasPrefix(*tableName, prefix) {
				dynamoDBAPI.DeleteTableWithContext(testCtx, &dynamodb.DeleteTableInput{
					TableName: tableName,
				})
			}
		}
		d, err := New(Config{
			DynamoDBAPI:               dynamoDBAPI,
			DefaultPrefix:             prefix,
			DefaultReadCapacityUnits:  10,
			DefaultWriteCapacityUnits: 10,
		})
		if err != nil {
			t.Fatal(err)
		}
		if err := d.CreateTables(testCtx); err != nil {
			t.Fatal(err)
		}
		return d
	})
}
