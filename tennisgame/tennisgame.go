package tennisgame

type TennisGame struct {
	firstPlayerScoreTimes  int
	secondPlayerScoreTimes int
}

var lookup = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Fifty",
}

func (game TennisGame) Score() string {
	if game.firstPlayerScoreTimes != game.secondPlayerScoreTimes {
		return lookup[game.firstPlayerScoreTimes] + " " + lookup[game.secondPlayerScoreTimes]
	}

	return lookup[game.firstPlayerScoreTimes] + " All"
}

func (game *TennisGame) FirstPlayerScore() {
	game.firstPlayerScoreTimes++
}

func (game *TennisGame) SecondPlayerScore() {
	game.secondPlayerScoreTimes++
}

func NewTennisGame() *TennisGame {
	return &TennisGame{}
}
