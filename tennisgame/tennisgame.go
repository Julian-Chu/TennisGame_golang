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
	if isScoreDifferent(g) {
		if isReadyForWin(g) {
			if isAdv(g) {
				return getAdvPlayer(getPlayer(g))
			}
			return getWinner(getPlayer(g))
		}
		return getNormalScore(g)
	}

	if isDeuce(g) {
		return getDeuce()

	}
	return getAll(g)

}

func isScoreDifferent(g TennisGame) bool {
	return g.firstPlayerScoreTimes != g.secondPlayerScoreTimes
}

func getPlayer(g TennisGame) string {
	player := g.firstPlayerName
	if g.firstPlayerScoreTimes < g.secondPlayerScoreTimes {
		player = g.secondPlayerName
	}
	return player
}

func isReadyForWin(g TennisGame) bool {
	return g.firstPlayerScoreTimes > 3 || g.secondPlayerScoreTimes > 3
}

func isAdv(g TennisGame) bool {
	return math.Abs(float64(g.firstPlayerScoreTimes-g.secondPlayerScoreTimes)) == 1
}

func getAdvPlayer(player string) string {
	return player + " Adv"
}

func getWinner(player string) string {
	return player + " Win"
}

func getNormalScore(g TennisGame) string {
	return lookup[g.firstPlayerScoreTimes] + " " + lookup[g.secondPlayerScoreTimes]
}

func isDeuce(g TennisGame) bool {
	return g.firstPlayerScoreTimes >= 3
}

func getDeuce() string {
	return "Deuce"
}

func getAll(g TennisGame) string {
	return lookup[g.firstPlayerScoreTimes] + " All"
}

func (g *TennisGame) FirstPlayerScore() {
	g.firstPlayerScoreTimes++
}

func (g *TennisGame) SecondPlayerScore() {
	g.secondPlayerScoreTimes++
}
