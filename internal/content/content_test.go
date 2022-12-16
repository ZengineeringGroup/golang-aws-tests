package content_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zengineeringgroup/golang-aws-tests/internal/content"
)

type ContentMock struct {
	mock.Mock
}

func (m *ContentMock) Get()     {}
func (m ContentMock) Put()      {}
func (m *ContentMock) GetStar() {}
func (m ContentMock) PutStar()  {}

type ContentSuite struct {
	suite.Suite
	d content.DataLayer
}

func (s *ContentSuite) SetupSuite() {
	cm := new(ContentMock)
	s.d = content.DataLayer{
		Manager: cm,
	}
}

func TestContentSuite(t *testing.T) {
	suite.Run(t, new(ContentSuite))
}
