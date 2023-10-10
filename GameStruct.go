package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"golang.org/x/image/font"
)

type GameStruct struct {
	player          *ebiten.Image
	xLoc            int
	yloc            int
	background      *ebiten.Image
	enemyPict       *ebiten.Image
	laserPict       *ebiten.Image
	moveUp          bool
	moveDown        bool
	score           int
	backgroundXView int
	audioContext    *audio.Context
	pewPewSound     *audio.Player
	popSound        *audio.Player
	lasers          []LaserStruct
	enemies         []EnemyStruct
	typeface        font.Face
}

type LaserStruct struct {
	laser *ebiten.Image
	xLoc  int
	yLoc  int
	//deltaX int
}

type EnemyStruct struct {
	enemy *ebiten.Image
	xLoc  int
	yLoc  int
	//deltaX int
}
