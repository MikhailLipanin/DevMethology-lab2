package internal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MikhailLipanin/DevMethology-lab2/internal/games"
)

const (
	greetingsMsg    = "Welcome to the Brain Games!"
	askNameMsg      = "May I have your name? "
	successRoundMsg = "Correct!"
	successGameMsg  = "Congratulations, %s!"
	failedRoundMsg  = "'%v' is wrong answer ;(. Correct answer was '%v'"
	failedGameMsg   = "Let's try again, %s!"
)

type Game struct {
	playerName   string
	playerAnswer string
	CntOfRounds  int
	reader       *bufio.Scanner
	mode         games.GameModeProvider
}

func NewGame(cntOfRounds int, gameProvider games.GameModeProvider) *Game {
	return &Game{
		CntOfRounds: cntOfRounds,
		reader:      bufio.NewScanner(os.Stdin),
		mode:        gameProvider,
	}
}

func (g *Game) SuccessRound() {
	fmt.Println(successRoundMsg)
}

func (g *Game) SuccessGame() {
	fmt.Println(fmt.Sprintf(successGameMsg, g.playerName))
}

func (g *Game) FailRound() {
	fmt.Println(fmt.Sprintf(failedRoundMsg, g.playerAnswer, g.mode.GetExpectedResult().String()))
}

func (g *Game) FailGame() {
	fmt.Println(fmt.Sprintf(failedGameMsg, g.playerName))
}

func (g *Game) Greetings() {
	fmt.Println(greetingsMsg)
	fmt.Print(askNameMsg)
	g.playerName = g.readLine()
}

func (g *Game) AskPlayer() {
	fmt.Printf("Question: %v\n", g.mode.GetRoundData())
}

func (g *Game) GetPlayerAnswer() {
	fmt.Print("Your answer: ")
	g.playerAnswer = g.readLine()
}

func (g *Game) readLine() string {
	var txt string
	for g.reader.Scan() {
		txt = g.reader.Text()
		if len(txt) == 0 {
			fmt.Println("You have provided empty string. Try again!")
			continue
		}
		break
	}
	return txt
}

func (g *Game) Start() {
	g.Greetings()
	for i := 0; i < g.CntOfRounds; i++ {
		g.mode.PrintGameRules()
		g.mode.ProcessNewRound()
		g.AskPlayer()
		g.GetPlayerAnswer()
		if !g.mode.CompareResult(g.playerAnswer) {
			g.FailRound()
			g.FailGame()
			return
		}
		g.SuccessRound()
	}
	g.SuccessGame()
}
