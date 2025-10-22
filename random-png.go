package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	w, h := 256, 256
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r := uint8(rand.Intn(40))
			g := uint8(rand.Intn(120))
			b := uint8(rand.Intn(21))
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	f, _ := os.Create("random_valid.png")
	defer f.Close()
	png.Encode(f, img) // menghasilkan PNG valid yang pasti bisa dibuka
}
