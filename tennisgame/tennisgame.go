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
	if game.firstPlayerScoreTimes != game.secondPlayerScoreTimes {
		player := game.firstPlayerName
		if game.firstPlayerScoreTimes < game.secondPlayerScoreTimes {
			player = game.secondPlayerName
		}
		if game.firstPlayerScoreTimes > 3 || game.secondPlayerScoreTimes > 3 {
			if math.Abs(float64(game.firstPlayerScoreTimes-game.secondPlayerScoreTimes)) == 1 {
				return player + " Adv"
			}
			return player + " Win"
		}
		return lookup[game.firstPlayerScoreTimes] + " " + lookup[game.secondPlayerScoreTimes]
	}

	if game.firstPlayerScoreTimes >= 3 {
		return "Deuce"
	}
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
