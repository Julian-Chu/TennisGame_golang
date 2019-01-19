package tennisgame

type TennisGame struct {
	firstPlayerScoreTimes int
}

func (game TennisGame) Score() string {
	if game.firstPlayerScoreTimes == 1 {
		return "Fifteen Love"
	}
	if game.firstPlayerScoreTimes == 2 {
		return "Thirty Love"
	}

	return "Love All"
}

func (game *TennisGame) FirstPlayerScore() {
	game.firstPlayerScoreTimes++
}

func NewTennisGame() *TennisGame {
	return &TennisGame{}
}
