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
	t.TennisGame = tennisgame.NewTennisGame("Joey", "Tom")
}

func (t TennisGameTestSuite) Test_LoveAll() {
	t.ScoreShouldBe("Love All")
}

func (t TennisGameTestSuite) Test_FifteenLove() {
	t.TennisGame.FirstPlayerScore()
	t.ScoreShouldBe("Fifteen Love")
}

func (t TennisGameTestSuite) Test_ThirtyLove() {
	t.GivenFirstPlayerScore(2)
	t.ScoreShouldBe("Thirty Love")
}

func (t TennisGameTestSuite) Test_FiftyLove() {
	t.GivenFirstPlayerScore(3)
	t.ScoreShouldBe("Fifty Love")
}

func (t TennisGameTestSuite) Test_LoveFifteen() {
	t.TennisGame.SecondPlayerScore()
	t.ScoreShouldBe("Love Fifteen")
}

func (t TennisGameTestSuite) Test_LoveThirty() {
	t.GivenSecondPlayerScore(2)
	t.ScoreShouldBe("Love Thirty")
}

func (t TennisGameTestSuite) Test_FifteenAll() {
	t.GivenFirstPlayerScore(1)
	t.GivenSecondPlayerScore(1)
	t.ScoreShouldBe("Fifteen All")
}

func (t TennisGameTestSuite) Test_ThirtyAll() {
	t.GivenFirstPlayerScore(2)
	t.GivenSecondPlayerScore(2)
	t.ScoreShouldBe("Thirty All")
}

func (t TennisGameTestSuite) Test_Deuce() {
	t.GivenDeuce()
	t.ScoreShouldBe("Deuce")
}

func (t TennisGameTestSuite) Test_FirstPlayerAdv() {
	t.GivenDeuce()
	t.GivenFirstPlayerScore(1)
	t.ScoreShouldBe("Joey Adv")
}

func (t TennisGameTestSuite) Test_SecondPlayerAdv() {
	t.GivenDeuce()
	t.GivenSecondPlayerScore(1)
	t.ScoreShouldBe("Tom Adv")
}

func (t TennisGameTestSuite) Test_Deuce44() {
	t.GivenDeuce()
	t.GivenFirstPlayerScore(1)
	t.GivenSecondPlayerScore(1)
	t.ScoreShouldBe("Deuce")
}

func (t TennisGameTestSuite) Test_FirstPlayerWin() {
	t.GivenDeuce()
	t.GivenFirstPlayerScore(2)
	t.ScoreShouldBe("Joey Win")
}

func (t TennisGameTestSuite) GivenDeuce() {
	t.GivenFirstPlayerScore(3)
	t.GivenSecondPlayerScore(3)
}

func (t TennisGameTestSuite) GivenSecondPlayerScore(times int) {
	for i := 0; i < times; i++ {
		t.TennisGame.SecondPlayerScore()
	}
}

func (t TennisGameTestSuite) GivenFirstPlayerScore(times int) {
	for i := 0; i < times; i++ {
		t.TennisGame.FirstPlayerScore()
	}
}

func (t TennisGameTestSuite) ScoreShouldBe(expected string) bool {
	return assert.Equal(t.T(), expected, t.TennisGame.Score(), "")
}
