package games

import (
	"DevMethology-lab2/pkg/math"
	"DevMethology-lab2/pkg/rand"
	"DevMethology-lab2/pkg/slices"
	"fmt"
	"strconv"
)

const (
	gameQuestion       = "Find the smallest common multiple of given numbers."
	lowIntervalOfNums  = 1
	highIntervalOfNums = 100
)

type resultType int

func (r resultType) String() string {
	return fmt.Sprintf("%d", r)
}

type LcmGame struct {
	data           []int
	cntOfNums      int
	expectedResult resultType
}

func NewLcmGame(cntOfNums int) *LcmGame {
	return &LcmGame{
		cntOfNums: cntOfNums,
	}
}

func (l *LcmGame) PrintGameRules() {
	fmt.Println(gameQuestion)
}

func (l *LcmGame) ProcessNewRound() {
	l.data = l.generateQuestionData()
	l.expectedResult = l.evaluateResult()
}

func (l *LcmGame) GetRoundData() []any {
	return slices.Map(l.data, func(v int) any {
		return any(v)
	})
}

func (l *LcmGame) GetExpectedResult() fmt.Stringer {
	return l.expectedResult
}

func (l *LcmGame) CompareResult(result string) bool {
	res, err := l.parseResult(result)
	if err != nil {
		return false
	}
	if res != l.expectedResult {
		return false
	}
	return true
}

func (l *LcmGame) parseResult(result string) (resultType, error) {
	val, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return resultType(val), nil
}

func (l *LcmGame) generateQuestionData() []int {
	res := make([]int, 0, l.cntOfNums)
	for i := 0; i < l.cntOfNums; i++ {
		res = append(res, rand.RandIntForInterval(lowIntervalOfNums, highIntervalOfNums))
	}
	return res
}

func (l *LcmGame) evaluateResult() resultType {
	return resultType(math.LcmOfArray(l.data))
}
