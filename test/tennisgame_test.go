package tennisgame_test

import (
	"TennisGame/tennisgame"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TennisGameSuite struct {
	suite.Suite
	*tennisgame.TennisGame
}

func TestSuiteInit(t *testing.T) {
	suite.Run(t, new(TennisGameSuite))
}

func (t *TennisGameSuite) SetupTest() {
	t.TennisGame = tennisgame.NewTennisGame("Joey", "Tom")

}

func (t TennisGameSuite) Test_LoveAll() {
	t.ScoreShouldBe("Love All")
}

func (t TennisGameSuite) Test_FifteenLove() {
	t.TennisGame.FirstPlayerScore()
	t.ScoreShouldBe("Fifteen Love")
}

func (t TennisGameSuite) Test_ThirtyLove() {
	t.FirstPlayerScoreTimes(2)
	t.ScoreShouldBe("Thirty Love")
}

func (t TennisGameSuite) Test_FortyLove() {
	t.FirstPlayerScoreTimes(3)
	t.ScoreShouldBe("Forty Love")
}

func (t TennisGameSuite) Test_LoveFifteen() {
	t.TennisGame.SecondPlayerScore()
	t.ScoreShouldBe("Love Fifteen")
}

func (t TennisGameSuite) Test_LoveThirty() {
	t.SecondPlayerScoreTimes(2)
	t.ScoreShouldBe("Love Thirty")
}

func (t TennisGameSuite) Test_FifteenAll() {
	t.FirstPlayerScoreTimes(1)
	t.SecondPlayerScoreTimes(1)
	t.ScoreShouldBe("Fifteen All")
}

func (t TennisGameSuite) Test_ThirtyAll() {
	t.FirstPlayerScoreTimes(2)
	t.SecondPlayerScoreTimes(2)
	t.ScoreShouldBe("Thirty All")
}

func (t TennisGameSuite) Test_Deuce() {
	t.FirstPlayerScoreTimes(3)
	t.SecondPlayerScoreTimes(3)
	t.ScoreShouldBe("Deuce")
}

func (t TennisGameSuite) Test_FirstPlayerAdv() {
	t.FirstPlayerScoreTimes(4)
	t.SecondPlayerScoreTimes(3)
	t.ScoreShouldBe("Joey Adv")
}

func (t TennisGameSuite) Test_SecondPlayerAdv() {
	t.FirstPlayerScoreTimes(3)
	t.SecondPlayerScoreTimes(4)
	t.ScoreShouldBe("Tom Adv")
}

func (t TennisGameSuite) Test_Deuce44() {
	t.FirstPlayerScoreTimes(4)
	t.SecondPlayerScoreTimes(4)
	t.ScoreShouldBe("Deuce")
}

func (t TennisGameSuite) Test_FirstPlayerWin() {
	t.FirstPlayerScoreTimes(5)
	t.SecondPlayerScoreTimes(3)
	t.ScoreShouldBe("Joey Win")
}

func (t TennisGameSuite) SecondPlayerScoreTimes(times int) {
	for i := 0; i < times; i++ {
		t.TennisGame.SecondPlayerScore()
	}
}

func (t TennisGameSuite) FirstPlayerScoreTimes(times int) {
	for i := 0; i < times; i++ {
		t.TennisGame.FirstPlayerScore()
	}
}

func (t TennisGameSuite) ScoreShouldBe(expect string) bool {
	return t.Equal(expect, t.TennisGame.Score())
}
