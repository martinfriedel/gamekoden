package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func drawImage(screen *ebiten.Image, img *ebiten.Image, x int, y int) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(gameScale, gameScale)
	options.GeoM.Translate(float64(x), float64(y))
	options.Filter = ebiten.FilterLinear
	screen.DrawImage(img, options)
}

func loadImageOrCrash(fileLocation string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
