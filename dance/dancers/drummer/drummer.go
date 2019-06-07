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
	Goal drum.Drum // Niveaus counts dances per Level
	Fail drum.Drum // Deadend counts fail ends per Level
	Call drum.Drum // Grooves counts solutions per length
	Leaf drum.Drum // UpDates counts unLink per Level
	// Note: No drum for Push & Pop, as these occur in sync with Call
	Verbose bool
}

// NewDrums provides a fresh ensemble of drums
func NewDrums(cap int, verbose bool) *Drums {
	var d = new(Drums)
	return d.init(cap, verbose)
}

func (d *Drums) init(cap int, verbose bool) *Drums {
	d.Verbose = verbose

	d.Goal = drum.NewDrum("Grooves", cap)
	d.Fail = drum.NewDrum("Deadend", cap)
	d.Call = drum.NewDrum("Niveaus", cap)
	d.Leaf = drum.NewDrum("UpDates", cap)

	d.Goal.Verbose = d.Verbose
	d.Fail.Verbose = d.Verbose
	d.Call.Verbose = d.Verbose
	d.Leaf.Verbose = d.Verbose

	return d
}

// Print has all drums print iff Verbose
func (d *Drums) Print() {
	if d.Verbose {
		d.Goal.Print()
		d.Fail.Print()
		d.Call.Print()
		d.Leaf.Print()
	}
}
