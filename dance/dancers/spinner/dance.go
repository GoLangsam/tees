// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinner

import (
	"github.com/GoLangsam/tees"
	"github.com/GoLangsam/tees/dance/dancing"
)

// ========================================================

// Dance keeps a dancer turning to dance dancing or not
func Dance(d *dancing.Dancing, l *tees.Tees) {

	if d.Turning.OnLeaf(l){						// YES We have to abort
		return							// ... we quietly return
	}

	if d.Turning.OnGoal(l){						// YES We have a solution
		if d.Verbose {d.Beating.Dance(l)}			// ... we count our happyness
		return
	}
	next, ok := d.Turning.OnFail(l); if !ok {			// YES We have a failure
		if d.Verbose {d.Beating.OnFail()}			// ... we count our suffering
		return

	} else {							// YES We have to go on goaling
		if d.Verbose {d.Beating.OnGoal()}			// ... we count our effort
		if next == nil { panic("Confusing reply from Turning.OnFail: Cannot look into nil!")}
	}
	d.Level++
	d.Turning.Dance(next)
	d.Level--
}
