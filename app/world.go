package app

import (
	"fmt"
	"math/rand"
)

const (
	SleepIteration      = 100
	ClearScreenSequence = "\033c\x0c"
	//	MoveCursorAndClearScreenSequence move cursor and clear screen in linux terminal.
	//  	[1;1H - this attribute move cursor ;
	//		\033[2J - this attribute clear screen ;
	//
	//	In the linux terminal you may use terminal commands to move your cursor, such as
	//
	//	printf("\033[8;5Hhello"); // Move to (8, 5) and output hello
	//
	//	other similar commands:
	//
	// 	printf("\033[XA"); // Move up X lines;
	// 	printf("\033[XB"); // Move down X lines;
	// 	printf("\033[XC"); // Move right X column;
	// 	printf("\033[XD"); // Move left X column;
	// 	printf("\033[2J"); // Clear screen
	// 	Keep in mind that this is not a standardised solution, and therefore your code will not be platform independent.
	MoveCursorAndClearScreenSequence = "\033[1;1H\033[2J"
	// yellowBlock creates new block with Yellow color. Ypu can choose any other colors:
	//	 Background colors:
	//	On_Black="\[\033[40m\]"       # Black
	//	On_Red="\[\033[41m\]"         # Red
	//	On_Green="\[\033[42m\]"       # Green
	//	On_Yellow="\[\033[43m\]"      # Yellow
	//	On_Blue="\[\033[44m\]"        # Blue
	//	On_Purple="\[\033[45m\]"      # Purple
	//	On_Cyan="\[\033[46m\]"        # Cyan
	//	On_White="\[\033[47m\]"       # White
	yellowBlock = "\033[43m  "
	blueBlock   = "\033[44m  "
	Reset       = "\033[0m"
)

type Universe [][]bool

// World represents our Universe, which will contain a two-dimensional field of cells.
// With a boolean type, each cell will be either alive (true) or dead (false).
type World struct {
	universe Universe
	width    int
	height   int
}

// MakeWorld creates new Universe with row height 'height' and column 'width' per row.
// Function returns empty world.
func MakeWorld(width, height int) (*World, error) {
	if width <= 0 && height <= 0 {
		return nil, ErrNotValidWidthAndHeight
	}
	if width <= 0 {
		return nil, ErrWidthNotValid
	}
	if height <= 0 {
		return nil, ErrHeightNotValid
	}

	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return &World{
		universe: u,
		width:    width,
		height:   height,
	}, nil
}

// Display prints to screen
func (w World) Display() {
	for _, row := range w.universe {
		for _, cell := range row {
			switch {
			case cell:
				fmt.Print(yellowBlock) // yellow
			default:
				fmt.Print(blueBlock) // blue
			}
		}
		fmt.Println(Reset)
	}
}

// Seed randomly places approximately 25% of alive cells with value 'true'.
func (w *World) Seed() {
	for _, row := range w.universe {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

// Alive determines whether a cell is alive or dead. Method just looks at the cage in the World slice.
// If the boolean value is true, then the cell is alive.
// If the coordinates are outside the universe, we return to the beginning.
func (w *World) Alive(x, y int) bool {
	y = (w.height + y) % w.height
	x = (w.width + x) % w.width
	return w.universe[y][x]
}

// Neighbors a method, which count alive neighbors for the specified cell, from 0 to 8.
// Instead of accessing the universe data directly, use the Alive method to make the universe return to the beginning.
func (w *World) Neighbors(x, y int) int {
	var neighbors int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x { // we count adjoining neighbors and not the question cell.
				continue
			}
			if w.Alive(j, i) {
				neighbors++
			}
		}
	}
	return neighbors
}

// NextState determines if a cell has two, three, or more neighbors.
//	This method shows: should current cell be alive at the next generation or not
//	Also this method implement rules of task:
//		1. Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
//		2. Any live cell with two or three live neighbors lives on to the next generation.
//		3. Any live cell with more than three live neighbors dies, as if by overcrowding.
//		4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction. +
func (w *World) NextState(neighbours int, alive bool) bool {
	if neighbours < 4 && neighbours > 1 && alive {
		return true
	} else if neighbours == 3 && !alive {
		return true
	} else {
		return false
	}
}

// Step updates the state of the next world (b) from current world (a).
// To complete the simulation, you need to go through every cell in the World and determine what the Next state
// should be.
func Step(a, b *World) {
	for i := 0; i < a.height; i++ {
		for j := 0; j < a.width; j++ {
			b.universe[i][j] = a.NextState(a.Neighbors(j, i), a.Alive(j, i))
		}
	}
}
