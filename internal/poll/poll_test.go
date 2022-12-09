package poll_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zengineeringgroup/golang-aws-tests/internal/poll"
)

func (s *PollSuite) TestPoll() {
	call := s.SQS.On("ReceiveMessageWithContext", mock.Anything, mock.Anything, mock.Anything)
	call.Return(&sqs.ReceiveMessageOutput{}, errors.New("error"))

	msgs, err := s.Poller.Poll(context.Background(), "queue")
	s.Error(err)
	s.Nil(msgs)
}

type SQSMock struct {
	mock.Mock
	sqsiface.SQSAPI
}

func (m *SQSMock) ReceiveMessageWithContext(ctx aws.Context, input *sqs.ReceiveMessageInput, opts ...request.Option) (*sqs.ReceiveMessageOutput, error) {
	args := m.Called(ctx, input, opts)
	return args[0].(*sqs.ReceiveMessageOutput), args.Error(1)
}

type PollSuite struct {
	suite.Suite
	Poller poll.Poller
	SQS    *SQSMock
}

func (s *PollSuite) SetupSuite() {
	sqsAPI := new(SQSMock)
	s.SQS = sqsAPI
	s.Poller = poll.Poller{
		SQS: sqsAPI,
	}
}

func TestPollSuite(t *testing.T) {
	s := new(PollSuite)
	suite.Run(t, s)
}
