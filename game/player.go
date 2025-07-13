package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,
	}
	return &Player{
		image:             image,
		game:              game,
		position:          position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Reset() {
	bounds := p.image.Bounds()
	halfW := float64(bounds.Dx()) / 2
	p.position.X = (screenWidth / 2) - halfW
	p.position.Y = 500

	p.laserLoadingTimer.Reset()
}

func (p *Player) Update() {
	speed := 6.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}
	playerWidth := float64(p.image.Bounds().Dx())
	if p.position.X < 0 {
		p.position.X = 0
	} else if p.position.X+playerWidth > screenWidth {
		p.position.X = screenWidth - playerWidth
	}

	p.laserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()

		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfW,
			p.position.Y - halfH/2,
		}
		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewRect(p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()))
}
