package tennisgame

type TennisGame struct {
	firstPlayerScoreTimes  int
	secondPlayerScoreTimes int
}

var lookup = map[int]string{
	1: "Fifteen",
	2: "Thirty",
	3: "Fifty",
}

func (game TennisGame) Score() string {
	if game.firstPlayerScoreTimes > 0 {
		return lookup[game.firstPlayerScoreTimes] + " Love"
	}
	if game.secondPlayerScoreTimes > 0 {
		return "Love " + lookup[game.secondPlayerScoreTimes]
	}

	return "Love All"
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
