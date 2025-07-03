package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"sync"

	"golang.org/x/image/draw"
)

const (
	scale = 4
)

var (
	pathSlice = []string{
		// "/home/arc/go-projects/pixelart-to-grid/img/clubs.png",
		// "/home/arc/go-projects/pixelart-to-grid/img/diamonds.png",
		// "/home/arc/go-projects/pixelart-to-grid/img/hearts.png",
		// "/home/arc/go-projects/pixelart-to-grid/img/spades.png",
		"./img/milk.png",
	}
	colorCount = make(map[color.Color]int)
	mu         sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	for i, path := range pathSlice {
		wg.Add(1)
		go func(id int, p string) {
			defer wg.Done()

			img, err := NewPixelImage(p)
			checkError(err)

			fmt.Println(p, img)
			resizedImg := Resize(img, draw.NearestNeighbor)

			DrawGrid(resizedImg)

			res := DrawNums(resizedImg, img.GetColorSet())

			outPath := fmt.Sprintf("./out/%d.png", i)
			out, err := os.Create(outPath)
			checkError(err)
			defer out.Close()

			err = png.Encode(out, res)
			checkError(err)
		}(i, path)
	}

	wg.Wait()

	fmt.Println(colorCount)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
