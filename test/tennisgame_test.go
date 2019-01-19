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
	t.ScoreShouldBe("Love All")
}

func (t TennisGameTestSuite) Test_FifteenLove() {
	t.TennisGame.FirstPlayerScore()
	t.ScoreShouldBe("Fifteen Love")
}

func (t TennisGameTestSuite) Test_ThirtyLove() {
	t.TennisGame.FirstPlayerScore()
	t.TennisGame.FirstPlayerScore()
	t.ScoreShouldBe("Thirty Love")
}

func (t TennisGameTestSuite) ScoreShouldBe(expected string) bool {
	return assert.Equal(t.T(), expected, t.TennisGame.Score(), "")
}
