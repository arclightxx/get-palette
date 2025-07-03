package main

import (
	"image"
	"image/color"
	"log"
	"strconv"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/opentype"
)

func  DrawNums(src image.Image, colors []color.RGBA) image.Image {
	dc := gg.NewContextForImage(src)
	palette := getColorPalette(colors)

	// Загрузка шрифта с проверкой ошибок
	f, err := opentype.Parse(gomedium.TTF)
	if err != nil {
		log.Fatal("Font parsing error:", err)
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    14, // Уменьшенный размер для теста
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("Face creation error:", err)
	}
	dc.SetFontFace(face)

	width, height := src.Bounds().Max.X, src.Bounds().Max.Y
	for y := 0; y < height; y += scale {
		for x := 0; x < width; x += scale {
			// Берем центральную точку клетки
			cx, cy := x+scale/2, y+scale/2

			// Получаем цвет и индекс в палитре
			currColor := src.At(cx, cy)
			paletteIndex := palette.Index(currColor)

			mu.Lock()
			colorCount[currColor]++
			mu.Unlock()

			// Определяем цвет текста (инвертированный)
			r, g, b, _ := currColor.RGBA()
			if r == 0 && g == 0 && b == 0 {
				dc.SetColor(color.White)
			} else {
				dc.SetColor(color.Black)
			}

			// Рисуем цифру
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
