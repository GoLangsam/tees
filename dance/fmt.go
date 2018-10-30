// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dance

import (
	"fmt"
)

// ========================================================

func (d *Dance) Print() {
	d.Drummer.Print()
}

// ========================================================

// Print prints the current stack (= the solution)
func (d *Dance) PrintGoal() {
	fmt.Print("Solution: ")
	fmt.Println(d.Stacker.Len())
	for _, l := range d.Stacker.Get() {
		l.List().PrintValue()
		fmt.Print(": ")
		for e := l.AwayList().Front(); e != nil; e = e.Next() {
			e.AwayList().PrintValue()
			fmt.Print(" ")
		}
		fmt.Println(".")
	}
}
