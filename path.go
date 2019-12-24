package diff

// StepType is the steps of the levenshtein path.
type StepType int

const (
	// EqualStep indicates A and B entries are equal.
	EqualStep StepType = iota

	// AddedStep indicates A was added.
	AddedStep

	// RemovedStep indicates A was removed.
	RemovedStep
)

// Comparable is a simple interface for generic difference determination.
type Comparable interface {

	// ALength is the length of the first list being compared.
	ALength() int

	// BLength is the length of the secton list being compared.
	BLength() int

	// Equals determines if the entries in the two given indices are equal.
	Equals(aIndex, bIndex int) bool
}

// Path gets the difference path for the two given items.
func Path(comp Comparable) []StepType {
	result := []StepType{}
	aLen, bLen := comp.ALength(), comp.BLength()
	costMatrix := levenshteinDistance(comp, aLen, bLen)
	_, path := traverseLevenshteinDistance(comp, aLen, bLen, costMatrix)

	// Group additions and removals together.
	aRun, bRun := []StepType{}, []StepType{}
	for _, step := range path {
		switch step {
		case EqualStep:
			result = append(result, aRun...)
			result = append(result, bRun...)
			aRun, bRun = []StepType{}, []StepType{}
			result = append(result, EqualStep)
		case AddedStep:
			bRun = append(bRun, AddedStep)
		case RemovedStep:
			aRun = append(aRun, RemovedStep)
		}
	}
	result = append(result, aRun...)
	result = append(result, bRun...)
	return result
}

// costMatrix is a container for the levenshtein costs.
type costMatrix struct {
	costs [][]int
}

// newCostMatrix creates a new cost matrix.
func newCostMatrix(aLen, bLen int) *costMatrix {
	costs := make([][]int, aLen)
	for aIndex := 0; aIndex < aLen; aIndex++ {
		costs[aIndex] = make([]int, bLen)
	}
	return &costMatrix{
		costs: costs,
	}
}

// setCost sets the path cost.
func (cm *costMatrix) setCost(aIndex, bIndex, cost int) {
	cm.costs[aIndex-1][bIndex-1] = cost
}

// getCost gets the path cost.
func (cm *costMatrix) getCost(aIndex, bIndex int) int {
	if aIndex <= 0 {
		return bIndex
	}
	if bIndex <= 0 {
		return aIndex
	}
	return cm.costs[aIndex-1][bIndex-1]
}

// minimum gets the minimum value of the given three values
func minimum(a, b, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}

// levenshteinDistance gets cost matrix of levenshtein distances.
// This filles out the cost matrix.
// See https://en.wikipedia.org/wiki/Levenshtein_distance
func levenshteinDistance(comp Comparable, aLen, bLen int) *costMatrix {
	costs := newCostMatrix(aLen, bLen)
	for aIndex := 1; aIndex <= aLen; aIndex++ {
		for bIndex := 1; bIndex <= bLen; bIndex++ {

			// skips any cost for equal values in the inputs.
			skipCost := 1
			if comp.Equals(aIndex-1, bIndex-1) {
				skipCost = 0
			}

			// get the minimum of entry skip entry from a, skip entry from b, and skip entry from both
			costA := costs.getCost(aIndex-1, bIndex) + 1
			costB := costs.getCost(aIndex, bIndex-1) + 1
			costC := costs.getCost(aIndex-1, bIndex-1) + skipCost

			// calculate the minimum path cost and set cost
			costs.setCost(aIndex, bIndex, minimum(costA, costB, costC))
		}
	}
	return costs
}

// fillSteps will fill a path of the given length with the given step.
func fillSteps(count int, step StepType) []StepType {
	path := make([]StepType, count)
	for i := count - 1; i >= 0; i-- {
		path[i] = step
	}
	return path
}

// traverseLevenshteinDistance gets the path with the lowest cost.
func traverseLevenshteinDistance(comp Comparable, aIndex, bIndex int, costs *costMatrix) (int, []StepType) {
	// base case when one of the inputs are empty
	if aIndex <= 0 {
		return bIndex, fillSteps(bIndex, AddedStep)
	}
	if bIndex <= 0 {
		return aIndex, fillSteps(aIndex, RemovedStep)
	}

	// get the minimum of entry skip entry from a, skip entry from b, and skip entry from both
	costA := costs.getCost(aIndex-1, bIndex)
	costB := costs.getCost(aIndex, bIndex-1)
	costC := costs.getCost(aIndex-1, bIndex-1)
	minCost := minimum(costA, costB, costC)

	// calculate the minimum path cost and set cost
	minPathCost := minCost + 2
	minPath := []StepType{}
	if costA <= minCost {
		// costA is minimum
		cost, path := traverseLevenshteinDistance(comp, aIndex-1, bIndex, costs)
		cost++
		if cost < minPathCost {
			minPathCost = cost
			minPath = append(path, RemovedStep)
		}
	}

	if costB <= minCost {
		// costB is minimum
		cost, path := traverseLevenshteinDistance(comp, aIndex, bIndex-1, costs)
		cost++
		if cost < minPathCost {
			minPathCost = cost
			minPath = append(path, AddedStep)
		}
	}

	if costC <= minCost {
		if comp.Equals(aIndex-1, bIndex-1) {
			// costC is minimum and entries equal
			cost, path := traverseLevenshteinDistance(comp, aIndex-1, bIndex-1, costs)
			// Do not add to cost since the values are equal
			if cost < minPathCost {
				minPathCost = cost
				minPath = append(path, EqualStep)
			}
		} else {
			// costC is minimum and entries different
			cost, path := traverseLevenshteinDistance(comp, aIndex-1, bIndex-1, costs)
			cost++
			if cost < minPathCost {
				minPathCost = cost
				minPath = append(path, RemovedStep, AddedStep)
			}
		}
	}

	return minPathCost, minPath
}
