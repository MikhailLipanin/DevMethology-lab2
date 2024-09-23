package games

import "fmt"

type GameModeProvider interface {
	PrintGameRules()
	ProcessNewRound()
	GetRoundData() []any
	GetExpectedResult() fmt.Stringer
	CompareResult(string) bool
}
