package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"sync"

	"golang.org/x/image/draw"
)

func main() {
	var wg sync.WaitGroup

	cfg := NewConfig()
	flag.Parse()
	pathSlice := ParsePath(cfg.inputPath)

	for _, path := range pathSlice {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()

			f := OpenFile(path)
			defer f.Close()

			img, _, err := image.Decode(f)
			checkError(err)

			if cfg.scale != 0 {
				img = Resize(img, draw.NearestNeighbor, cfg.scale)
			}
			pixelImage := NewPixelImage(img)

			grid := pixelImage.DrawGrid(cfg.cellSize)

			res := DrawNums(grid, pixelImage.GetColors(), cfg.cellSize)

			outDir := "./out/"
			outName := fmt.Sprintf("schema-%s", ParseName(path))

			out, err := os.Create(outDir + outName)
			checkError(err)
			defer out.Close()

			err = png.Encode(out, res)
			checkError(err)
		}(path)
	}

	wg.Wait()
}

func OpenFile(path string) *os.File {
	f, err := os.Open(path)
	checkError(err)
	return f
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
