package main

import (
	"image"

	"golang.org/x/image/draw"
)

func Resize(src image.Image, interpolator draw.Interpolator) *image.RGBA {
	srcRect := src.Bounds()
	srcW := srcRect.Max.X
	srcH := srcRect.Max.Y
	dstImg := image.NewRGBA(image.Rect(0, 0, srcW/scale, srcH/scale))
	dstRect := dstImg.Bounds()

	interpolator.Scale(dstImg, dstRect, src, srcRect, draw.Over, nil)

	return dstImg
}
