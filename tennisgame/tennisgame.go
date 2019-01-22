package tennisgame

import "math"

type TennisGame struct {
	firstPlayerScoreTimes  int
	secondPlayerScoreTimes int
	firstPlayerName        string
	secondPlayerName       string
}

func NewTennisGame(firstPlayerName string, secondPlayerName string) *TennisGame {
	return &TennisGame{firstPlayerName: firstPlayerName, secondPlayerName: secondPlayerName}
}

var lookup = map[int]string{

	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Fifty",
}

func (g TennisGame) Score() string {
	if g.firstPlayerScoreTimes != g.secondPlayerScoreTimes {
		player := g.firstPlayerName
		if g.firstPlayerScoreTimes < g.secondPlayerScoreTimes {
			player = g.secondPlayerName
		}
		if g.firstPlayerScoreTimes > 3 || g.secondPlayerScoreTimes > 3 {
			if math.Abs(float64(g.firstPlayerScoreTimes-g.secondPlayerScoreTimes)) == 1 {
				return player + " Adv"
			}
		}
		return lookup[g.firstPlayerScoreTimes] + " " + lookup[g.secondPlayerScoreTimes]
	}

	if g.firstPlayerScoreTimes == 3 {
		return "Deuce"
	}
	return lookup[g.firstPlayerScoreTimes] + " All"

}

func (g *TennisGame) FirstPlayerScore() {
	g.firstPlayerScoreTimes++
}

func (g *TennisGame) SecondPlayerScore() {
	g.secondPlayerScoreTimes++
}
