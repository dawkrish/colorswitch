package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)


const (
	ModeTitle Mode = iota
	ModeGame
	ModeOver

	screenWidth = 800
	screenHeight = 600
)

var (
	gameMode Mode = ModeTitle
	playerBall Ball
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	playerBall.cntrX = screenWidth/2
	playerBall.cntrY = screenHeight/2
	playerBall.radius = 50
	playerBall.color  = rl.Black
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// INPUT-START
		switch gameMode {
		case ModeTitle:
			if rl.IsKeyPressed(rl.KeySpace) {
				gameMode = ModeGame
			}
		case ModeGame:
			if rl.IsKeyPressed(rl.KeyR) {
				gameMode = ModeOver
			}
		case ModeOver:
			if rl.IsKeyPressed(rl.KeyEnter) {
				gameMode = ModeGame
			}
		}
		// INPUT-END

		// UPDATE-START

		// UPDATE-END

		// DRAW-START
		rl.BeginDrawing()
		
		switch gameMode {
		case ModeTitle:
			rl.ClearBackground(rl.Lime)
			rl.DrawText("Title Screen", 100, 150, 20, rl.RayWhite)
			rl.DrawText("Press Space to START", 100, 350, 20, rl.RayWhite)
		case ModeGame:
			rl.ClearBackground(rl.Beige)
			rl.DrawText("Game Screen", 100, 150, 20, rl.RayWhite)
			rl.DrawCircle(playerBall.cntrX, playerBall.cntrY, playerBall.radius, playerBall.color)
			rl.DrawText("Press R to END", 100, 350, 20, rl.RayWhite)
		case ModeOver:
			rl.ClearBackground(rl.Red)
			rl.DrawText("Game Over", 100, 150, 20, rl.RayWhite)
			rl.DrawText("Press ENTER to restart", 100, 350, 20, rl.RayWhite)
		}

		rl.EndDrawing()
		// DRAW-END

	}
}
