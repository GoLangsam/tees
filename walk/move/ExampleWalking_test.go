// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package move_test
// package main

import (
	"fmt"

	"github.com/GoLangsam/tees/list"
	"github.com/GoLangsam/tees/list/math"
	"github.com/GoLangsam/tees/list/walk/move"
)

func ExampleGotoNext() {
//func main() {
	// Create a new list
	var chess = list.NewList( "Chess" )

	// Create new lists
	var cols = chess.AddBeam( "Spalten", "A", "B", "C", "D", "E" )
	var rows = chess.AddBeam( "Zeilen", 1, 2, 3, 4, 5 )
	fmt.Println("Starting")

	var board = math.Xross( cols, rows )
	board.PrintAways( "Schach" )
	board.Root().Away().List().PrintAways( "Brett" )

	if false {
	board.PrintAtomValues( "board \t")
	board.Root().Away().List().PrintAtomValues( "board.Root().Away().List().Root().Away() \t")
	board.Front().PrintAtomValues("board.Front() \t")
	board.Front().Away().PrintAtomValues("board.Front().Away() \t")
	board.Front().Away().List().PrintAtomValues("board.Front().Away().List() \t")
	board.Front().Away().Next().PrintAtomValues("board.Front().Away().Next() \t")
	board.Front().Away().List().Front().PrintAtomValues("board.Front().Away().List().Front() \t")
	board.Front().Away().List().Front().Away().PrintAtomValues("board.Front().Away().List().Front().Away() \t")
	}

	var s = board.Front().Away().List().Front()		// A|1
	var m = s.Next().Next().Away().Next().Next().Away()	// C|3
	var e = board.Back().Away().List().Back()		// E|5
	if s.IsRoot() || m.IsRoot() || e.IsRoot() {}

	fmt.Println( "\nJumpNext" )
//	move.GotoNext.PrintFullWalk( s )
	move.GotoNext.PrintFullWalk( m )
//	move.GotoNext.PrintFullWalk( e )

	fmt.Println( "\nJumpDown" )
//	move.MoveDown.PrintFullWalk( s )
	move.MoveDown.PrintFullWalk( m )
//	move.MoveDown.PrintFullWalk( e )

	fmt.Println( "\nJumpRightDown" )
//	move.DiagRightDown.PrintFullWalk( s )
	move.DiagRightDown.PrintFullWalk( m )
//	move.DiagRightDown.PrintFullWalk( e )

	fmt.Println( "\nChessRook" )
//	move.ChessRook.PrintFullWalk( s )
	move.ChessRook.PrintFullWalk( m )
//	move.ChessRook.PrintFullWalk( e )

	fmt.Println( "\nChessBishop" )
//	move.ChessBishop.PrintFullWalk( s )
	move.ChessBishop.PrintFullWalk( m )
//	move.ChessBishop.PrintFullWalk( e )

	fmt.Println( "\nChessKnight" )
//	move.ChessKnight.PrintFullWalk( s )
	move.ChessBishop.PrintFullWalk( m )
//	move.ChessKnight.PrintFullWalk( e )

	fmt.Println( "\nChessKing" )
//	move.ChessKing.PrintFullWalk( s )
	move.ChessKing.PrintFullWalk( m )
//	move.ChessKing.PrintFullWalk( e )

	fmt.Println( "\nChessQueen" )
//	move.ChessQueen.PrintFullWalk( s )
	move.ChessQueen.PrintFullWalk( m )
//	move.ChessQueen.PrintFullWalk( e )

	// Output:
	// Starting
	// Schach: List=Spalten | Total=5
	// => : List=A	A|1 | A|2 | A|3 | A|4 | A|5 | Total=5
	// => : List=B	B|1 | B|2 | B|3 | B|4 | B|5 | Total=5
	// => : List=C	C|1 | C|2 | C|3 | C|4 | C|5 | Total=5
	// => : List=D	D|1 | D|2 | D|3 | D|4 | D|5 | Total=5
	// => : List=E	E|1 | E|2 | E|3 | E|4 | E|5 | Total=5
	// Brett: List=Zeilen | Total=5
	// => : List=1	1|A | 1|B | 1|C | 1|D | 1|E | Total=5
	// => : List=2	2|A | 2|B | 2|C | 2|D | 2|E | Total=5
	// => : List=3	3|A | 3|B | 3|C | 3|D | 3|E | Total=5
	// => : List=4	4|A | 4|B | 4|C | 4|D | 4|E | Total=5
	// => : List=5	5|A | 5|B | 5|C | 5|D | 5|E | Total=5
	//
	// JumpNext
	// From: C|3> C|4 -> End <	Distance: 1
	// Grab: C|3> C|4 -> End <	Distance: 1
	// Haul: C|3> C|4 -> C|5 -> End <	Distance: 2
	//
	// JumpDown
	// From: C|3> D|3 -> End <	Distance: 1
	// Grab: C|3> 3|C -> 3|D -> D|3 -> End <	Distance: 1
	// Haul: C|3> D|3 -> E|3 -> End <	Distance: 2
	//
	// JumpRightDown
	// From: C|3> D|4 -> End <	Distance: 2
	// Grab: C|3> C|4 -> 4|C -> 4|D -> D|4 -> End <	Distance: 2
	// Haul: C|3> D|4 -> E|5 -> End <	Distance: 4
	//
	// ChessRook
	// From: C|3> C|2 -> C|4 -> B|3 -> D|3 -> End <	Distance: 4
	// Grab: C|3> C|2 -> C|4 -> 3|C -> 3|B -> B|3 -> 3|C -> 3|D -> D|3 -> End <	Distance: 4
	// Haul: C|3> C|2 -> C|1 -> C|4 -> C|5 -> B|3 -> A|3 -> D|3 -> E|3 -> End <	Distance: 8
	//
	// ChessBishop
	// From: C|3> B|4 -> D|4 -> B|2 -> D|2 -> End <	Distance: 8
	// Grab: C|3> C|4 -> 4|C -> 4|B -> B|4 -> C|4 -> 4|C -> 4|D -> D|4 -> C|2 -> 2|C -> 2|B -> B|2 -> C|2 -> 2|C -> 2|D -> D|2 -> End <	Distance: 8
	// Haul: C|3> B|4 -> A|5 -> D|4 -> E|5 -> B|2 -> A|1 -> D|2 -> E|1 -> End <	Distance: 16
	//
	// ChessKnight
	// From: C|3> B|4 -> D|4 -> B|2 -> D|2 -> End <	Distance: 8
	// Grab: C|3> C|4 -> 4|C -> 4|B -> B|4 -> C|4 -> 4|C -> 4|D -> D|4 -> C|2 -> 2|C -> 2|B -> B|2 -> C|2 -> 2|C -> 2|D -> D|2 -> End <	Distance: 8
	// Haul: C|3> B|4 -> A|5 -> D|4 -> E|5 -> B|2 -> A|1 -> D|2 -> E|1 -> End <	Distance: 16
	//
	// ChessKing
	// From: C|3> C|2 -> C|4 -> B|3 -> D|3 -> B|4 -> D|4 -> B|2 -> D|2 -> End <	Distance: 12
	// Grab: C|3> C|2 -> C|4 -> 3|C -> 3|B -> B|3 -> 3|C -> 3|D -> D|3 -> C|4 -> 4|C -> 4|B -> B|4 -> C|4 -> 4|C -> 4|D -> D|4 -> C|2 -> 2|C -> 2|B -> B|2 -> C|2 -> 2|C -> 2|D -> D|2 -> End <	Distance: 12
	// Haul: C|3> C|2 -> C|1 -> C|4 -> C|5 -> B|3 -> A|3 -> D|3 -> E|3 -> B|4 -> A|5 -> D|4 -> E|5 -> B|2 -> A|1 -> D|2 -> E|1 -> End <	Distance: 24
	//
	// ChessQueen
	// From: C|3> C|2 -> C|4 -> B|3 -> D|3 -> B|4 -> D|4 -> B|2 -> D|2 -> End <	Distance: 12
	// Grab: C|3> C|2 -> C|4 -> 3|C -> 3|B -> B|3 -> 3|C -> 3|D -> D|3 -> C|4 -> 4|C -> 4|B -> B|4 -> C|4 -> 4|C -> 4|D -> D|4 -> C|2 -> 2|C -> 2|B -> B|2 -> C|2 -> 2|C -> 2|D -> D|2 -> End <	Distance: 12
	// Haul: C|3> C|2 -> C|1 -> C|4 -> C|5 -> B|3 -> A|3 -> D|3 -> E|3 -> B|4 -> A|5 -> D|4 -> E|5 -> B|2 -> A|1 -> D|2 -> E|1 -> End <	Distance: 24
}
