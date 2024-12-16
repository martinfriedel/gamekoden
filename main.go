package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
)

var hero *ebiten.Image
var grass *ebiten.Image
var grass2 *ebiten.Image

const (
	tileSize   = 128
	gameWidth  = 1280
	gameHeight = 720
	gameScale  = 0.5
	levelSize  = 100
)

type Game struct {
	playerX     int
	playerY     int
	playerSpeed int
	level       [][]int
}

func (g *Game) generateLevel() {
	g.level = make([][]int, levelSize)
	for y := 0; y < levelSize; y++ {
		g.level[y] = make([]int, levelSize)
		for x := 0; x < levelSize; x++ {
			g.level[y][x] = rand.Intn(2)
		}
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerY -= g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerY += g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerX = g.playerX - g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX = g.playerX + g.playerSpeed
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// render the background
	const realSize = tileSize * gameScale
	for y := 0; y < levelSize; y++ {
		for x := 0; x < levelSize; x++ {
			img := grass
			if g.level[x][y] == 1 {
				img = grass2
			}
			drawImage(screen, img, 0-g.playerX+x*realSize, 0-g.playerY+y*realSize)
		}
	}

	// render the player
	drawImage(screen, hero, gameWidth/2, gameHeight/2)

	// var position string = fmt.Sprintf("player x:%v y:%v", g.playerX, g.playerY)
	// ebitenutil.DebugPrint(screen, position)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameWidth, gameHeight
}

func init() {
	hero = loadImageOrCrash("img/hero.png")
	grass = loadImageOrCrash("img/grass.png")
	grass2 = loadImageOrCrash("img/grass2.png")
}

func main() {
	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Megaspiel")
	game := &Game{
		playerX:     2000,
		playerY:     2000,
		playerSpeed: 3,
	}
	game.generateLevel()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
