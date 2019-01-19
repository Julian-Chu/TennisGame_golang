package tennisgame

import "math"

type TennisGame struct {
	firstPlayerScoreTimes  int
	secondPlayerScoreTimes int
	firstPlayerName        string
	secondPlayerName       string
}

var lookup = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Fifty",
}

func (game TennisGame) Score() string {
	if game.isScoreDifferent() {
		if game.IsReadyForWin() {
			if game.IsAdv() {
				return game.getAdvPlayer(game.getPlayer())
			}
			return game.getWinner(game.getPlayer())
		}
		return game.getNormalScore()
	}

	if game.IsDeuce() {
		return game.getDeuce()
	}
	return game.getAll()
}

func (game TennisGame) isScoreDifferent() bool {
	return game.firstPlayerScoreTimes != game.secondPlayerScoreTimes
}

func (game TennisGame) getPlayer() string {
	player := game.firstPlayerName
	if game.firstPlayerScoreTimes < game.secondPlayerScoreTimes {
		player = game.secondPlayerName
	}
	return player
}

func (game TennisGame) IsReadyForWin() bool {
	return game.firstPlayerScoreTimes > 3 || game.secondPlayerScoreTimes > 3
}

func (game TennisGame) IsAdv() bool {
	return math.Abs(float64(game.firstPlayerScoreTimes-game.secondPlayerScoreTimes)) == 1
}

func (game TennisGame) getAdvPlayer(player string) string {
	return player + " Adv"
}

func (game TennisGame) getWinner(player string) string {
	return player + " Win"
}

func (game TennisGame) getNormalScore() string {
	return lookup[game.firstPlayerScoreTimes] + " " + lookup[game.secondPlayerScoreTimes]
}

func (game TennisGame) IsDeuce() bool {
	return game.firstPlayerScoreTimes >= 3
}

func (game TennisGame) getDeuce() string {
	return "Deuce"
}

func (game TennisGame) getAll() string {
	return lookup[game.firstPlayerScoreTimes] + " All"
}

func (game *TennisGame) FirstPlayerScore() {
	game.firstPlayerScoreTimes++
}

func (game *TennisGame) SecondPlayerScore() {
	game.secondPlayerScoreTimes++
}

func NewTennisGame(firstPlayerName, secondPlayerName string) *TennisGame {
	return &TennisGame{firstPlayerName: firstPlayerName, secondPlayerName: secondPlayerName}
}
