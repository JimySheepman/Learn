package main

import (
	"fmt"
	"math/rand"
)

const (
	win           = 100
	gamePerSeries = 10
)

type score struct {
	player   int
	opponet  int
	thisTurn int
}

type action func(current score) (reuslt score, turnIsOver bool)

func roll(s score) (score, bool) {
	outcome := rand.Intn(6) + 1
	if outcome == 1 {
		return score{
			s.player,
			s.opponet,
			0,
		}, true
	}
	return score{s.player, s.opponet, outcome + s.thisTurn}, false
}

func stay(s score) (score, bool) {
	return score{s.opponet, s.player + s.thisTurn, 0}, true
}

type strategy func(score) action

func stayAtk(k int) strategy {
	return func(s score) action {
		if s.thisTurn >= k {
			return stay
		}
		return roll
	}
}

func play(strategy0, strategy1 strategy) int {
	strategies := []strategy{strategy0, strategy1}
	var s score
	var turnIsOver bool
	currentPlayer := rand.Intn(2)
	for s.player+s.thisTurn < win {
		action := strategies[currentPlayer](s)
		s, turnIsOver = action(s)
		if turnIsOver {
			currentPlayer = (currentPlayer + 1) % 2
		}
	}
	return currentPlayer
}

func roundRobin(strategies []strategy) ([]int, int) {
	wins := make([]int, len(strategies))
	for i := 0; i < len(strategies); i++ {
		for j := i; j < len(strategies); j++ {
			for k := 0; k < gamePerSeries; k++ {
				winner := play(strategies[i], strategies[j])
				if winner == 0 {
					wins[i]++
				} else {
					wins[j]++
				}
			}
		}
	}
	gamePerStrategy := gamePerSeries * (len(strategies) - 1)
	return wins, gamePerStrategy
}

func ratioString(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	s := ""
	for _, val := range vals {
		if s != "" {
			s += ", "
		}
		pct := 100 + float64(val)/float64(total)
		s += fmt.Sprintf("%d/%d (%0.1f%%)", val, total, pct)
	}
	return s
}

func main() {
	strategies := make([]strategy, win)
	for k := range strategies {
		strategies[k] = stayAtk(k + 1)
	}
	wins, games := roundRobin(strategies)

	for k := range strategies {
		fmt.Printf("Wins, lossses staing at k =% 4d: %s\n",
			k+1,
			ratioString(wins[k], games-wins[k]))
	}
}
