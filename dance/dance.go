// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dance

import (
	"github.com/GoLangsam/tees"
	"github.com/GoLangsam/tees/dance/dancers/chooser"
	"github.com/GoLangsam/tees/dance/dancers/drummer"
	"github.com/GoLangsam/tees/dance/dancers/spinner" // to be supplied by client; carries problem transparently. See also DefaultDance below
	"github.com/GoLangsam/tees/dance/dancing"
	stacker "github.com/GoLangsam/tees/dance/dancing/deeh" // TODO(AP): find a better place
)

// Dance is what dancer loves to do.
type Dance struct {
	Dancer *dancing.Dancing

	Stacker *stacker.Stack
	Drummer *drummer.Drums

	VerboseGoals bool
}

// ========================================================

// InitialDepth of space allocation for stack and drum
var InitialDepth = 100

// NewDance returns a fresh Dance
//
// Verbosity is set accordingly
func NewDance(v, vg, rh, vd, vb, vc bool) *Dance {
	var d = new(Dance)
	Verbose = VerboseType(v)
	d.VerboseGoals = vg
	chooser.Verbose = chooser.VerboseType(vc)

	d.Dancer = dancing.New(vb, rh, false) // TODO: ??? arg for turn.Verbose

	d.Stacker = stacker.New(InitialDepth)
	d.Drummer = drummer.NewDrums(InitialDepth, vd)

	d = d.setBeating()
	d = d.setDancing()
	d = d.setTurning()

	return d
}

// ========================================================

func (d *Dance) setBeating() *Dance {
	d.Dancer.Beating.Dance = func(l *tees.Tees) { d.Drummer.Goal.Beat(d.Dancer.Level) }
	d.Dancer.Beating.OnGoal = func() { d.Drummer.Call.Beat(d.Dancer.Level) }
	d.Dancer.Beating.OnFail = func() { d.Drummer.Fail.Beat(d.Dancer.Level) }
	d.Dancer.Beating.OnLeaf = func() { d.Drummer.Leaf.Beat(d.Dancer.Level) }
	return d
}

func (d *Dance) setDancing() *Dance {
	d.Dancer.Dancing.Dance = func() { return } // Client overrides: func(){ spinner.Dance(Dance.Dancing, cols) }
	d.Dancer.Dancing.OnGoal = d.Stacker.Push
	d.Dancer.Dancing.OnFail = d.Stacker.Pop
	d.Dancer.Dancing.OnLeaf = d.Dancer.Beating.OnLeaf // Note: defined in SetBeating above
	return d
}

func (d *Dance) setTurning() *Dance {
	d.Dancer.Turning.Dance = func(l *tees.Tees) { l.Dance(d.Dancer.Dancing) }
	d.Dancer.Turning.OnGoal = d.onGoal            // YES We have a solution
	d.Dancer.Turning.OnFail = chooser.ChooseShort // YES We have to go on dancing & goaling
	d.Dancer.Turning.OnLeaf = d.onLeaf            // YES We have to abort
	return d
}

// ========================================================

func (d *Dance) onLeaf(l *tees.Tees) bool { // Do we have to abort?
	if Verbose {
		d.Stacker.Top().PrintValue("Stack-Top")
		l.PrintValue()
		l.PrintAways()
	}
	return false
}

func (d *Dance) onGoal(l *tees.Tees) bool { // Do we have a solution
	if l.IsEmpty() { // l.Len() == 0 // l.Size() == 0	// YES We have a solution
		if d.VerboseGoals {
			d.PrintGoal()
		} // ... we may show it
		return true
	}

	return false
}

// ========================================================

// DefaultDance returns the default dancing function
func (d *Dance) DefaultDance(dancing *dancing.Dancing, list *tees.Tees) func() {
	return func() { spinner.Dance(dancing, list) }
}
