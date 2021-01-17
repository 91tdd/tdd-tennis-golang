package tennis

import "fmt"

var lookUp = make(map[int]string, 4)

func init() {
	lookUp[0] = "love"
	lookUp[1] = "fifteen"
	lookUp[2] = "thirty"
	lookUp[3] = "forty"
}

type Tennis struct {
	firstPlayerScoreTimes  int
	secondPlayerScoreTimes int
	FirstPlayerScoreName   string
}

func (t *Tennis) Score() string {
	if t.secondPlayerScoreTimes == t.firstPlayerScoreTimes {
		if t.firstPlayerScoreTimes >= 3 && t.secondPlayerScoreTimes >= 3 {
			return "deuce"
		}
		return fmt.Sprintf("%s all", lookUp[t.firstPlayerScoreTimes])
	}
	if t.firstPlayerScoreTimes >= 3 && t.secondPlayerScoreTimes >= 3 {
		if t.firstPlayerScoreTimes > t.secondPlayerScoreTimes {
			return fmt.Sprintf("%s adv", t.FirstPlayerScoreName)
		}
		return "Elsa adv"
	}
	return fmt.Sprintf("%s %s", lookUp[t.firstPlayerScoreTimes], lookUp[t.secondPlayerScoreTimes])
}

func (t *Tennis) FirstPlayerScore() {
	t.firstPlayerScoreTimes++
}

func (t *Tennis) SecondPlayerScore() {
	t.secondPlayerScoreTimes++
}
