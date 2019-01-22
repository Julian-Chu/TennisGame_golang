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

func (g TennisGame) Score() string {
	if g.firstPlayerScoreTimes > 0 {
		return lookup[g.firstPlayerScoreTimes] + " Love"
	}
	if g.secondPlayerScoreTimes > 0 {
		return "Love " + lookup[g.secondPlayerScoreTimes]
	}

	return "Love All"

}

func (g *TennisGame) FirstPlayerScore() {
	g.firstPlayerScoreTimes++
}

func (g *TennisGame) SecondPlayerScore() {
	g.secondPlayerScoreTimes++
}
