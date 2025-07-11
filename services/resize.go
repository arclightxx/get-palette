package services

import (
	"fmt"
	"image"

	"github.com/arclightxx/getpalette/errors"
	"golang.org/x/image/draw"
)

const (
	minW = 32
	maxW = 1980
)

func Resize(src image.Image, interpolator draw.Interpolator, scale int) *image.RGBA {
	srcRect := src.Bounds()
	scaledW, err := applyScale(src.Bounds().Dx(), scale)
	errors.CheckError(err)
	scaledH, err := applyScale(srcRect.Dy(), scale)
	errors.CheckError(err)
	dstImg := image.NewRGBA(image.Rect(0, 0, scaledW, scaledH))
	dstRect := dstImg.Bounds()

	interpolator.Scale(dstImg, dstRect, src, srcRect, draw.Over, nil)

	return dstImg
}

func applyScale(n, s int) (int, error) {
	if s < 0 {
		ret := n / s * -1
		if ret < minW {
			err := fmt.Errorf("min width is %dpx", minW)
			return ret, err
		} else {
			return ret, nil
		}
	} else {
		ret := n * s
		if ret > maxW {
			err := fmt.Errorf("max width is %dpx", maxW)
			return ret, err
		} else {
			return ret, nil
		}
	}
}
