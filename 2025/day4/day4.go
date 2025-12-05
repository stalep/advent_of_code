package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type PaperGrid struct {
	Grid [][]Paper
}

type Paper struct {
	X int
	Y int
	HasPaper bool
	MarkedForRemoval bool
}

var PERIOD byte = "@"[0]

func CreateNewPaperGrid(inputMap []string) [][]Paper {
	PaperGrid := make([][]Paper, len(inputMap))
	// all strings should be equal length, lets get the length of the first one
	width := len(inputMap[0])
	for y := 0; y < len(inputMap); y++ {
		PaperGrid[y] = make([]Paper, width)
		for x := 0; x < width; x++ {
			paper := inputMap[y][x] == PERIOD
			PaperGrid[y][x] = Paper{X: x, Y: y, HasPaper: paper, MarkedForRemoval: false}
		}
	}

	return PaperGrid
}

func ReadFileGeneratePaperGrid(filename string) [][]Paper {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %s", err)
		return nil
	}

	return CreateNewPaperGrid(lines)
}

func FindNumberOfAccessablePapers(grid [][]Paper) int {
	sum := 0
	for y, papers := range grid {
		for x, paper := range papers {
			if(paper.HasPaper) {
				adjacentPapers := FindAdjacentPaperRolls(grid, paper)
				//since we're not checking for self, we check for < 5
				if adjacentPapers < 5 {
					grid[y][x].MarkedForRemoval = true
					sum++
				}
			}
		}
	}
	return sum
}

func FindAdjacentPaperRolls(grid [][]Paper, paper Paper) int {
	maxY := len(grid)
	maxX := len(grid[0])

	adjacentPapers := 0
	for x := paper.X - 1; x <= paper.X + 1; x++ {
		if (x >= 0 && x < maxX) {
			for y := paper.Y - 1; y <= paper.Y + 1; y++ {
				if (y >= 0 && y < maxY) {
					if grid[y][x].HasPaper {
						adjacentPapers++
					}
				}
			}
		}
	}
	return adjacentPapers
}

func ClearPaperRollsMarkedForRemoval(grid [][]Paper) {
	for y, papers := range grid {
		for x, paper := range papers {
			if(paper.MarkedForRemoval) {
				grid[y][x].HasPaper = false
			}
		}
	}
}

func main() {

	paperGrid := ReadFileGeneratePaperGrid("input")

	numAccessablePapers := FindNumberOfAccessablePapers(paperGrid)

	fmt.Printf("Task1, found number of accessable paper rolls %d\n", numAccessablePapers)

	got := 1
	totalRemovedRolls := 0

	for i := 0; i < got; i = 0 {
		got = FindNumberOfAccessablePapers(paperGrid)
		totalRemovedRolls += got
		ClearPaperRollsMarkedForRemoval(paperGrid)
	}

	fmt.Printf("Task2, total number of removed paper rolls %d\n", totalRemovedRolls)
}
