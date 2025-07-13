package game

import (
	"fmt"
	"image/color"
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StateGameOver
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Timer
	score            int
	star             []*Stars
	state            GameState
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
		state:            StateMenu,
	}
	g.player = NewPlayer(g)

	for i := 0; i < 100; i++ {
		g.star = append(g.star, NewStar())
	}

	return g
}

func (g *Game) Update() error {
	switch g.state {
	case StateMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.state = StatePlaying
		}
	case StatePlaying:
		g.player.Update()

		for _, l := range g.lasers {
			l.Update()
		}

		g.meteorSpawnTimer.Update()
		if g.meteorSpawnTimer.IsReady() {
			g.meteorSpawnTimer.Reset()
			m := NewMeteor()
			g.meteors = append(g.meteors, m)
		}
		for _, m := range g.meteors {
			m.Update()
		}

		for _, m := range g.meteors {
			if m.Collider().Intersects(g.player.Collider()) {
				g.state = StateGameOver
			}
		}

		for i := len(g.meteors) - 1; i >= 0; i-- {
			m := g.meteors[i]
			for j := len(g.lasers) - 1; j >= 0; j-- {
				l := g.lasers[j]
				if m.Collider().Intersects(l.Collider()) {
					g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
					g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
					g.score++
					break
				}
			}
		}

		for _, s := range g.star {
			s.Update()
		}
	case StateGameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Reset()
			g.state = StatePlaying
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateMenu:
		msg := "> Iniciar Jogo"
		text.Draw(screen, msg, assets.FontUi, 100, 100, color.White)

	case StatePlaying:
		for _, s := range g.star {
			s.Draw(screen)
		}
		g.player.Draw(screen)
		for _, l := range g.lasers {
			l.Draw(screen)
		}
		for _, m := range g.meteors {
			m.Draw(screen)
		}
		text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.FontUi, 20, 100, color.White)

	case StateGameOver:
		msg := "GAME OVER\nAperte ENTER para reiniciar"
		text.Draw(screen, msg, assets.FontUi, 20, 100, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeigth int) (int, int) {
	return screenWidth, screenHeigth
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player.Reset()
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
}
