package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

func (game *SpaceShooter) Draw(screen *ebiten.Image) {
	drawOps := ebiten.DrawImageOptions{}
	const repeat = 3
	backgroundWidth := game.background.Bounds().Dx()
	for count := 0; count < repeat; count += 1 {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(float64(backgroundWidth*count),
			float64(-1000))
		drawOps.GeoM.Translate(float64(game.backgroundXView), 0)
		screen.DrawImage(game.background, &drawOps)
	}
	drawOps.GeoM.Reset()
	drawOps.GeoM.Translate(float64(game.xloc), float64(game.yloc))
	screen.DrawImage(game.player, &drawOps)
	for _, laser := range game.lasers {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(float64(game.xloc), float64(game.yloc))
		screen.DrawImage(laser.laser, &drawOps)
	}
	for _, enemy := range game.enemies {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(float64(enemy.xLoc), float64(enemy.yLoc))
		screen.DrawImage(enemy.enemy, &drawOps)
	}
	//font, err := truetype.Parse(gomedium.TTF)
	//if err != nil {
	//	fmt.Println("Unable to load laser image", err)
	//}
	//DrawCenteredText(screen, font, "Score: "+strconv.Itoa(game.score), 500, 10)
}

func (game SpaceShooter) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
