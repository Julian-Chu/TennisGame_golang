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

func ScoreShouldBe(t *testing.T, expected string, g *tennisgame.TennisGame) bool {
	return assert.Equal(t, expected, g.Score(), "")
}

func NewGame() *tennisgame.TennisGame {
	g := &tennisgame.TennisGame{}
	return g
}
