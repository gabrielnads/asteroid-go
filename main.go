package main

import (
	"log"
	"my-game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
