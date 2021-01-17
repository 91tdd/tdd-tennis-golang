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
}

func (t *Tennis) Score() string {
	if t.secondPlayerScoreTimes > 0 {
		return fmt.Sprintf("%s %s", lookUp[t.firstPlayerScoreTimes], lookUp[t.secondPlayerScoreTimes])
	}
	if t.firstPlayerScoreTimes > 0 {
		return fmt.Sprintf("%s %s", lookUp[t.firstPlayerScoreTimes], lookUp[t.secondPlayerScoreTimes])
	}
	return "love all"
}

func (t *Tennis) FirstPlayerScore() {
	t.firstPlayerScoreTimes++
}

func (t *Tennis) SecondPlayerScore() {
	t.secondPlayerScoreTimes++
}
