package diff

const (
	// removeCost gives the cost to remove A at the given index.
	removeCost = -1

	// addCost gives the cost to add B at the given index.
	addCost = -1

	/// substitionCost gives the substition cost for replacing A with B at the given indices.
	substitionCost = -2

	/// equalCost gives the cost for A and B being equal.
	equalCost = 0
)

// container is a container for the comparable used to determine subset and
// revered reading of the data in the comparable.
type container struct {
	comp    Comparable
	aOffset int
	aLength int
	bOffset int
	bLength int
	reverse bool
}

// newContainer creates a new comparable container with the given subset and reverse settings.
func newContainer(comp Comparable, aOffset, aLength, bOffset, bLength int, reverse bool) *container {
	return &container{
		comp:    comp,
		aOffset: aOffset,
		aLength: aLength,
		bOffset: bOffset,
		bLength: bLength,
		reverse: reverse,
	}
}

// fullContainer creates a new comparable for a full container.
func fullContainer(comp Comparable) *container {
	return newContainer(comp, 0, comp.ALength(), 0, comp.BLength(), false)
}

// aAdjust gets the A index adjusted by the container's condition.
func (cont *container) aAdjust(aIndex int) int {
	if cont.reverse {
		return cont.aLength - 1 - aIndex + cont.aOffset
	}
	return aIndex + cont.aOffset
}

// bAdjust gets the B index adjusted by the container's condition.
func (cont *container) bAdjust(bIndex int) int {
	if cont.reverse {
		return cont.bLength - 1 - bIndex + cont.bOffset
	}
	return bIndex + cont.bOffset
}

// Sub creates a new comparable container for a subset and reverse relative to this container's settings.
func (cont *container) Sub(aLow, aHigh, bLow, bHigh int, reverse bool) *container {
	if cont.reverse {
		return newContainer(cont.comp, cont.aAdjust(aHigh), aHigh-aLow, cont.bAdjust(bHigh), bHigh-bLow, !reverse)
	}
	return newContainer(cont.comp, cont.aAdjust(aLow), aHigh-aLow, cont.bAdjust(bLow), bHigh-bLow, reverse)
}

// ALength is the length of the first list being compared.
func (cont *container) ALength() int {
	return cont.aLength
}

// BLength is the length of the second list being compared.
func (cont *container) BLength() int {
	return cont.bLength
}

// Equals determines if the entries in the two given indices are equal.
func (cont *container) Equals(aIndex, bIndex int) bool {
	return cont.comp.Equals(cont.aAdjust(aIndex), cont.bAdjust(bIndex))
}

// SubstitionCost determines the substition cost for the given indices.
func (cont *container) SubstitionCost(i, j int) int {
	if cont.Equals(i, j) {
		return equalCost
	}
	return substitionCost
}
