package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
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

	thisGameInstance := SpaceShooter{
		player: playerPict, xloc: 15, yloc: 500,
		background: backgroundPict,
		laser:      laserPict,
	}
	err = ebiten.RunGame(&thisGameInstance)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}
