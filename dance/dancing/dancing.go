// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dancing

import (
	"github.com/GoLangsam/tees/dance/dancing/beat" // Note: No drum for Push & Pop, as these occur in sync with Call
	"github.com/GoLangsam/tees/dance/dancing/deeh"
	"github.com/GoLangsam/tees/dance/dancing/turn"
)

// Dancing consolidates beet, deeh and turn Dancings
// and allows to keep track of the Level
type Dancing struct {
	Beating *beat.Dancing
	Dancing *deeh.Dancing
	Turning *turn.Dancing

	Level int
}

// New returns a fresh Dancing
// with applicable verbosities set
func New(vd, vb, vt bool) *Dancing {
	var d = new(Dancing)

	d.Beating = new(beat.Dancing)
	d.Dancing = new(deeh.Dancing)
	d.Turning = new(turn.Dancing)
	d.Level = 0

	Verbose = VerboseType(vd)
	beat.Verbose = beat.VerboseType(vb)
//	deeh.Verbose = deeh.VerboseType(vl)	// intentionally does not Verbose
	turn.Verbose = turn.VerboseType(vt)

	return d
}
