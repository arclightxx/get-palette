package main

import "flag"

const (
	minW = 32
	maxW = 1980
)

type Config struct {
	inputPath     string
	outputPath    string
	scale         int
	cellSize      int
	gridThickness int
}

func NewConfig() *Config {
	c := &Config{}
	flag.StringVar(&c.inputPath, "input", "", "Path to input file")
	flag.StringVar(&c.outputPath, "output", "", "Path to output file")
	flag.IntVar(&c.scale, "scale", 0, "Scale < 0 to shrink the image\nScale > 0 to grow the image")
	flag.IntVar(&c.cellSize, "cellsize", 2, "Target cell size in pixels")
	flag.IntVar(&c.gridThickness, "grid", 1, "Grid line thickness")

	return c
}
