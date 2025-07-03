package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"sort"

	_ "image/png"
)

type PixelImage struct {
	image.Image
	colorSet map[color.RGBA]bool
}

func NewPixelImage(path string) (*PixelImage, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	srcImg, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	colorSet := make(map[color.RGBA]bool)
	pi := &PixelImage{
		srcImg,
		colorSet,
	}
	pi.fillColorSet()

	return pi, nil
}

func (pi *PixelImage) fillColorSet() {
	img := pi.Image
	rect := img.Bounds()
	maxX := rect.Max.X
	maxY := rect.Max.Y
	minX := rect.Min.X
	minY := rect.Min.Y

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rgbaColor := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}

			pi.colorSet[rgbaColor] = true
		}
	}
}

func (pi *PixelImage) GetColorSet() []color.RGBA {
	ret := make([]color.RGBA, 0)

	for k, v := range pi.colorSet {
		if v {
			ret = append(ret, k)
		}
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].R < ret[j].R || ret[i].G < ret[j].G || ret[i].B < ret[j].B
	})

	return ret
}

func (pi *PixelImage) String() string {
	return fmt.Sprintf("color set :%v", pi.colorSet)
}
