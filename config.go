package main

import "flag"

const (
	minW        = 32
	maxW        = 1980
)

type Config struct {
	inputPath     string
	outoutPath    string
	cellSize      int
	gridThickness int
}

func NewConfig() *Config {
	c := &Config{}
	flag.StringVar(&c.inputPath, "input", "./img", "Path to input file")
	flag.StringVar(&c.outoutPath, "output", "./out", "Path to output file")
	flag.IntVar(&c.cellSize, "cellsize", 1, "Target cell size in pixels")
	flag.IntVar(&c.gridThickness, "grid", 1, "Grid line thickness")

	return c
}
