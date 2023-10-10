package main

import "github.com/hajimehoshi/ebiten/v2"

func NewLaser(x int, y int, image *ebiten.Image) LaserStruct {
	return LaserStruct{
		laser: image,
		xLoc:  x,
		yLoc:  y,
		//deltaX: Delta,
	}
}
