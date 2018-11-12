package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GO doesn't have a concept of exception
var LocationOccupied = fmt.Errorf("LocationOccupied")

// PHP doesn't have a concept of nested classes (or even classes)
var cachedDirections = [8][2]int{
	{-1, 1}, {0, 1}, {1, 1},    // above
	{-1, 0}, {1, 0},            // sides
	{-1, -1}, {0, -1}, {1, -1}, // below
}

type Cell struct {
	x, y       int
	alive      bool
	next_state int
	neighbours []*Cell
}

func (c *Cell) to_char() rune {
	if c.alive {
		return 'o'
	} else {
		return ' '
	}
}

type World struct {
	width, height int
	tick          int

	cells [World_Width + 1][World_Height + 1]*Cell
}

// Like constructor
func newWorld(width, height int) *World {
	w := &World{width: width, height: height, cells: [World_Width + 1][World_Height + 1]*Cell{}}

	w.populate_cells()
	w.prepopulate_neighbours()
	return w
}

func (w *World) _tick() {
	var cell *Cell
	// First determine the action for all cells
	for y := 0; y <= w.height; y++ {
		for x := 0; x <= w.width; x++ {
			cell = w.cells[x][y]
			alive_neighbours := w.alive_neighbours_around(cell)

			if !cell.alive && alive_neighbours == 3 {
				cell.next_state = 1
			} else if alive_neighbours < 2 || alive_neighbours > 3 {
				cell.next_state = 0
			}
		}
	}

	// Then execute the determined action for all cells
	for y := 0; y <= w.height; y++ {
		for x := 0; x <= w.width; x++ {
			cell = w.cells[x][y]
			if cell.next_state == 1 {
				cell.alive = true
			} else if cell.next_state == 0 {
				cell.alive = false
			}
		}
	}
	w.tick++
}

// Implement first using string concatenation. Then implement any
// special string builders, and use whatever runs the fastest
func (w *World) render() string {
	// The following works but is slower
	// rendering := ""
	// for y := 0; y <= w.height; y++ {
	// 	for x := 0; x <= w.width; x++ {
	// 		cell, _ := w.cell_at(x, y)
	//		rendering += string(cell.to_char())
	//	}
	//	rendering += "\n"
	// }
	// return rendering
	var cell *Cell
	var rendering = strings.Builder{}
	for y := 0; y <= w.height; y++ {
		for x := 0; x <= w.width; x++ {
			cell, _ = w.cell_at(x, y)
			rendering.WriteRune(cell.to_char())
		}
		rendering.WriteRune('\n')
	}
	return rendering.String()
}

func (w *World) populate_cells() {
	rand.Seed(time.Now().UnixNano())

	for y := 0; y <= w.height; y++ {
		for x := 0; x <= w.width; x++ {
			alive := rand.Intn(100) <= 20
			w.add_cell(x, y, alive)
		}
	}
}

func (w *World) prepopulate_neighbours() {
	var cell *Cell
	for y := 0; y <= w.height; y++ {
		for x := 0; x <= w.width; x++ {
			cell = w.cells[x][y]
			w.neighbours_around(cell)
		}
	}
}

func (w *World) add_cell(x, y int, alive bool) *Cell {
	if _, empty := w.cell_at(x, y); empty {
		panic(LocationOccupied)
	}

	c := &Cell{x: x, y: y, alive: alive}

	w.cells[x][y] = c
	return c
}

func (w *World) cell_at(x, y int) (*Cell, bool) {
	if x >= 0 && y >= 0 && x <= w.width && y <= w.height {
		return w.cells[x][y], false
	}

	return nil, true
}

func (w *World) neighbours_around(cell *Cell) []*Cell {
	if cell.neighbours == nil {
		for _, set := range cachedDirections {
			neighbour, empty := w.cell_at(cell.x+set[0], cell.y+set[1])

			if !empty {
				cell.neighbours = append(cell.neighbours, neighbour)
			}
		}
	}
	return cell.neighbours
}

// Implement first using filter/lambda if available. Then implement
// foreach and for. Retain whatever implementation runs the fastest
func (w *World) alive_neighbours_around(cell *Cell) int {
	alive_neighbours := 0
	for _, neighbour := range w.neighbours_around(cell) {
		if neighbour.alive {
			alive_neighbours++
		}
	}
	return alive_neighbours
}
