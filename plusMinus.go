package diff

import (
	"./comparable"
	"./step"
)

// PlusMinus gets the labelled difference between the two slices.
// It formats the results by prepending a "+" to new strings in [b],
// a "-" for any to removed strings from [a], and " " if the strings are the same.
func PlusMinus(a, b []string) []string {
	path := Diff(comparable.NewString(a, b))

	result := make([]string, 0, path.Total())
	aIndex, bIndex := 0, 0
	path.Read(func(step step.Type, count int) {
		switch step {
		case step.Equal:
			for i := count - 1; i >= 0; i-- {
				result = append(result, " "+a[aIndex])
				aIndex++
				bIndex++
			}
		case step.Added:
			for i := count - 1; i >= 0; i-- {
				result = append(result, "+"+b[bIndex])
				bIndex++
			}
		case step.Removed:
			for i := count - 1; i >= 0; i-- {
				result = append(result, "-"+a[aIndex])
				aIndex++
			}
		}
	})
	return result
}
