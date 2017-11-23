package diff

import (
	"strings"
)

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
