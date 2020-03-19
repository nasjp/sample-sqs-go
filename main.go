package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/nasjp/sample-sqs-go/services"
)

const (
	AWS_REGION            = "ap-northeast-1"
	MAIN_QUEUE_NAME       = "test-main-queue"
	AWS_ACCESS_KEY_ID     = ""
	AWS_SECRET_ACCESS_KEY = ""
	END_POINT             = ""
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
	})
	if err != nil {
		return err
	}

	q := &services.Queue{
		Client: sqs.New(sess),
		URL:    END_POINT,
	}

	msgs, err := q.GetMessages(20)
	if err != nil {
		return err
	}

	fmt.Println("Messages:")
	for _, msg := range msgs {
		fmt.Printf("%s> %s: %s\n", msg.From, msg.To, msg.Msg)
	}

	return nil
}
