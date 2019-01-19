package tennisgame

type TennisGame struct {
	firstPlayerScoreTimes int
}

var lookup = map[int]string{
	1: "Fifteen",
	2: "Thirty",
}

func (game TennisGame) Score() string {
	if game.firstPlayerScoreTimes == 1 || game.firstPlayerScoreTimes == 2 {
		return lookup[game.firstPlayerScoreTimes] + " Love"
	}

	return "Love All"
}

func (game *TennisGame) FirstPlayerScore() {
	game.firstPlayerScoreTimes++
}

func NewTennisGame() *TennisGame {
	return &TennisGame{}
}
