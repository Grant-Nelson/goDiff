package container

import (
	"github.com/Grant-Nelson/goDiff/comparable"
	"github.com/Grant-Nelson/goDiff/internal/collector"
)

const (
	// RemoveCost gives the cost to remove A at the given index.
	RemoveCost = 1

	// AddCost gives the cost to add B at the given index.
	AddCost = 1

	// SubstitionCost gives the substition cost for replacing A with B at the given indices.
	SubstitionCost = 2

	// EqualCost gives the cost for A and B being equal.
	EqualCost = 0
)

type (
	// Diff is the interface for a diff algorithm.
	Diff interface {

		// CanDiff determines if the diff algorithm can handle a container with
		// the amount of data inside of the given container. If this returns false a
		// larger matrix, cache, vector, or whatever would be needed to perform the diff.
		CanDiff(cont *Container) bool

		// Diff performs the algorithm on the given container
		// and writes the results to the collector.
		Diff(cont *Container, col *collector.Collector)
	}

	// Container is a container for the comparable used to determine subset and
	// revered reading of the data in the comparable.
	Container struct {
		comp    comparable.Comparable
		aOffset int
		aLength int
		bOffset int
		bLength int
		reverse bool
	}
)

// min gets the minimum value from the two given values.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// check that the container also implements the comparable.
var _ comparable.Comparable = (*Container)(nil)

// newSub creates a new comparable for the given range.
func newSub(comp comparable.Comparable, aOffset, aLength, bOffset, bLength int, reverse bool) *Container {
	return &Container{
		comp:    comp,
		aOffset: aOffset,
		aLength: aLength,
		bOffset: bOffset,
		bLength: bLength,
		reverse: reverse,
	}
}

// New creates a new comparable for a full container.
func New(comp comparable.Comparable) *Container {
	return newSub(comp,
		0, comp.ALength(),
		0, comp.BLength(),
		false)
}

// aAdjust gets the A index adjusted by the container's condition.
func (cont *Container) aAdjust(aIndex int) int {
	if cont.reverse {
		return cont.aLength - 1 - aIndex + cont.aOffset
	}
	return aIndex + cont.aOffset
}

// bAdjust gets the B index adjusted by the container's condition.
func (cont *Container) bAdjust(bIndex int) int {
	if cont.reverse {
		return cont.bLength - 1 - bIndex + cont.bOffset
	}
	return bIndex + cont.bOffset
}

// ALength is the length of the first list being compared.
func (cont *Container) ALength() int {
	return cont.aLength
}

// BLength is the length of the second list being compared.
func (cont *Container) BLength() int {
	return cont.bLength
}

// Equals determines if the entries in the two given indices are equal.
func (cont *Container) Equals(aIndex, bIndex int) bool {
	return cont.comp.Equals(cont.aAdjust(aIndex), cont.bAdjust(bIndex))
}

// SubstitionCost determines the substition cost for the given indices.
func (cont *Container) SubstitionCost(i, j int) int {
	if cont.Equals(i, j) {
		return EqualCost
	}
	return SubstitionCost
}

// Sub creates a new comparable container for a subset and reverse relative to this container's settings.
func (cont *Container) Sub(aLow, aHigh, bLow, bHigh int, reverse bool) *Container {
	if cont.reverse {
		return newSub(cont.comp,
			cont.aAdjust(aHigh), aHigh-aLow,
			cont.bAdjust(bHigh), bHigh-bLow,
			!reverse)
	}
	return newSub(cont.comp,
		cont.aAdjust(aLow), aHigh-aLow,
		cont.bAdjust(bLow), bHigh-bLow,
		reverse)
}

// Reduce determines how much of the edges of this container are equal.
// The amount before and after which are equal are returned and
// the reduced subcontainer is returned.
func (cont *Container) Reduce() (sub *Container, before, after int) {
	width := min(cont.aLength, cont.bLength)
	for before = 0; before < width; before++ {
		if !cont.Equals(before, before) {
			break
		}
	}

	width = width - before
	for after = 0; after < width; after++ {
		if !cont.Equals(cont.aLength-1-after, cont.bLength-1-after) {
			break
		}
	}

	return cont.Sub(
		before, cont.aLength-after,
		before, cont.bLength-after,
		cont.reverse), before, after
}
