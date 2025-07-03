package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"sort"
	"strconv"
	"sync"

	_ "image/png"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/opentype"
)

type PixelImage struct {
	mu sync.Mutex

	image.Image
	colorCount map[color.RGBA]int
}

func NewPixelImage(src image.Image) *PixelImage {
	colorCount := make(map[color.RGBA]int)
	pi := &PixelImage{
		sync.Mutex{},
		src,
		colorCount,
	}
	pi.fillColorCountKeys()

	return pi
}

func (pi *PixelImage) fillColorCountKeys() {
	img := pi.Image

	for y := range img.Bounds().Dy() {
		for x := range img.Bounds().Dx() {
			r, g, b, a := img.At(x, y).RGBA()
			if a == 0 {
				continue
			}

			rgbaColor := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			pi.colorCount[rgbaColor] = 0
		}
	}
}

func (pi *PixelImage) GetColors() []color.RGBA {
	colors := make([]color.RGBA, 0, len(pi.colorCount))

	for k := range pi.colorCount {
		colors = append(colors, k)
	}

	sort.Slice(colors, func(i, j int) bool {
		li := 0.299*float64(colors[i].R) + 0.587*float64(colors[i].G) + 0.114*float64(colors[i].B)
		lj := 0.299*float64(colors[j].R) + 0.587*float64(colors[j].G) + 0.114*float64(colors[j].B)
		return li > lj
	})

	fmt.Println(colors)

	return colors
}

func (pi *PixelImage) IncrementColorCount(c color.RGBA) {
	pi.mu.Lock()
	pi.colorCount[c]++
	pi.mu.Unlock()
}

func (pi *PixelImage) DrawGrid(cellSize int) draw.Image {
	bounds := pi.Image.Bounds()
	result := image.NewRGBA(bounds)

	draw.Draw(result, bounds, pi.Image, bounds.Min, draw.Src)

	width, height := bounds.Dx(), bounds.Dy()

	for x := 0; x < width; x += cellSize {
		for y := range height {
			result.Set(x, y, color.Black)
		}
	}

	for y := 0; y < height; y += cellSize {
		for x := range width {
			result.Set(x, y, color.Black)
		}
	}

	return result
}

func DrawNums(src draw.Image, colors []color.RGBA, cellSize int) image.Image {
	dc := gg.NewContextForImage(src)
	palette := getColorPalette(colors)

	f, err := opentype.Parse(gomedium.TTF)
	checkError(err)

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    float64(cellSize) * 0.5,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	checkError(err)
	dc.SetFontFace(face)

	width, height := src.Bounds().Dx(), src.Bounds().Dy()
	for y := 0; y < height; y += cellSize {
		for x := 0; x < width; x += cellSize {
			cx, cy := x+cellSize/2, y+cellSize/2

			currColor := src.At(cx, cy)
			paletteIndex := palette.Index(currColor)

			r, g, b, _ := currColor.RGBA()
			if r == 0 && g == 0 && b == 0 {
				dc.SetColor(color.White)
			} else {
				dc.SetColor(color.Black)
			}

			dc.DrawStringAnchored(
				strconv.Itoa(paletteIndex+1),
				float64(cx),
				float64(cy),
				0.5, 0.5)
		}
	}

	return dc.Image()
}

func getColorPalette(colors []color.RGBA) color.Palette {
	ret := color.Palette{}

	for _, v := range colors {
		ret = append(ret, v)
	}

	return ret
}

func (pi *PixelImage) String() string {
	return fmt.Sprintf("colors :%v", pi.colorCount)
}
