package lesson6

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func Draw() {
	no_color := color.Color(color.Alpha{0})
	red := color.Color(color.RGBA{200, 30, 30, 255})
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 380, 380))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{no_color}, image.ZP, draw.Src)

	for x := 0; x < 380; x++ {
		//y := x + 15
		rectImg.Set(x, 100, red)
	}

	for y := 0; y < 380; y++ {
		//x := y + 15
		rectImg.Set(190, y, green)
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
