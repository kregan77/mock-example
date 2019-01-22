package thing

import (
	"testing"

	"github.com/kregan77/mock-example/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ThingTestSuite struct {
	suite.Suite
	thinger mocks.Thinger
}

func (s *ThingTestSuite) SetupSuite() {
	s.thinger = mocks.Thinger{}
}

func (s *ThingTestSuite) SetupTest() {
	s.thinger.On("Thing", mock.Anything).Return("default")
}

func (s *ThingTestSuite) TestFuncDontCareAboutThinger() {
	expected := "default"
	actual := FuncToTest(&s.thinger, "doesn't matter")
	assert.Equal(s.T(), expected, actual, "ruh roh")
}

func (s *ThingTestSuite) TestFuncOverrideThingerMock() {
	expected := "something different"
	val := "specific value"
	s.thinger.On("Thing", val).Return(expected)
	actual := FuncToTest(&s.thinger, val)
	assert.Equal(s.T(), expected, actual, "ruh roh")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ThingTestSuite))
}
