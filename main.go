package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerLookDirection int

const (
	LookUp PlayerLookDirection = iota
	LookRight
	LookDown
	LookLeft
)

type GameAsset struct {
	Name              string
	Image             rl.Image
	Texture           rl.Texture2D
	BoundingBox       rl.BoundingBox
	X                 int
	Y                 int
	Look              PlayerLookDirection
	Hp                uint
	Xp                uint
	MovementSpeed     uint
	AnimationFrame    uint
	AnimationFrameMax uint
}

func NewGameAsset(name string, assetFileName string, posX int, posY int, flip bool) GameAsset {
	sprite := rl.LoadImage(assetFileName)

	if flip {
		rl.ImageFlipHorizontal(sprite)
	}

	texture := rl.LoadTextureFromImage(sprite)

	return GameAsset{
		Name:              name,
		Image:             *sprite,
		Texture:           texture,
		X:                 posX,
		Y:                 posY,
		Look:              LookRight,
		Hp:                100,
		Xp:                0,
		MovementSpeed:     3,
		AnimationFrame:    0,
		AnimationFrameMax: 5,
	}
}

type World struct {
	Player   GameAsset
	Monsters []GameAsset
	Input    PlayerInput
}

type PlayerInput struct {
	Right            bool
	Left             bool
	Up               bool
	Down             bool
	ToggleFullscreen bool
	Movement         bool
}

var globalFrameCounter uint

func main() {
	var screenWidth int32 = 800
	var screenHeight int32 = 600
	var spriteScale float32 = 3
	var frameSpeed uint = 8

	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagBorderlessWindowedMode)
	rl.InitWindow(screenWidth, screenHeight, "raylibtest")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerAsset := NewGameAsset("Dude", "assets/Heroes/Knight/Run/Run-Sheet.png", 0, 0, false)
	//playerAsset := NewGameAsset("Dude", "assets/Heroes/Wizzard/Run/Run-Sheet.png", 0, 0, false)
	grassAsset := NewGameAsset("Grass", "assets/Environment/Green Woods/Assets/Tiles.png", 0, 0, false)

	for !rl.WindowShouldClose() {
		// Input Stuff
		input := ProcessInput()
		UpdatePlayerInput(&playerAsset, &input)

		// Draw stuff
		rl.BeginDrawing()

		rl.ClearBackground(rl.DarkGreen)

		grassAsset.DrawTile(spriteScale)
		playerAsset.DrawAnimation(spriteScale, 90, 64)

		globalFrameCounter++
		if globalFrameCounter >= 60/frameSpeed {
			if input.Movement {
				playerAsset.AdvancePlayerAnimationFrame()
			}
			globalFrameCounter = 0
		}

		rl.EndDrawing()
	}
}

// Load sprite
// Horizontally flip the texture by negating the width value
func (asset *GameAsset) DrawAsset(spriteScale float32) {
	var flipTexture float32 = -1.0
	if asset.Look == LookRight {
		flipTexture = 1.0
	}

	rl.DrawTexturePro(asset.Texture,
		rl.Rectangle{X: 0,
			Y:      0,
			Height: float32(asset.Texture.Height),
			Width:  flipTexture * float32(asset.Texture.Width)},
		rl.Rectangle{
			X:      float32(asset.X),
			Y:      float32(asset.Y),
			Height: float32(asset.Texture.Height) * spriteScale,
			Width:  float32(asset.Texture.Width) * spriteScale},
		rl.Vector2{X: 0, Y: 0}, 0.0, rl.White)

}

func (asset *GameAsset) DrawAnimation(spriteScale float32, offset uint, spriteSize uint) {
	var flipTexture float32 = -1.0
	if asset.Look == LookRight {
		flipTexture = 1.0
	}

	rl.DrawTexturePro(asset.Texture,
		rl.Rectangle{X: float32(spriteSize * asset.AnimationFrame),
			Y:      float32(offset),
			Height: float32(spriteSize),
			Width:  float32(spriteSize) * flipTexture},
		rl.Rectangle{
			X:      float32(asset.X),
			Y:      float32(asset.Y),
			Height: float32(spriteSize) * spriteScale,
			Width:  float32(spriteSize) * spriteScale},
		rl.Vector2{X: 0, Y: 0}, 0.0, rl.White)
}

// Load sprite
// Horizontally flip the texture by negating the width value
func (asset *GameAsset) DrawTile(spriteScale float32) {
	var flipTexture float32 = -1.0
	if asset.Look == LookRight {
		flipTexture = 1.0
	}

	rl.DrawTexturePro(asset.Texture,
		rl.Rectangle{X: 0,
			Y:      0,
			Height: float32(asset.Texture.Height),
			Width:  flipTexture * float32(asset.Texture.Width)},
		rl.Rectangle{
			X:      float32(asset.X),
			Y:      float32(asset.Y),
			Height: float32(asset.Texture.Height) * spriteScale,
			Width:  float32(asset.Texture.Width) * spriteScale},
		rl.Vector2{X: 0, Y: 0}, 0.0, rl.White)
}

func ProcessInput() PlayerInput {
	return PlayerInput{
		Up:               rl.IsKeyPressed(rl.KeyW) || rl.IsKeyDown(rl.KeyW),
		Down:             rl.IsKeyPressed(rl.KeyS) || rl.IsKeyDown(rl.KeyS),
		Left:             rl.IsKeyPressed(rl.KeyA) || rl.IsKeyDown(rl.KeyA),
		Right:            rl.IsKeyPressed(rl.KeyD) || rl.IsKeyDown(rl.KeyD),
		ToggleFullscreen: rl.IsKeyPressed(rl.KeyF11),
	}
}

// Update Player location with input
func UpdatePlayerInput(player *GameAsset, input *PlayerInput) {
	input.Movement = false

	if input.Right {
		player.X = player.X + 1*int(player.MovementSpeed)
		input.Movement = true
		if player.Look == LookLeft {
			player.Look = LookRight
		}
	}

	if input.Left {
		player.X = player.X - 1*int(player.MovementSpeed)
		input.Movement = true
		if player.Look == LookRight {
			player.Look = LookLeft
		}
	}

	if input.Up {
		player.Y = player.Y - 1*int(player.MovementSpeed)
		input.Movement = true
	}

	if input.Down {
		player.Y = player.Y + 1*int(player.MovementSpeed)
		input.Movement = true
	}
}

func (asset *GameAsset) AdvancePlayerAnimationFrame() {
	asset.AnimationFrame++
	if asset.AnimationFrame > asset.AnimationFrameMax {
		asset.AnimationFrame = 0
	}
}
