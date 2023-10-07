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

	for _, enemy := range game.enemies {
		enemy.xLoc -= 2
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
	}
	if game.moveUp && game.yloc >= 5 {
		game.yloc -= 3
	}
	if game.moveDown && game.yloc <= 940 {
		game.yloc += 3
	}
	return nil
}
