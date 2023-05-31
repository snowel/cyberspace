package main

import (
	//"math"
 	//tea "github.com/charmbracelet/bubbletea"
	//dgf "github.com/snowel/dgf"
)

// For a set of coordinates, give the DGF intersectoins ID
// Negative value means it's not a valid one.
func TermCoordToIntersect(x, y, xOff, yOff, size int) int {
  yCoor := y - yOff + 1
  xTemp := x - xOff + 1
	if xTemp > size * 2 { return -1 } // cursor is of the board to the rigth
	if yCoor > size { return -1 } // cursor is off the board to bottom
	if x < xOff { return -1 } // curosor is off the board to the left
	if y < yOff { return -1 } // above the board
	var xCoor int
  if xTemp % 2 != 0 {
		xCoor = (xTemp + 1 ) / 2
  } else {
		return -1
  }

  return xCoor + (yCoor - 1) * size
}

/* -- Boordinate practice -- */
