package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"sync"

	"github.com/arclightxx/getpalette/entities"
	"github.com/arclightxx/getpalette/errors"
	"github.com/arclightxx/getpalette/services"
	"golang.org/x/image/draw"
)

func main() {
	var wg sync.WaitGroup

	cfg := NewConfig()
	flag.Parse()
	pathSlice := services.ParsePath(cfg.inputPath)

	for _, path := range pathSlice {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()

			f := OpenFile(path)
			defer f.Close()

			img, _, err := image.Decode(f)
			errors.CheckError(err)
			rgba := image.NewRGBA(img.Bounds())
			draw.Draw(rgba, img.Bounds(), img, img.Bounds().Min, draw.Src)

			if cfg.scale != 0 {
				rgba = services.Resize(img, draw.NearestNeighbor, cfg.scale)
			}
			pixelImage := entities.NewPixelImage(rgba)

			pixelImage.DrawGrid(cfg.cellSize)

			pixelImage.DrawNums(cfg.cellSize)

			outDir := "./out/"
			outName := fmt.Sprintf("schema-%s", services.ParseName(path))

			out, err := os.Create(outDir + outName)
			errors.CheckError(err)
			defer out.Close()

			err = png.Encode(out, pixelImage)
			errors.CheckError(err)
		}(path)
	}

	wg.Wait()
}

func OpenFile(path string) *os.File {
	f, err := os.Open(path)
	errors.CheckError(err)
	return f
}
