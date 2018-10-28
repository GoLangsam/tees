// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import (
	"fmt"
)

// KataAkas common interface
type KataAkas interface {
	//	From(Here) has to have differing returns!
	Grab(Here) (Trail, Distance)
	Haul(Here) (Trail, Distance)
	PrintFullWalk(Here)
	PrintWalker(string, Here)
	Walker(Here) Walk
}

func (steps Kata) PrintFullWalk(e Here) {
	fmt.Print("From: ")
	from, d := steps.From(e)
	e.PrintValue()
	fmt.Print("> ")
	from.PrintValue()
	fmt.Print(" -> ")
	fmt.Print("End <")
	fmt.Printf("\tDistance: %v\n", d)

	fmt.Print("Grab: ")
	grab, d := steps.Grab(e)
	e.PrintValue()
	fmt.Print("> ")
	for _, e := range grab {
		if e != nil {
			e.PrintValue()
			fmt.Print(" -> ")
		}
	}
	fmt.Print("End <")
	fmt.Printf("\tDistance: %v\n", d)

	fmt.Print("Haul: ")
	haul, d := steps.Haul(e)
	e.PrintValue()
	fmt.Print("> ")
	for _, e := range haul {
		if e != nil {
			e.PrintValue()
			fmt.Print(" -> ")
		}
	}
	fmt.Print("End <")
	fmt.Printf("\tDistance: %v\n", d)
}

func (jumps Akas) PrintFullWalk(e Here) {
	fmt.Print("From: ")
	from, d := jumps.From(e)
	e.PrintValue()
	fmt.Print("> ")
	for _, e := range from {
		if e != nil {
			e.PrintValue()
			fmt.Print(" -> ")
		}
	}
	fmt.Print("End <")
	fmt.Printf("\tDistance: %v\n", d)

	fmt.Print("Grab: ")
	grab, d := jumps.Grab(e)
	e.PrintValue()
	fmt.Print("> ")
	for _, e := range grab {
		if e != nil {
			e.PrintValue()
			fmt.Print(" -> ")
		}
	}
	fmt.Print("End <")
	fmt.Printf("\tDistance: %v\n", d)

	fmt.Print("Haul: ")
	haul, d := jumps.Haul(e)
	e.PrintValue()
	fmt.Print("> ")
	for _, e := range haul {
		if e != nil {
			e.PrintValue()
			fmt.Print(" -> ")
		}
	}
	fmt.Print("End <")
	fmt.Printf("\tDistance: %v\n", d)
}

func (steps Kata) PrintWalker(name string, from Here) {
	steps.Walker(from).Print(name, from)
}

func (jumps Akas) PrintWalker(name string, from Here) {
	jumps.Walker(from).Print(name, from)
}

func (next Walk) Print(name string, from Here) {
	from.PrintValue(name + "\t=>")
	for x := next(); x != nil; x = next() {
		x.PrintValue(" ->")
	}
	fmt.Println()
}
