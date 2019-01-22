package test

import (
	"TennisGame/tennisgame"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LoveAll(t *testing.T) {
	g := NewGame()
	ScoreShouldBe(t, "Love All", g)
}

func Test_FifteenLove(t *testing.T) {
	g := NewGame()
	g.FirstPlayerScore()
	ScoreShouldBe(t, "Fifteen Love", g)
}

func Test_ThirtyLove(t *testing.T) {
	g := NewGame()
	GivenFirstPlayerScore(g, 2)
	ScoreShouldBe(t, "Thirty Love", g)
}

func Test_FiftyLove(t *testing.T) {
	g := NewGame()
	GivenFirstPlayerScore(g, 3)
	ScoreShouldBe(t, "Fifty Love", g)
}

func Test_LoveFifteen(t *testing.T) {
	g := NewGame()
	g.SecondPlayerScore()
	ScoreShouldBe(t, "Love Fifteen", g)
}

func Test_LoveThirty(t *testing.T) {
	g := NewGame()
	GivenSecondPlayerScore(g, 2)
	ScoreShouldBe(t, "Love Thirty", g)
}

func Test_FifteenAll(t *testing.T) {
	g := NewGame()
	GivenFirstPlayerScore(g, 1)
	GivenSecondPlayerScore(g, 1)
	ScoreShouldBe(t, "Fifteen All", g)
}

func Test_ThirtyAll(t *testing.T) {
	g := NewGame()
	GivenFirstPlayerScore(g, 2)
	GivenSecondPlayerScore(g, 2)
	ScoreShouldBe(t, "Thirty All", g)
}

func GivenSecondPlayerScore(g *tennisgame.TennisGame, times int) {
	for i := 0; i < times; i++ {
		g.SecondPlayerScore()
	}
}

func GivenFirstPlayerScore(g *tennisgame.TennisGame, times int) {
	for i := 0; i < times; i++ {
		g.FirstPlayerScore()
	}
}

func ScoreShouldBe(t *testing.T, expected string, g *tennisgame.TennisGame) bool {
	return assert.Equal(t, expected, g.Score(), "")
}

func NewGame() *tennisgame.TennisGame {
	g := &tennisgame.TennisGame{}
	return g
}
