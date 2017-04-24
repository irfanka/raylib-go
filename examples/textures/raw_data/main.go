package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	raylib.InitWindow(screenWidth, screenHeight, "raylib [textures] example - texture from raw data")

	// NOTE: Textures MUST be loaded after Window initialization (OpenGL context is required)

	// Load RAW image data (384x512, 32bit RGBA, no file header)
	fudesumiRaw := raylib.LoadImageRaw("texture_formats/fudesumi.raw", 384, 512, raylib.UncompressedR8g8b8a8, 0)
	fudesumi := raylib.LoadTextureFromImage(fudesumiRaw) // Upload CPU (RAM) image to GPU (VRAM)
	raylib.UnloadImage(fudesumiRaw)                      // Unload CPU (RAM) image data

	// Generate a checked texture by code (1024x1024 pixels)
	width := 1024
	height := 1024

	// Dynamic memory allocation to store pixels data (Color type)
	pixels := make([]raylib.Color, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if ((x/32+y/32)/1)%2 == 0 {
				pixels[y*height+x] = raylib.Orange
			} else {
				pixels[y*height+x] = raylib.Gold
			}
		}
	}

	// Load pixels data into an image structure and create texture
	checkedIm := raylib.LoadImageEx(pixels, int32(width), int32(height))
	checked := raylib.LoadTextureFromImage(checkedIm)
	raylib.UnloadImage(checkedIm) // Unload CPU (RAM) image data

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RayWhite)

		raylib.DrawTexture(checked, screenWidth/2-checked.Width/2, screenHeight/2-checked.Height/2, raylib.Fade(raylib.White, 0.5))
		raylib.DrawTexture(fudesumi, 430, -30, raylib.White)

		raylib.DrawText("CHECKED TEXTURE ", 84, 100, 30, raylib.Brown)
		raylib.DrawText("GENERATED by CODE", 72, 164, 30, raylib.Brown)
		raylib.DrawText("and RAW IMAGE LOADING", 46, 226, 30, raylib.Brown)

		raylib.EndDrawing()
	}

	raylib.UnloadTexture(fudesumi) // Texture unloading
	raylib.UnloadTexture(checked)  // Texture unloading

	raylib.CloseWindow()
}
