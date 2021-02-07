package diff

type (
	// StepType is the steps of the levenshtein path.
	StepType int

	// Step is a continuous group of step types.
	Step struct {

		// step is the type for this group.
		Type StepType

		// count is the number of the given type in the group.
		Count int
	}

	// PathCallback is the function signature for calling back steps in the path.
	PathCallback func(step Step)
)

const (
	// Equal indicates A and B entries are equal.
	Equal StepType = iota

	// Added indicates A was added.
	Added

	// Removed indicates A was removed.
	Removed
)

// WalkPath calls back the difference path for the given comparable.
func WalkPath(comp Comparable, hndl PathCallback) {
	addRun := 0
	insertAdd := func() {
		if addRun > 0 {
			hndl(Step{Type: Added, Count: addRun})
			addRun = 0
		}
	}

	removeRun := 0
	insertRemove := func() {
		if removeRun > 0 {
			hndl(Step{Type: Removed, Count: removeRun})
			removeRun = 0
		}
	}

	equalRun := 0
	insertEqual := func() {
		if equalRun > 0 {
			hndl(Step{Type: Equal, Count: equalRun})
			equalRun = 0
		}
	}

	b := newPathBuilder(comp, func(step Step) {
		switch step.Type {
		case Added:
			insertEqual()
			addRun += step.Count
			break

		case Removed:
			insertEqual()
			removeRun += step.Count
			break

		case Equal:
			insertRemove()
			insertAdd()
			equalRun += step.Count
			break
		}
	})
	b.Build()

	insertEqual()
	insertRemove()
	insertAdd()
}

// pathBuilder is a Levenshtein/Hirschberg path builder used for diffing two comparable sources.
type pathBuilder struct {

	// baseCont is the source comparable to create the path for.
	baseCont *container

	// hndl is the callback to return path steps.
	hndl PathCallback

	// scoreFront is the score vector at the front of the score calculation.
	scoreFront []int

	// scoreBack is the score vector at the back of the score calculation.
	scoreBack []int

	// scoreOther is the score vector to store off a result vector to.
	scoreOther []int
}

/// newPathBuilder creates a new path builder.
func newPathBuilder(comp Comparable, hndl PathCallback) *pathBuilder {
	len := comp.BLength() + 1
	return &pathBuilder{
		baseCont:   fullContainer(comp),
		hndl:       hndl,
		scoreFront: make([]int, len),
		scoreBack:  make([]int, len),
		scoreOther: make([]int, len),
	}
}

// swapScores swaps the front and back score vectors.
func (b *pathBuilder) swapScores() {
	b.scoreBack, b.scoreFront = b.scoreFront, b.scoreBack
}

// storeScore swaps the back and other score vectors.
func (b *pathBuilder) storeScore() {
	b.scoreBack, b.scoreOther = b.scoreOther, b.scoreBack
}

// max gets the maximum value of the three given values.
func max(a, b, c int) int {
	result := a
	if result < b {
		result = b
	}
	if result < c {
		result = c
	}
	return result
}

// calculateScore calculate the Needleman-Wunsch score.
// At the end of this calculation the score is in the back vector.
func (b *pathBuilder) calculateScore(cont *container) {
	aLen := cont.ALength()
	bLen := cont.BLength()

	b.scoreBack[0] = 0
	for j := 1; j <= bLen; j++ {
		b.scoreBack[j] = b.scoreBack[j-1] + addCost
	}

	for i := 1; i <= aLen; i++ {
		b.scoreFront[0] = b.scoreBack[0] + removeCost
		for j := 1; j <= bLen; j++ {
			b.scoreFront[j] = max(
				b.scoreBack[j-1]+cont.SubstitionCost(i-1, j-1),
				b.scoreBack[j]+removeCost,
				b.scoreFront[j-1]+addCost)
		}

		b.swapScores()
	}
}

// findPivot finds the pivot between the other score and the reverse of the back score.
// The pivot is the index of the maximum sum of each element in the two scores.
func (b *pathBuilder) findPivot(bLength int) int {
	index := 0
	max := b.scoreOther[0] + b.scoreBack[bLength]
	for j := 1; j <= bLength; j++ {
		value := b.scoreOther[j] + b.scoreBack[bLength-j]
		if value > max {
			max = value
			index = j
		}
	}
	return index
}

// aEdge handles when at the edge of the A source subset in the given container.
func (b *pathBuilder) aEdge(cont *container) {
	aLen := cont.ALength()
	bLen := cont.BLength()

	if aLen <= 0 {
		if bLen > 0 {
			b.hndl(Step{Type: Added, Count: bLen})
		}
		return
	}

	split := -1
	for j := 0; j < bLen; j++ {
		if cont.Equals(0, j) {
			split = j
			break
		}
	}

	if split < 0 {
		b.hndl(Step{Type: Removed, Count: 1})
		b.hndl(Step{Type: Added, Count: bLen})
	} else {
		if split > 0 {
			b.hndl(Step{Type: Added, Count: split})
		}
		b.hndl(Step{Type: Equal, Count: 1})
		if split < bLen-1 {
			b.hndl(Step{Type: Added, Count: bLen - split - 1})
		}
	}
}

// bEdge Handles when at the edge of the B source subset in the given container.
func (b *pathBuilder) bEdge(cont *container) {
	aLen := cont.ALength()
	bLen := cont.BLength()

	if bLen <= 0 {
		if aLen > 0 {
			b.hndl(Step{Type: Removed, Count: aLen})
		}
		return
	}

	split := -1
	for i := 0; i < aLen; i++ {
		if cont.Equals(i, 0) {
			split = i
			break
		}
	}

	if split < 0 {
		b.hndl(Step{Type: Removed, Count: aLen})
		b.hndl(Step{Type: Added, Count: 1})
	} else {
		if split > 0 {
			b.hndl(Step{Type: Removed, Count: split})
		}
		b.hndl(Step{Type: Equal, Count: 1})
		if split < aLen-1 {
			b.hndl(Step{Type: Removed, Count: aLen - split - 1})
		}
	}
}

// breakupPath performs the Hirschberg divide and conquer and returns the path.
func (b *pathBuilder) breakupPath(cont *container) {
	aLen := cont.ALength()
	bLen := cont.BLength()

	if aLen <= 1 {
		b.aEdge(cont)
		return
	}

	if bLen <= 1 {
		b.bEdge(cont)
		return
	}

	aMid := aLen / 2
	b.calculateScore(cont.Sub(0, aMid, 0, bLen, false))
	b.storeScore()
	b.calculateScore(cont.Sub(aMid, aLen, 0, bLen, true))
	bMid := b.findPivot(bLen)

	b.breakupPath(cont.Sub(0, aMid, 0, bMid, false))
	b.breakupPath(cont.Sub(aMid, aLen, bMid, bLen, false))
}

// Build builds the diff path for the base content.
func (b *pathBuilder) Build() {
	b.breakupPath(b.baseCont)
}
