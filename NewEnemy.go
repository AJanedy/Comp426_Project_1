package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

func NewEnemy(MaxWidth int, MaxHeight int, image *ebiten.Image) Enemy {
	return Enemy{
		enemy: image,
		xLoc:  rand.Intn(MaxWidth),
		yLoc:  rand.Intn(MaxHeight),
	}
}
