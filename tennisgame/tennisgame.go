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

func (g TennisGame) Score() string {
	if g.firstPlayerScoreTimes != g.secondPlayerScoreTimes {
		return lookup[g.firstPlayerScoreTimes] + " " + lookup[g.secondPlayerScoreTimes]
	}

	return lookup[g.firstPlayerScoreTimes] + " All"

}

func (g *TennisGame) FirstPlayerScore() {
	g.firstPlayerScoreTimes++
}

func (g *TennisGame) SecondPlayerScore() {
	g.secondPlayerScoreTimes++
}
