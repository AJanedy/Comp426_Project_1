package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type SpaceShooter struct {
	player          *ebiten.Image
	laser           *ebiten.Image
	xloc            int
	yloc            int
	moveUp          bool
	moveDown        bool
	score           int
	backgroundXView int
	background      *ebiten.Image
}

type Shot struct {
	xLoc   int
	yLoc   int
	deltaX int
}

type Enemy struct {
	enemy  *ebiten.Image
	xLoc   int
	yLoc   int
	deltaX int
}

type gameSounds struct {
	audioContext *audio.Context
	soundPlayer  *audio.Player
	counter      int
}
