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

// StepGroup is a continuous group of step types.
type StepGroup struct {

	// step is the type for this group.
	Step StepType

	// count is the number of the given type in the group.
	Count int
}

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
// See https://en.wikipedia.org/wiki/Levenshtein_distance
func Path(comp Comparable) []StepGroup {
	result := []StepGroup{}
	dd := newDiffData(comp)
	dd.setMovement(comp.ALength(), comp.BLength())

	addRun := 0
	insertAdd := func() {
		if addRun > 0 {
			result = append(result, StepGroup{Step: AddedStep, Count: addRun})
			addRun = 0
		}
	}

	removeRun := 0
	insertRemove := func() {
		if removeRun > 0 {
			result = append(result, StepGroup{Step: RemovedStep, Count: removeRun})
			removeRun = 0
		}
	}

	equalRun := 0
	insertEqual := func() {
		if equalRun > 0 {
			result = append(result, StepGroup{Step: EqualStep, Count: equalRun})
			equalRun = 0
		}
	}

	dd.traverseBackwards(func(step StepType, count int) {
		switch step {
		case EqualStep:
			insertAdd()
			insertRemove()
			equalRun += count
			break
		case AddedStep:
			insertEqual()
			addRun += count
			break
		case RemovedStep:
			insertEqual()
			removeRun += count
			break
		}
	})
	insertAdd()
	insertRemove()
	insertEqual()

	for i, j := len(result)-1, 0; i > j; i, j = i-1, j+1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// minimum gets the minimum value of the given values
func minimum(first int, rest ...int) int {
	min := first
	for _, value := range rest {
		if value < min {
			min = value
		}
	}
	return min
}
