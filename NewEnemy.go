package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func NewEnemy(MaxWidth int, MaxHeight int, image *ebiten.Image) EnemyStruct {
	return EnemyStruct{
		enemy: image,
		xLoc:  MaxWidth,
		yLoc:  MaxHeight,
	}
}
