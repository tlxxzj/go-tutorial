package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (m Image) At(x, y int) color.Color {
	v := (uint8)(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
