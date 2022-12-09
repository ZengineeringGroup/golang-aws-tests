package poll

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type Poller struct {
	SQS sqsiface.SQSAPI
}

func (p *Poller) Poll(ctx context.Context, queueUrl string) ([]*sqs.Message, error) {
	input := sqs.ReceiveMessageInput{
		QueueUrl: &queueUrl,
	}

	out, err := p.SQS.ReceiveMessageWithContext(ctx, &input)
	if err != nil {
		return nil, err
	}

	return out.Messages, nil
}
