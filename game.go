package main

import (
	"errors"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var movementIsEnabled = true
var frameCount int = 0

// Available color palette
var colorLight color.Color = color.RGBA{199, 240, 216, 255}
var colorDark color.Color = color.RGBA{67, 82, 61, 255}

// Game state
const mapX int = 28
const mapY int = 4

// Game State
type Game struct {
	Height      int
	Width       int
	MapPosition int
	Map         [mapY][mapX]int
}

var game = Game{48, 84, 0, [mapY][mapX]int{
	{0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1},
	{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0}}}

var pixelSize int = game.Height / 4
var coloums = 7

var player *ebiten.Image
var playerPosition int = 0

func main() {
	if err := ebiten.Run(update, game.Width, game.Height, 5, "VrumVrum"); err != nil {
		panic(err)
	}
}

func update(screen *ebiten.Image) error {

	screen.Fill(colorLight)

	// Pause game if it's not not on top
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	handleInput()

	// Draw user
	drawPixel(0, playerPosition, screen)

	if game.MapPosition+coloums > len(game.Map[0]) {
		os.Exit(0)
	}

	for i := game.MapPosition; i < game.MapPosition+coloums; i++ {
		for j := 0; j < 4; j++ {
			if game.Map[j][i] == 1 {
				drawPixel(i-game.MapPosition, j, screen)
			}
		}
	}

	// User crashed with wall
	if game.Map[playerPosition][game.MapPosition] == 1 {
		return errors.New("Game Over")
	}

	// Screen update and player movement is reduced to 1 in every 60 frames
	frameCount++
	if frameCount%60 == 0 {
		game.MapPosition++
		movementIsEnabled = true
	}

	return nil
}

func drawPixel(x int, y int, screen *ebiten.Image) {
	var pixel *ebiten.Image
	if pixel == nil {
		pixel, _ = ebiten.NewImage(pixelSize, pixelSize, ebiten.FilterNearest)
	}
	pixel.Fill(colorDark)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x*pixelSize), float64(y*pixelSize))
	screen.DrawImage(pixel, opts)
}

func handleInput() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if movementIsEnabled {
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			if playerPosition < mapY-1 {
				playerPosition++
				movementIsEnabled = false
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			if playerPosition > 0 {
				playerPosition--
				movementIsEnabled = false
			}
		}
	}
	return nil
}
