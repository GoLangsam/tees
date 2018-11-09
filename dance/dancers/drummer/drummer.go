// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package drummer supplies drums (counters)
package drummer

import (
	"github.com/GoLangsam/tees/dance/dancers/drummer/drum" // == "github.com/GoLangsam/container/oneway/drum"
)

// Drums consolidates all needed drums
type Drums struct {
	Goal *drum.Drum // Niveaus counts dances per Level
	Fail *drum.Drum // Deadend counts fail ends per Level
	Call *drum.Drum // Grooves counts solutions per length
	Leaf *drum.Drum // UpDates counts unLink per Level
	// Note: No drum for Push & Pop, as these occur in sync with Call
}

// NewDrums provides a fresh ensemble of drums
func NewDrums(cap int, verbose bool) *Drums {
	var d = new(Drums)
	d = d.init(cap, verbose)
	return d
}

func (d *Drums) init(cap int, verbose bool) *Drums {
	d.Goal = drum.NewDrum("Grooves", cap)
	d.Fail = drum.NewDrum("Deadend", cap)
	d.Call = drum.NewDrum("Niveaus", cap)
	d.Leaf = drum.NewDrum("UpDates", cap)

	if verbose {
		drum.Verbose = true
	} else {
		drum.Verbose = false
	}

	return d
}

// Print has all drums print iff Verbose
func (d *Drums) Print() {
	if Verbose {
		d.Goal.Print()
		d.Fail.Print()
		d.Call.Print()
		d.Leaf.Print()
	}
}
