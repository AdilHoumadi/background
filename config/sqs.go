package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const access = ""
const secret = ""
const region = ""

func GetSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				access,
				secret,
				"",
			),
		},
	}))
}

func GetSQSClient() *sqs.SQS {
	awsSession := GetSession()
	return sqs.New(awsSession)
}

func GetQueueUrl(queueName string) string {
	svc := GetSQSClient()
	params := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}
	log.Println(params)
	res, err := svc.GetQueueUrl(params)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return *res.QueueUrl
}
