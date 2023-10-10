package main

import (
	"github.com/co0p/tankism/lib/collision"
)

func CheckCollision(enemy EnemyStruct, laser LaserStruct, game *GameStruct) bool {
	laserBounds := collision.BoundingBox{
		X:      float64(laser.xLoc),
		Y:      float64(laser.yLoc),
		Width:  float64(laser.laser.Bounds().Dx()),
		Height: float64(laser.laser.Bounds().Dy()),
	}
	enemyBounds := collision.BoundingBox{
		X:      float64(enemy.xLoc),
		Y:      float64(enemy.yLoc),
		Width:  float64(enemy.enemy.Bounds().Dx()),
		Height: float64(enemy.enemy.Bounds().Dy()),
	}
	if collision.AABBCollision(laserBounds, enemyBounds) {
		return true
	}
	return false
}

func RemoveIndex(enemies []EnemyStruct, index int) []EnemyStruct {
	copied := []EnemyStruct{}
	copied = append(copied, enemies[:index]...)
	return append(copied, enemies[index+1:]...)
}
