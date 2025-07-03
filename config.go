package main

const (
	minW        = 32
	maxW        = 1980
	minCellSize = 1
)

type Config struct {
	FilePath string
	Scale    int // Scale < 0 to shrink the image; Scale > 0 to grow the image
	CellSize int
}
