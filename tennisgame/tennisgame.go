package tennisgame

import "math"

type TennisGame struct {
	firstPlayerName        string
	secondPlayerName       string
	firstPlayerScoreTimes  int
	secondPlayerScoreTimes int
}

func NewTennisGame(firstPlayerName string, secondPlayerName string) *TennisGame {
	return &TennisGame{firstPlayerName: firstPlayerName, secondPlayerName: secondPlayerName}
}

var lookup = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Forty",
}

func (g TennisGame) Score() string {
	if g.isScoreDifferent() {
		if g.isReadyForWin() {
			if g.isAdv() {
				return g.getAdvPlayer()
			}
			return g.getWinner()
		}
		return g.getNormalScore()
	}

	if g.isDeuce() {
		return "Deuce"
	}
	return g.getAll()
}

func (g TennisGame) getAll() string {
	return lookup[g.firstPlayerScoreTimes] + " All"
}

func (g TennisGame) isDeuce() bool {
	return g.firstPlayerScoreTimes >= 3
}

func (g TennisGame) getNormalScore() string {
	return lookup[g.firstPlayerScoreTimes] + " " + lookup[g.secondPlayerScoreTimes]
}

func (g TennisGame) getWinner() string {
	return g.getPlayer() + " Win"
}

func (g TennisGame) getAdvPlayer() string {
	return g.getPlayer() + " Adv"
}

func (g TennisGame) isAdv() bool {
	return math.Abs(float64(g.firstPlayerScoreTimes-g.secondPlayerScoreTimes)) == 1
}

func (g TennisGame) isReadyForWin() bool {
	return g.firstPlayerScoreTimes > 3 || g.secondPlayerScoreTimes > 3
}

func (g TennisGame) isScoreDifferent() bool {
	return g.firstPlayerScoreTimes != g.secondPlayerScoreTimes
}

func (g TennisGame) getPlayer() string {
	player := g.firstPlayerName
	if g.firstPlayerScoreTimes < g.secondPlayerScoreTimes {
		player = g.secondPlayerName
	}
	return player
}

func (g *TennisGame) FirstPlayerScore() {
	g.firstPlayerScoreTimes++
}

func (g *TennisGame) SecondPlayerScore() {
	g.secondPlayerScoreTimes++

}
