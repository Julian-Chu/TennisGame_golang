package tennisgame

type TennisGame struct {
	firstPlayerScoreTimes int
}

var lookup = map[int]string{
	1: "Fifteen",
	2: "Thirty",
}

func (g TennisGame) Score() string {
	if g.firstPlayerScoreTimes == 1 || g.firstPlayerScoreTimes == 2 {
		return lookup[g.firstPlayerScoreTimes] + " Love"
	}

	return "Love All"

}

func (g *TennisGame) FirstPlayerScore() {
	g.firstPlayerScoreTimes++
}
