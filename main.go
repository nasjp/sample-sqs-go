package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/nasjp/sample-sqs-go/services"
)

var sess = session.Must(session.NewSession())

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Queue URL required.")
		os.Exit(1)
	}

}

func run() error {
	q := services.Queue{
		Client: sqs.New(sess),
		URL:    os.Args[1],
	}

	msgs, err := q.GetMessages(20)
	if err != nil {
		return err
	}

	fmt.Println("Messages:")
	for _, msg := range msgs {
		fmt.Printf("%s>%s: %s\n", msg.From, msg.To, msg.Msg)
	}

	return nil
}
