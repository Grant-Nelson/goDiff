package diff

import "strings"

type stepType int

const (
	equalStep stepType = iota
	diffStep
	addedStep
	removedStep
)

// PartDiff gets the labelled differences between the values.
// The given seperator will split the values and join the result.
func PartDiff(a, b, sep string) string {
	aParts := strings.Split(a, sep)
	bParts := strings.Split(b, sep)
	result := Diff(aParts, bParts)
	return strings.Join(result, sep)
}

// Diff gets the labelled difference between the two slices.
func Diff(a, b []string) []string {
	result := []string{}
	aRun, bRun := []string{}, []string{}
	aIndex, bIndex := 0, 0
	_, path := levenshteinDistance(a, b, len(a), len(b))
	for _, step := range path {
		switch step {
		case equalStep:
			result = append(result, aRun...)
			result = append(result, bRun...)
			aRun, bRun = []string{}, []string{}
			result = append(result, " "+a[aIndex])
			aIndex++
			bIndex++
		case diffStep:
			aRun = append(aRun, "-"+a[aIndex])
			bRun = append(bRun, "+"+b[bIndex])
			aIndex++
			bIndex++
		case addedStep:
			bRun = append(bRun, "+"+b[bIndex])
			bIndex++
		case removedStep:
			aRun = append(aRun, "-"+a[aIndex])
			aIndex++
		}
	}
	result = append(result, aRun...)
	result = append(result, bRun...)
	return result
}

// levenshteinDistance gets the levenshtein distance and calculates the path.
// The path is the simplest difference between the two sets.
// See https://en.wikipedia.org/wiki/Levenshtein_distance
func levenshteinDistance(a, b []string, aIndex, bIndex int) (int, []stepType) {
	// base case, empty sets
	if aIndex <= 0 {
		minPath := []stepType{}
		for i := bIndex; i > 0; i-- {
			minPath = append(minPath, addedStep)
		}
		return bIndex, minPath
	}
	if bIndex <= 0 {
		minPath := []stepType{}
		for i := aIndex; i > 0; i-- {
			minPath = append(minPath, removedStep)
		}
		return aIndex, minPath
	}

	// get the minimum of entry skip entry from a, skip entry from b, and skip entry from both
	skipA, aPath := levenshteinDistance(a, b, aIndex-1, bIndex)
	skipB, bPath := levenshteinDistance(a, b, aIndex, bIndex-1)
	skipC, cPath := levenshteinDistance(a, b, aIndex-1, bIndex-1)

	// calculate the minimum path and add costs
	skipA++
	skipB++
	cStep := equalStep
	if a[aIndex-1] != b[bIndex-1] {
		skipC++
		cStep = diffStep
	}
	skipMin, minPath, minStep := skipA, aPath, removedStep
	if skipB < skipMin {
		skipMin, minPath, minStep = skipB, bPath, addedStep
	}
	if skipC < skipMin {
		skipMin, minPath, minStep = skipC, cPath, cStep
	}
	return skipMin, append(minPath, minStep)
}
