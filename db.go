package sample_circleci

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	db *dynamodb.DynamoDB

	userTable string
)

func init() {
	dbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	region := os.Getenv("AWS_REGION")

	userTable = os.Getenv("DYNAMO_TABLE_USER")
	if userTable == "" {
		log.Fatal(`env variable "DYNAMO_TABLE_USER" is required`)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(dbEndpoint),
		Region:   aws.String(region),
	}))
	db = dynamodb.New(sess)
}
