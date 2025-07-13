package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2 // Metade da Largura da Imagem do Laser
	halfH := float64(bounds.Dy()) / 2 // Metade da Altura da Imagem do Laser

	position.X -= halfW
	position.Y -= halfH

	return &Laser{
		image:    image,
		position: position,
	}

}

func (l *Laser) Update() {
	speed := 7.0
	l.position.Y += -speed

}
func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// posicao X e Y que a imagem sera desenhada
	op.GeoM.Translate(l.position.X, l.position.Y)
	// Desenha a Imagem na Tela
	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()

	return NewRect(l.position.X,
		l.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()))
}
