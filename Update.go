package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (game *SpaceShooter) Update() error {
	backgroundWidth := game.background.Bounds().Dx()
	maxX := backgroundWidth * 2
	game.backgroundXView -= 4
	game.backgroundXView %= maxX

	//FIXME: I feel like this should loop through the slice of enemy sprites to move them
	//In main(), line 34, an empty slice is made, it is then filled to capacity with
	//randomly placed enemy sprites and passed into the game struct
	for _, enemy := range game.enemies {
		enemy.xLoc -= 2
	}
	//FIXME: In the same regard I feel as though this should be looping through laser sprites to move them
	//In main(), line 29, empty slice made, that empty slice is then passed into the game struct
	//The list gets lasers appended in this function on line 43
	for _, laser := range game.lasers {
		laser.xLoc += 2
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		game.moveUp = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
		game.moveUp = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		game.moveDown = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		game.moveDown = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		game.soundPlayer.Rewind()
		game.soundPlayer.Play()
		fmt.Println("pew")
		game.lasers = append(game.lasers, NewLaser(game.xloc+65, game.yloc+25, game.laserPict))
	}
	if game.moveUp && game.yloc >= 5 {
		game.yloc -= 3
	}
	if game.moveDown && game.yloc <= 940 {
		game.yloc += 3
	}
	return nil
}
