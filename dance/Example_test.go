// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dance_test

//package main

import (
	"testing"

	"github.com/GoLangsam/tees/dance"
	"github.com/GoLangsam/tees/list/test"
)

// ========================================================

func NQueensR(anz int) {

	//	fmt.Println()
	//	fmt.Println( "Queens on a", anz, "Board" )

	var Dance = dance.NewDance(false, false, false, false, false, false)

	var cols = test.NQueensR(anz)
	Dance.Dancer.Dancing.Dance = Dance.DefaultDance(Dance.Dancer, cols)
	Dance.Dancer.Dancing.Dance()
	Dance.Print()
}

// ========================================================

func NQueensRTest(beg, end int) {
	for i := beg; i <= end; i++ {
		NQueensR(i)
	}
}

// ========================================================

func ExampleNQueensR() {
	// func main() {

	beg := 1
	end := 13

	NQueensRTest(beg, end)
}

// ========================================================

func BenchmarkNQueensR(b *testing.B) {
	b.ReportAllocs()
	for i := 1; i < b.N; i++ {
		NQueensR(12)
	}
}

// ========================================================
