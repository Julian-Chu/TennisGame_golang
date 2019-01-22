package tennisgame

type TennisGame struct {
	firstPlayerScoreTimes int
}

func (g TennisGame) Score() string {
	if g.firstPlayerScoreTimes == 1 {
		return "Fifteen Love"
	}
	if g.firstPlayerScoreTimes == 2 {
		return "Thirty Love"
	}

	return "Love All"

}

func (g *TennisGame) FirstPlayerScore() {
	g.firstPlayerScoreTimes++
}
