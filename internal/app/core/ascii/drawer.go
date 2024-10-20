package ascii

import (
	"image"
	"image/color"
	"io"
)

const (
	ramp    = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
	newLine = "\n"
)

func Draw(w io.Writer, img image.Image) {
	max := img.Bounds().Max
	scaleX, scaleY := 2, 1

	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			c := avgPixel(img, x, y, scaleX, scaleY)
			w.Write([]byte(string(ramp[len(ramp)*c/65536])))
		}

		w.Write([]byte(newLine))
	}
}

func avgPixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += grayscale(img.At(i, j))
			cnt++
		}
	}

	return sum / cnt
}

func grayscale(c color.Color) int {
	r, g, b, _ := c.RGBA()

	return int(0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b))
}
