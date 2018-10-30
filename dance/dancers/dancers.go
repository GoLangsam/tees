// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package dancers provides factories of
// exemplary and default implementations
// of the constituents common to dancers
// (read: what it takes to be a dancer)
// in it's subdirectories:
//  spinner
//  - a dancer and some lower-level Turner
//  chooser
//  - given Tees: tells where (and if at all) to continue;
//  - knows `tees.Tees`, and calls them `among`
//  drummer
//  - brings his own drum
//  - is useful for statistical counting
package dancers

// ========================================================
