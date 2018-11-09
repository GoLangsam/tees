// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//package dance_test
package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/GoLangsam/tees"
	"github.com/GoLangsam/tees/dance"
	"github.com/GoLangsam/tees/dance/dancers/chooser"
	"github.com/GoLangsam/tees/list/test"
)

// ========================================================

var (
	beg = flag.Int("Beg", 1, "Begin with Boardsize (min=1, max=End)")
	end = flag.Int("End", 12, "End with Boardsize (min=1)")

	fastF = flag.Bool("Fast", false, "Use DanceFast, or")
	listF = flag.Bool("List", false, "Use list/Danc, or")
	slowF = flag.Bool("Slow", false, "Use DanceSlow")

	sortF = flag.Bool("Sort", false, "Use Sort-by-size")

	beats = flag.Bool("b", false, "Beats : Print counters per Level (sets -d)")
	drums = flag.Bool("d", false, "Drums : Print counters (sets -r)")
	rhyth = flag.Bool("r", false, "Rhythm: Collect counters")
	goals = flag.Bool("s", false, "Verbose Scoring: Print solutions")
	times = flag.Bool("t", true,  "Verbose Timings: Print times")
	choos = flag.Bool("u", false, "Verbose Chooser: Print all choices made")
	verbs = flag.Bool("v", false, "Verbose Dancing: Prints a lot of details")
	tests = flag.Bool("x", false, "Execute Chooser-test")
	dancs = flag.Bool("y", false, "Execute Dancing-test (-Fast -norm -Slow -List)")

//	s = flag.String("s", " ", "separator")
)

func flagParse() {
	flag.Parse()
	if *beats {	*drums = true	}
	if *drums {	*rhyth = true	}

	if *beg < 1 {	*beg = 1	}
	if *end < 1 {	*end = 1	}
	if *beg > *end {*beg = *end	}

	if *slowF && *fastF {	panic("Either or: -Slow, or -Fast")	}
	if *slowF && *listF {	panic("Either or: -Slow, or -List")	}
	if *listF && *fastF {	panic("Either or: -List, or -Fast")	}
}

// ========================================================

// Chooser - for chooser tests
var Chooser chooser.Chooser = chooser.ChooseShort // default

func nQueensR(anz int) {

	runtime.GC()

	var t time.Time
	if *times {	t = time.Now()		}

	var Dance *dance.Dance = dance.NewDance(*verbs, *goals, *rhyth, *drums, *beats, *choos)

	var		dance = func(l *tees.Tees) 	{	l.Dance(		Dance.Dancer.Dancing)	}
	if *slowF {	dance = func(l *tees.Tees) 	{	l.DanceSlow(		Dance.Dancer.Dancing)	}	}
	if *listF {	dance = func(l *tees.Tees) 	{	tees.Dance( l.Root(),	Dance.Dancer.Dancing)	}	}
	if *fastF {	dance = func(l *tees.Tees) 	{	l.DanceFast(		Dance.Dancer.Dancing)	}	}
	Dance.Dancer.Turning.Dance = dance
	Dance.Dancer.Turning.OnFail = Chooser // only needed for Choose-tests!

	var dancing = test.NQueensR(anz)
	if *sortF {	dancing = tees.Arrange(dancing, func(e *tees.This) int {return e.Size()})	}
	Dance.Dancer.Dancing.Dance = Dance.DefaultDance(Dance.Dancer, dancing)
	Dance.Dancer.Dancing.Dance()
	Dance.Print()

	if *times {	fmt.Println("Time =", time.Since(t))		}
}

// ========================================================

func nQueensRDancingTest(beg, end int) {
	for i := beg; i <= end; i++ {
		fmt.Println()
		fmt.Println("Queens on a", i, "*", i, "Board")
		dancingTest(i)
	}
}

func dancingTest(anz int) {
	*fastF, *listF, *slowF = true,  false, false
	fmt.Println("Dacing Fast")
	nQueensR(anz)

//	time.Sleep(1 * time.Second)
	*fastF, *listF, *slowF = false, false, false
	fmt.Println("Dacing norm")
	nQueensR(anz)

//	time.Sleep(1 * time.Second)
	*fastF, *listF, *slowF = false, false, true
	fmt.Println("Dacing Slow")
	nQueensR(anz)

//	time.Sleep(1 * time.Second)
	*fastF, *listF, *slowF = false, true,  false
	fmt.Println("Dacing List")
	nQueensR(anz)
}

// ========================================================

func nQueensRChooserTest(beg, end int) {
	for i := beg; i <= end; i++ {
		fmt.Println()
		fmt.Println("Queens on a", i, "*", i, "Board")
		chooserTest(i)
	}
}
func chooserTest(anz int) {
	for i, c := range chooser.GetChoosers() {
		Chooser = c
		fmt.Println("Chooser No", i+1, "=", chooser.ChooserName(i))
		nQueensR(anz)
	}
}

func chooserNQueensRTest(beg, end int) {
	for i, c := range chooser.GetChoosers() {
		Chooser = c
		fmt.Println()
		fmt.Println("Chooser No", i+1, "=", chooser.ChooserName(i))
		nQueensRTest(beg, end)
	}
}
func nQueensRTest(beg, end int) {
	for i := beg; i <= end; i++ {
		fmt.Println("Queens on a", i, "*", i, "Board")
		nQueensR(i)
	}
}

// ========================================================

// func ExampleTestChoosers(t *testing.T) {
func main() {
	flagParse()

	if *tests {
		if *beg == *end {
			nQueensRChooserTest(*beg, *end)
		} else {
			chooserNQueensRTest(*beg, *end)
		}
	} else if *dancs {
		if *beg == *end {
			nQueensRDancingTest(*beg, *end)
		} else {
			nQueensRDancingTest(*beg, *end)
		}
	} else {
		if *beg == *end {
			fmt.Println()
			fmt.Println("Queens on a", *beg, "*", *beg, "Board")
			nQueensR(*beg)
		} else {
			fmt.Println()
			nQueensRTest(*beg, *end)
		}
	}
}

// ========================================================
