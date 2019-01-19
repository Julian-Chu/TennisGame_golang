package tennisgame

type TennisGame struct {
}

func (game TennisGame) Score() string {
	return "Love All"
}

func NewTennisGame() *TennisGame {
	return &TennisGame{}
}
