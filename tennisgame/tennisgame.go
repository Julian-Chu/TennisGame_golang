package tennisgame

type TennisGame struct {
	firstPlayerScoreTimes int
}

var lookup = map[int]string{
	1: "Fifteen",
	2: "Thirty",
	3: "Fifty",
}

func (g TennisGame) Score() string {
	if g.firstPlayerScoreTimes > 0 {
		return lookup[g.firstPlayerScoreTimes] + " Love"
	}

	return "Love All"

}

func (g *TennisGame) FirstPlayerScore() {
	g.firstPlayerScoreTimes++
}
