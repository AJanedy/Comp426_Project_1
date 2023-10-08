package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Space Shooter")

	//New image from file returns image as image.Image (_) and ebiten.Image
	backgroundPict, _, err := ebitenutil.NewImageFromFile("background.png")
	if err != nil {
		fmt.Println("Unable to load background image:", err)
	}

	playerPict, _, err := ebitenutil.NewImageFromFile("spaceship.png")
	if err != nil {
		fmt.Println("Unable to load player image", err)
	}

	laserPict, _, err := ebitenutil.NewImageFromFile("blue_laser.png")
	if err != nil {
		fmt.Println("Unable to load laser image", err)
	}
	allLasers := make([]Shot, 0, 20)

	enemyPict, _, err := ebitenutil.NewImageFromFile("Titan.png")
	if err != nil {
		fmt.Println("Unable to load enemy ship", err)
	}
	allEnemies := make([]Enemy, 0, 5)
	for i := 0; i < cap(allEnemies); i++ {
		allEnemies = append(allEnemies, NewEnemy(928, 928, enemyPict))
	}
	const SOUND_SAMPLE_RATE = 48000
	soundContext := audio.NewContext(SOUND_SAMPLE_RATE)

	thisGameInstance := SpaceShooter{
		player: playerPict, xloc: 15, yloc: 500,
		background:   backgroundPict,
		enemyPict:    enemyPict,
		laserPict:    laserPict,
		lasers:       allLasers,
		enemies:      allEnemies,
		audioContext: soundContext,
		soundPlayer:  LoadWav("pewpew.wav", soundContext),
		score:        0,
	}
	err = ebiten.RunGame(&thisGameInstance)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}
