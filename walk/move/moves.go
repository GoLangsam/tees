// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package move

import (
	"github.com/GoLangsam/tees/list/walk"
)

// For Your convenience: Some predefined Moves:

// for the basic GoTo's
	var GotoPrev	walk.Kata = []walk.GoTo{ walk.Prev }
	var GotoNext	walk.Kata = []walk.GoTo{ walk.Next }
	var GotoAway	walk.Kata = []walk.GoTo{ walk.Away }
	var GotoRoot	walk.Kata = []walk.GoTo{ walk.Root }
	var GotoFront	walk.Kata = []walk.GoTo{ walk.Front }
	var GotoBack	walk.Kata = []walk.GoTo{ walk.Back }
	var GotoHome	walk.Kata = []walk.GoTo{ walk.Home }

// for composed movement
// Orthogonal
// Hint: Think of a right-hand matrix (or chessboard), look at a horizontal element
// Note: A Spreadsheet or Matrix is usually left(!)-handed, Rows go down!!!
	var MoveAwayPrev	walk.Kata = []walk.GoTo{ walk.Away, walk.Prev, walk.Away }		//	= Up
	var MoveAwayNext	walk.Kata = []walk.GoTo{ walk.Away, walk.Next, walk.Away }		//	= Down
	var MoveUp		walk.Kata = MoveAwayPrev
	var MoveDown		walk.Kata = MoveAwayNext

// next/prev orthogonal fiber
	var TrixNextFront	walk.Kata = []walk.GoTo{ walk.Root, walk.Away, walk.Next, walk.Away, walk.Front }
	var TrixNextBack	walk.Kata = []walk.GoTo{ walk.Root, walk.Away, walk.Next, walk.Away, walk.Back }
	var TrixPrevFront	walk.Kata = []walk.GoTo{ walk.Root, walk.Away, walk.Prev, walk.Away, walk.Front}
	var TrixPrevBack	walk.Kata = []walk.GoTo{ walk.Root, walk.Away, walk.Prev, walk.Away, walk.Back }

// Diagonal
	var JumpDiagonalNU	walk.Kata = []walk.GoTo{ walk.Next, walk.Away, walk.Prev, walk.Away }		//	= Next, AwayPrev
	var JumpDiagonalND	walk.Kata = []walk.GoTo{ walk.Next, walk.Away, walk.Next, walk.Away }		//	= Next, AwayNext
	var JumpDiagonalPU	walk.Kata = []walk.GoTo{ walk.Prev, walk.Away, walk.Prev, walk.Away }		//	= Prev, AwayPrev
	var JumpDiagonalPD	walk.Kata = []walk.GoTo{ walk.Prev, walk.Away, walk.Next, walk.Away }		//	= Prev, AwayPrev

// Diagonal - chess: King:Jump, Bishop:Tour, Queen:Tour
	var DiagRightUp		walk.Kata = JumpDiagonalNU
	var DiagRightDown	walk.Kata = JumpDiagonalND
	var DiagLeftUp		walk.Kata = JumpDiagonalPU
	var DiagLeftDown	walk.Kata = JumpDiagonalPD

// Knight's
	var JumpKnight_NNU	walk.Kata = []walk.GoTo{ walk.Next, walk.Next, walk.Away, walk.Prev, walk.Away }	//	= Next, RightUp
	var JumpKnight_NND	walk.Kata = []walk.GoTo{ walk.Next, walk.Next, walk.Away, walk.Next, walk.Away }	//	= Next, RightDown
	var JumpKnight_PPU	walk.Kata = []walk.GoTo{ walk.Prev, walk.Prev, walk.Away, walk.Prev, walk.Away }	//	= Prev, RightUp
	var JumpKnight_PPD	walk.Kata = []walk.GoTo{ walk.Prev, walk.Prev, walk.Away, walk.Next, walk.Away }	//	= Prev, RightDown
	var JumpKnight_UUP	walk.Kata = []walk.GoTo{ walk.Away, walk.Prev, walk.Prev, walk.Away, walk.Prev }	//
	var JumpKnight_UUN	walk.Kata = []walk.GoTo{ walk.Away, walk.Prev, walk.Prev, walk.Away, walk.Next }	//
	var JumpKnight_DDP	walk.Kata = []walk.GoTo{ walk.Away, walk.Next, walk.Next, walk.Away, walk.Prev }	//
	var JumpKnight_DDN	walk.Kata = []walk.GoTo{ walk.Away, walk.Next, walk.Next, walk.Away, walk.Next }	//

// Chess
	var ChessKing		walk.Akas = []walk.Kata{ GotoPrev, GotoNext, MoveUp, MoveDown, DiagRightUp, DiagRightDown, DiagLeftUp, DiagLeftDown }
// Note: King Grab's only, for others use Haul or Walker!
	var ChessQueen		walk.Akas = ChessKing
	var ChessRook		walk.Akas = []walk.Kata{ GotoPrev, GotoNext, MoveUp, MoveDown }
	var ChessBishop		walk.Akas = []walk.Kata{                                       DiagRightUp, DiagRightDown, DiagLeftUp, DiagLeftDown }

	var ChessKnight		walk.Akas = []walk.Kata{
		JumpKnight_NNU,
		JumpKnight_NND,
		JumpKnight_PPU,
		JumpKnight_PPD,
		JumpKnight_UUP,
		JumpKnight_UUN,
		JumpKnight_DDP,
		JumpKnight_DDN }