// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinner

import (
	"github.com/GoLangsam/tees"
	"github.com/GoLangsam/tees/dance/dancing/turn"
)

// ========================================================

// TurnFast keeps a dancer turning to dance dancing or not - but without beating
func TurnFast(d *turn.Dancing, l *tees.Tees) {

	if d.OnLeaf(l){				return	}	// YES We have to abort
	if d.OnGoal(l){				return	}	// YES We have a solution
	next, ok := d.OnFail(l); if !ok {	return	}	// YES We have a failure

//	d.Level++
	d.Dance(next)
//	d.Level--
}

// TurnEver keeps a dancer turning to dance dancing or not - but with less turning, and without beating
//  Note: this may only be useful for benchmarks / OnLeaf-counting
func TurnEver(d *turn.Dancing, l *tees.Tees) {

//	if d.OnLeaf(l){				return	}	// YES We have to abort
//	if d.OnGoal(l){				return	}	// YES We have a solution
	if l.IsEmpty(){				return	}	// YES We have a solution - but we don't tell anyone
	next, ok := d.OnFail(l); if !ok {	return	}	// YES We have a failure

//	d.Level++
	d.Dance(next)
//	d.Level--
}
