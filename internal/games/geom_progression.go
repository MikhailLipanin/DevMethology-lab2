package games

import (
	"DevMethology-lab2/pkg/rand"
	"DevMethology-lab2/pkg/slices"
	"fmt"
	"strconv"
)

const (
	geomProgressionRules = "What number is missing in the progression?"
	minLenOfProgression  = 5
	maxLenOfProgression  = 10
	minBeginVal          = 1
	maxBeginVal          = 10
	hiddenValueVal       = -1
)

type GeomProgression struct {
	data           []int64
	expectedResult singleValResultType
}

func NewGeomProgression() *GeomProgression {
	return &GeomProgression{}
}

func (g *GeomProgression) PrintGameRules() {
	fmt.Println(geomProgressionRules)
}

func (g *GeomProgression) ProcessNewRound() {
	// due to game rules, inside this method we fill g.expectedResult
	g.data = g.generateQuestionData()
}

func (g *GeomProgression) GetRoundData() []any {
	return slices.Map(g.data, func(v int64) any {
		if v == hiddenValueVal {
			return any("..")
		}
		return any(v)
	})
}

func (g *GeomProgression) GetExpectedResult() fmt.Stringer {
	return g.expectedResult
}

func (g *GeomProgression) CompareResult(result string) bool {
	res, err := g.parseResult(result)
	if err != nil {
		return false
	}
	if res != g.expectedResult {
		return false
	}
	return true
}

func (g *GeomProgression) parseResult(result string) (singleValResultType, error) {
	val, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return singleValResultType(val), nil
}

func (g *GeomProgression) generateQuestionData() []int64 {
	progressionLen := rand.RandIntForInterval(minLenOfProgression, maxLenOfProgression)

	res := make([]int64, 0, progressionLen)
	res = append(res, int64(rand.RandIntForInterval(minBeginVal, maxBeginVal)))

	mul := int64(rand.RandIntForInterval(minBeginVal, maxBeginVal))
	for i := 1; i < progressionLen; i++ {
		res = append(res, res[i-1]*mul)
	}

	// hide random value of progression
	idOfMissedVal := rand.RandIntForInterval(0, progressionLen-1)
	g.expectedResult = singleValResultType(res[idOfMissedVal])
	res[idOfMissedVal] = hiddenValueVal

	return res
}
