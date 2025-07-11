package entities

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"sort"
	"strconv"
	"sync"

	_ "image/png"

	"github.com/arclightxx/getpalette/errors"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/opentype"
)

type PixelImage struct {
	mu sync.Mutex

	*image.RGBA
	colorCount map[color.RGBA]int
	width      int
	height     int
}

func NewPixelImage(src *image.RGBA) *PixelImage {
	colorCount := make(map[color.RGBA]int)
	pi := &PixelImage{
		sync.Mutex{},
		src,
		colorCount,
		src.Bounds().Dx(),
		src.Bounds().Dy(),
	}
	pi.fillColorCountKeys()

	return pi
}

func (pi *PixelImage) fillColorCountKeys() {
	img := pi.RGBA

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
		li := colors[i].R + colors[i].G + colors[i].B
		lj := colors[j].R + colors[j].G + colors[j].B
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

func (pi *PixelImage) DrawGrid(cellSize int) {
	for x := 0; x < pi.width; x += cellSize {
		for y := range pi.height {
			pi.RGBA.Set(x, y, color.Black)
		}
	}

	for y := 0; y < pi.height; y += cellSize {
		for x := range pi.width {
			pi.RGBA.Set(x, y, color.Black)
		}
	}
}

func (pi *PixelImage) DrawNums(cellSize int) {
	dc := gg.NewContextForImage(pi.RGBA)
	palette := getColorPalette(pi.GetColors())

	f, err := opentype.Parse(gomedium.TTF)
	errors.CheckError(err)

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    float64(cellSize) * 0.5,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	errors.CheckError(err)
	dc.SetFontFace(face)

	for y := 0; y < pi.height; y += cellSize {
		for x := 0; x < pi.width; x += cellSize {
			cx, cy := x+cellSize/2, y+cellSize/2

			currColor := pi.RGBA.At(cx, cy)
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

	draw.Draw(pi.RGBA, pi.Bounds(), dc.Image(), pi.Bounds().Min, draw.Src)
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
