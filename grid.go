package main

import (
	"image/color"
	"image/draw"
)

func DrawGrid(src draw.Image) {
	for y := range src.Bounds().Max.Y {
		for x := range src.Bounds().Max.X {
			if x%scale == 0 {
				src.Set(x, y, color.Black)
				src.Set(x+1, y, color.Black)
			}
			if y%scale == 0 {
				src.Set(x, y, color.Black)
				src.Set(x, y+1, color.Black)
			}
		}
	}
}
