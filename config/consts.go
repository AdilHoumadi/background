package config

import (
	"context"

	"github.com/remind101/mq-go"
)

const (
	ConsoleExec = "console"
	Demo        = "demo"

	// Env variable names
	Env     = "ENV"
	Verbose = "VERBOSE"

	// flags
	EnvFlag     = "environment"
	VerboseFlag = "verbose"

	// shortcut
	EnvShortcut     = "e"
	VerboseShortcut = "v"

	// flag description
	EnvDescription      = "Environment, you can set via ENV"
	UnSecureDescription = "Show or hide log data on the stdout"

	// Emails
	Email      = "email"
	EmailQueue = "email"
	EmailShort = "consume sqs queue email"
)

func Listen(handler func(m *mq.Message) error, queueUrl string) {
	done := make(chan struct{})
	processor := mq.HandlerFunc(handler)
	client := mq.WithClient(GetSQSClient())
	s := mq.NewServer(GetQueueUrl(queueUrl), processor, client)
	s.Start()
	defer s.Shutdown(context.Background())
	<-done
}
