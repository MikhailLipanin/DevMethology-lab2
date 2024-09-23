package main

import (
	"DevMethology-lab2/internal"
	"DevMethology-lab2/internal/games"
	"flag"
	"fmt"
	"os"
)

var (
	gameMode = flag.String("mode", "lcm", "Choose game mode to play - 'lcm' or 'geom-progression'. Default is 'lcm'.")
)

func main() {
	flag.Parse()

	var gameProvider games.GameModeProvider

	switch *gameMode {
	case "lcm":
		gameProvider = games.NewLcmGame(3)
	case "geom-progression":

	default:
		fmt.Println("Wrong game mode! Specify the game mode: 'lcm' or 'geom-progression'")
		os.Exit(1)
	}

	internal.NewGame(3, gameProvider).Start()
}
