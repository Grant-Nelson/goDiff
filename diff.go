package diff

import "strings"

// StepType is the steps of the levenshtein path.
type StepType int

const (
	EqualStep   StepType = iota // EqualStep indicates A and B entries are equal
	AddedStep                   // AddStep indicates B was added
	RemovedStep                 // RemoveStep indicates A was removed
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

// stringSliceComparable is a comparable for two string slices.
type stringSliceComparable struct {
	a []string
	b []string
}

func (comp *stringSliceComparable) ALength() int { return len(comp.a) }
func (comp *stringSliceComparable) BLength() int { return len(comp.b) }
func (comp *stringSliceComparable) Equals(aIndex, bIndex int) bool {
	return comp.a[aIndex] == comp.b[bIndex]
}

// Strings gets the labelled differences between the values.
// The given seperator will split the values and join the result.
func Strings(a, b, sep string) string {
	aParts := strings.Split(a, sep)
	bParts := strings.Split(b, sep)
	result := Slices(aParts, bParts)
	return strings.Join(result, sep)
}

// Slices gets the labelled difference between the two slices.
func Slices(a, b []string) []string {
	result := []string{}
	aIndex, bIndex := 0, 0
	path := Path(&stringSliceComparable{a: a, b: b})
	for _, step := range path {
		switch step {
		case EqualStep:
			result = append(result, " "+a[aIndex])
			aIndex++
			bIndex++
		case AddedStep:
			result = append(result, "+"+b[bIndex])
			bIndex++
		case RemovedStep:
			result = append(result, "-"+a[aIndex])
			aIndex++
		}
	}
	return result
}

// Path gets the difference path for the two given items.
func Path(comp Comparable) []StepType {
	result := []StepType{}
	_, path := levenshteinDistance(comp, comp.ALength(), comp.BLength())

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

// levenshteinDistance gets the levenshtein distance and calculates the path.
// The path is the simplest difference between the two sets.
// See https://en.wikipedia.org/wiki/Levenshtein_distance
func levenshteinDistance(comp Comparable, aIndex, bIndex int) (int, []StepType) {
	// base case, empty sets
	if aIndex <= 0 {
		minPath := []StepType{}
		for i := bIndex; i > 0; i-- {
			minPath = append(minPath, AddedStep)
		}
		return bIndex, minPath
	}
	if bIndex <= 0 {
		minPath := []StepType{}
		for i := aIndex; i > 0; i-- {
			minPath = append(minPath, RemovedStep)
		}
		return aIndex, minPath
	}

	// get the minimum of entry skip entry from a, skip entry from b, and skip entry from both
	skipA, aPath := levenshteinDistance(comp, aIndex-1, bIndex)
	skipB, bPath := levenshteinDistance(comp, aIndex, bIndex-1)
	skipC, cPath := levenshteinDistance(comp, aIndex-1, bIndex-1)

	// calculate the minimum path and add costs
	skipA++
	skipB++
	cSteps := []StepType{EqualStep}
	if !comp.Equals(aIndex-1, bIndex-1) {
		skipC++
		cSteps = []StepType{RemovedStep, AddedStep}
	}
	skipMin, minPath, minSteps := skipA, aPath, []StepType{RemovedStep}
	if skipB < skipMin {
		skipMin, minPath, minSteps = skipB, bPath, []StepType{AddedStep}
	}
	if skipC < skipMin {
		skipMin, minPath, minSteps = skipC, cPath, cSteps
	}
	return skipMin, append(minPath, minSteps...)
}
