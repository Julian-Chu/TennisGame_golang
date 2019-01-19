package tennisgame_test

import (
	"TennisGame/tennisgame"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TennisGameTestSuite struct {
	suite.Suite
	*tennisgame.TennisGame
}

func TestSuiteInit(t *testing.T) {
	suite.Run(t, new(TennisGameTestSuite))
}

func (t *TennisGameTestSuite) SetupTest() {
	t.TennisGame = tennisgame.NewTennisGame()
}

func (t TennisGameTestSuite) Test_LoveAll() {
	assert.Equal(t.T(), "Love All", t.TennisGame.Score(), "")

}
