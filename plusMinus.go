package godiff

import (
	"github.com/Grant-Nelson/goDiff/comparable"
	"github.com/Grant-Nelson/goDiff/step"
)

// PlusMinus gets the labelled difference between the two slices.
// It formats the results by prepending a "+" to new strings in [b],
// a "-" for any to removed strings from [a], and " " if the strings are the same.
// This will use the default diff configuration to perform the diff.
func PlusMinus(a, b []string) []string {
	return PlusMinusCustom(nil, a, b)
}

// PlusMinusCustom gets the labelled difference between the two slices.
// It formats the results by prepending a "+" to new strings in [b],
// a "-" for any to removed strings from [a], and " " if the strings are the same.
// This was can use any given diff algorithm.
func PlusMinusCustom(diff Algorithm, a, b []string) []string {
	if diff == nil {
		diff = DefaultDiff()
	}
	path := diff(comparable.NewString(a, b))

	result := make([]string, 0, path.Total())
	aIndex, bIndex := 0, 0
	path.Read(func(stepType step.Type, count int) {
		switch stepType {
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
