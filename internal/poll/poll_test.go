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
	// Being more specific with the On calls allows us to create different sets of logic for different input
	call := s.SQS.On("ReceiveMessageWithContext", mock.Anything, mock.Anything, mock.Anything)
	call.Return(&sqs.ReceiveMessageOutput{}, errors.New("error"))

	msgs, err := s.Poller.Poll(context.Background(), "queue")
	s.Error(err)
	s.Nil(msgs)
}

// SQSMock only needs mock.Mock and the interface it is mocking, sqsiface.SQSAPI
// We only need to define the methods that are actually used
type SQSMock struct {
	mock.Mock
	sqsiface.SQSAPI
}

// This simple mock allows us to override the function behavior within the tests itself
// using the mock.Call functionality
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
	// Create the mock and use it in our Poller
	sqsAPI := new(SQSMock)

	// Keep a reference in the Suite itself to avoid the need to case Poller.SQS in tests
	s.SQS = sqsAPI
	s.Poller = poll.Poller{
		SQS: sqsAPI,
	}
}

func TestPollSuite(t *testing.T) {
	s := new(PollSuite)
	suite.Run(t, s)
}
