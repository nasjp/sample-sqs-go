package services

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type Queue struct {
	Client sqsiface.SQSAPI
	URL    string
}

type Message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

func (q *Queue) GetMessages(waitTimeout int64) ([]Message, error) {
	params := sqs.ReceiveMessageInput{
		QueueUrl: aws.String(q.URL),
	}
	if waitTimeout > 0 {
		params.WaitTimeSeconds = aws.Int64(waitTimeout)
	}
	resp, err := q.Client.ReceiveMessage(&params)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages, %v", err)
	}

	msgs := make([]Message, len(resp.Messages))
	for i, msg := range resp.Messages {
		parsedMsg := Message{}
		if err := json.Unmarshal([]byte(aws.StringValue(msg.Body)), &parsedMsg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal message, %v", err)
		}

		msgs[i] = parsedMsg
	}

	return msgs, nil
}
