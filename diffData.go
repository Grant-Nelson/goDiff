package diff

import (
	"fmt"
)

type (
	// movementType is the type of movement
	movementType int

	// diffData is a container for the levenshtein costs and store subpath movements.
	diffData struct {
		comp  Comparable
		costs [][]int
		moves [][]movementType
	}
)

const (
	moveNotSet movementType = iota
	moveUp
	moveLeft
	moveUpLeft
	moveEqual
)

// newDiffData creates a new cost matrix and movement matrix.
func newDiffData(comp Comparable) *diffData {
	aLen, bLen := comp.ALength(), comp.BLength()
	costs := make([][]int, aLen)
	moves := make([][]movementType, aLen)
	for aIndex := 0; aIndex < aLen; aIndex++ {
		costs[aIndex] = make([]int, bLen)
		moves[aIndex] = make([]movementType, bLen)
		for bIndex := 0; bIndex < bLen; bIndex++ {
			costs[aIndex][bIndex] = -1
			moves[aIndex][bIndex] = moveNotSet
		}
	}
	return &diffData{
		comp:  comp,
		costs: costs,
		moves: moves,
	}
}

/// isEqual checks if the comparer is equal.
func (dd *diffData) isEqual(aIndex, bIndex int) bool {
	return dd.comp.Equals(aIndex-1, bIndex-1)
}

// setCost sets the path point.
func (dd *diffData) setCost(aIndex, bIndex int) int {
	// get the minimum of entry skip entry from a, skip entry from b, and skip entry from both
	costA := dd.getCost(aIndex-1, bIndex)
	costB := dd.getCost(aIndex, bIndex-1)
	minCost := minimum(costA, costB)

	costC := dd.getCost(aIndex-1, bIndex-1)
	if costC <= minCost {
		// skips any cost for equal values in the inputs
		skipCost := 0
		if dd.isEqual(aIndex, bIndex) {
			skipCost = -1
		}
		minCost = costC + skipCost
	}

	// calculate the minimum path cost and set cost
	minCost++
	dd.costs[aIndex-1][bIndex-1] = minCost
	return minCost
}

// getCost gets the path cost.
func (dd *diffData) getCost(aIndex, bIndex int) int {
	if aIndex <= 0 {
		return bIndex
	}
	if bIndex <= 0 {
		return aIndex
	}
	cost := dd.costs[aIndex-1][bIndex-1]
	if cost < 0 {
		cost = dd.setCost(aIndex, bIndex)
	}
	return cost
}

// setMovement determines the minimum path starting from this point
// and then sets the movement from this point towards the shorter path.
func (dd *diffData) setMovement(aIndex, bIndex int) int {
	// base case when one of the inputs are empty
	if aIndex <= 0 {
		return bIndex
	}
	if bIndex <= 0 {
		return aIndex
	}

	// Check if this subpath has already been solved.
	if dd.moves[aIndex-1][bIndex-1] != moveNotSet {
		return dd.costs[aIndex-1][bIndex-1]
	}

	// get the minimum of entry skip entry from a, skip entry from b, and skip entry from both
	costA := dd.getCost(aIndex-1, bIndex)
	costB := dd.getCost(aIndex, bIndex-1)
	costC := dd.getCost(aIndex-1, bIndex-1)
	minCost := minimum(costA, costB, costC)

	// calculate the minimum path cost and set movements
	minPathCost := minCost + 2
	minMove := moveNotSet

	if costA <= minCost {
		// costA is minimum
		cost := dd.setMovement(aIndex-1, bIndex) + 1
		if cost < minPathCost {
			minPathCost = cost
			minMove = moveLeft
		}
	}

	if costB <= minCost {
		// costB is minimum
		cost := dd.setMovement(aIndex, bIndex-1) + 1
		if cost < minPathCost {
			minPathCost = cost
			minMove = moveUp
		}
	}

	if costC <= minCost {
		cost := dd.setMovement(aIndex-1, bIndex-1)
		if dd.isEqual(aIndex, bIndex) {
			if cost < minPathCost {
				minPathCost = cost
				minMove = moveEqual
			}
		} else {
			cost++
			if cost < minPathCost {
				minPathCost = cost
				minMove = moveUpLeft
			}
		}
	}

	dd.moves[aIndex-1][bIndex-1] = minMove
	return minPathCost
}

// traverseBackwards handles traversing the diff path using the defined movements,
// however it traverses backwards.
func (dd *diffData) traverseBackwards(hndl func(StepType, int)) {
	aIndex := dd.comp.ALength()
	bIndex := dd.comp.BLength()
	for {
		if aIndex <= 0 {
			hndl(AddedStep, bIndex)
			return
		}

		if bIndex <= 0 {
			hndl(RemovedStep, aIndex)
			return
		}

		switch dd.moves[aIndex-1][bIndex-1] {
		case moveLeft:
			aIndex--
			hndl(RemovedStep, 1)
			break
		case moveUp:
			bIndex--
			hndl(AddedStep, 1)
			break
		case moveEqual:
			aIndex--
			bIndex--
			hndl(EqualStep, 1)
			break
		case moveUpLeft:
			aIndex--
			bIndex--
			hndl(AddedStep, 1)
			hndl(RemovedStep, 1)
			break
		default:
			panic(fmt.Errorf("hit not set at (%d, %d)", aIndex, bIndex))
		}
	}
}
