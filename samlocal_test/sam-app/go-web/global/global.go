package global

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

var (
	DB *dynamo.DB
)

func InitDynamoDB() {
	os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
	os.Setenv("AWS_SESSION_TOKEN", "dummy3")
	creds := credentials.NewEnvCredentials()
	creds.Get()
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: creds,
	}))
	DB = dynamo.New(sess)
	if DB == nil {
		panic("connect db wrong")
	}
	fmt.Println("InitDynamoDB...")
	fmt.Println("Tables in DynamoDB:")
	fmt.Println(DB.ListTables().All())
}
