package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math/rand"
)

func (game *GameStruct) Update() error {
	backgroundWidth := game.background.Bounds().Dx()
	maxX := backgroundWidth * 2
	game.backgroundXView -= 4
	game.backgroundXView %= maxX

	for i, _ := range game.enemies {
		game.enemies[i].xLoc -= 2
		if game.enemies[i].xLoc < -80 {
			game.enemies[i] = game.enemies[len(game.enemies)-1]
			game.enemies[len(game.enemies)-1] = NewEnemy(1010, rand.Intn(928), game.enemyPict)
			game.score -= 1
		}
	}
	for i, _ := range game.lasers {
		for j, _ := range game.enemies {
			if CheckCollision(game.enemies[j], game.lasers[i], game) {
				game.popSound.Rewind()
				game.popSound.Play()
				game.score += 1
				game.enemies[j] = NewEnemy(rand.Intn(928)+928, rand.Intn(928), game.enemyPict)
				game.lasers[i] = NewLaser(2000, 0, game.laserPict)
			}
		}
	}
	for i, _ := range game.lasers {
		game.lasers[i].xLoc += 10
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
		game.pewPewSound.Rewind()
		game.pewPewSound.Play()
		game.lasers = append(game.lasers, NewLaser(game.xLoc+65, game.yloc+25, game.laserPict))
	}
	if game.moveUp && game.yloc >= 5 {
		game.yloc -= 5
	}
	if game.moveDown && game.yloc <= 940 {
		game.yloc += 5
	}
	return nil
}
