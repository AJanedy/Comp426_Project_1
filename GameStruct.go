package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
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
	soundPlayer     *audio.Player
	lasers          []LaserStruct
	enemies         []EnemyStruct
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

type GameSoundsStruct struct {
	audioContext *audio.Context
	soundPlayer  *audio.Player
}
