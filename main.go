package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
	"math/rand"
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
	allLasers := make([]LaserStruct, 0, 20)

	enemyPict, _, err := ebitenutil.NewImageFromFile("Titan.png")
	if err != nil {
		fmt.Println("Unable to load enemy ship", err)
	}
	allEnemies := make([]EnemyStruct, 0, 5)
	for i := 0; i < cap(allEnemies); i++ {
		allEnemies = append(allEnemies, NewEnemy(rand.Intn(928)+928, rand.Intn(928), enemyPict))
	}
	const SOUND_SAMPLE_RATE = 48000
	soundContext := audio.NewContext(SOUND_SAMPLE_RATE)

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	scoreFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	thisGameInstance := GameStruct{
		player: playerPict, xLoc: 15, yloc: 500,
		background:   backgroundPict,
		enemyPict:    enemyPict,
		laserPict:    laserPict,
		lasers:       allLasers,
		enemies:      allEnemies,
		audioContext: soundContext,
		pewPewSound:  LoadWav("pewpew.wav", soundContext),
		popSound:     LoadWav("pop.wav", soundContext),
		score:        0,
		typeface:     scoreFont,
	}
	err = ebiten.RunGame(&thisGameInstance)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}
