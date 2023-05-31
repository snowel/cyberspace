package main

import (
	"fmt"
	dgf "github.com/snowel/dgf"
	)

var Starpoints19 = []int{61, 175, 289, 295, 181, 67, 73, 187, 301}
// TODO Board size starpoints
// TODO coordinates

func DrawBoard(board []byte, bstone string, wstone string, intersection string, starpoint string, bcursor string, wcursor string, cursorPosition int, leftMargin string) string {
	s := leftMargin
  	var starpoints []int
		size := (int)(board[3])
		switch size {
		case 19:
			starpoints = append(starpoints, Starpoints19...)
		}


    for index, v := range board {
				// Skip the Board offset bytes
				if index <= dgf.BOff {continue}
				// Get index i equal to intersections numbers
				i := index - dgf.BOff
				// Draw cursor when applicable
				if i == cursorPosition && dgf.IsLib(board, i) {
					  switch dgf.GetPlayer(board){
						case 1:
						 	s += fmt.Sprintf("%s ", bcursor)
						case 2:
							s += fmt.Sprintf("%s ", wcursor)}
					  if i % size == 0 {
						 s += fmt.Sprintf("\n%s", leftMargin)// This will result in a keft amrgin character at the start of the line under the baord... TODO
					  }
					  continue
				}	
				// Draw empy intersection when applicable
				if dgf.IsLib(board, i) {
					if dgf.Contains(i, starpoints){
					  s += fmt.Sprintf("%s ", starpoint)
					} else {
					  s += fmt.Sprintf("%s ", intersection)
					}
				}

				// Draw stones when applicable
				if v == 1 {
					  s += fmt.Sprintf("%s ", bstone)
				}
				if v == 2 {
					  s += fmt.Sprintf("%s ", wstone)
				}

				// Newline when appicable
				if i % size == 0 {
					s += fmt.Sprintf("\n%s", leftMargin) //Smae issue
				 }
			 }

			 return s
}

// Print out the captures and current turn of a board
// Sep is a string between the fields, may be a new line or other graphic
func DrawBoardExtras(board []byte, sep, bplayer, bstone, wplayer, wstone string) string{
  player := dgf.GetPlayer(board)
  var pstone string
 	switch player {
	case 1:
	  pstone = bplayer
	case 2:
	  pstone = wplayer
	}
	bcap, wcap := dgf.GetCaps(board)
	return fmt.Sprintf("Current Move: %s %s %s %s: %f %s %s %s: %f",pstone, sep, bstone, bplayer, bcap, sep, wstone, wplayer, wcap)
}

// Places a stone on a board without validation 
func PlaceStone(board []byte, inter int) {
 	dgf.ApplyMove(board, inter) // This doesn't validate moves
 }

// Plays a move, if valid both the record and board are updated
func PlayMove(board []byte, move int) bool {
	valid, newBoard := dgf.ValidateMoveWko(board, move)
	if valid {
		dgf.BoardCopy(board, newBoard)
		return true
	}

	return false
}
/*

func PlayMove(rec []int, board []byte, move int) bool {
	valid, newBoard := dgf.RecValidateMove(rec, move)
	if valid {
		dgf.BoardCopy(board, newBoard)
		rec = append(rec, move)
		return true
	}
	return false
}

*/
