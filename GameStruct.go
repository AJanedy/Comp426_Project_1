package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type SpaceShooter struct {
	player          *ebiten.Image
	xloc            int
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
	lasers          []Shot
	enemies         []Enemy
}

type Shot struct {
	laser  *ebiten.Image
	xLoc   int
	yLoc   int
	deltaX int
}

type Enemy struct {
	enemy *ebiten.Image
	xLoc  int
	yLoc  int
	//deltaX int
}

type GameSounds struct {
	audioContext *audio.Context
	soundPlayer  *audio.Player
}
