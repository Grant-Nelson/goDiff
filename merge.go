package godiff

import (
	"github.com/Grant-Nelson/goDiff/comparable"
	"github.com/Grant-Nelson/goDiff/step"
)

// Merge gets the labelled difference between the two slices
// using a similar output to the git merge differences output.
// This will use the default diff configuration to perform the diff.
func Merge(a, b []string) []string {
	return MergeCustom(nil, a, b)
}

// MergeCustom gets the labelled difference between the two slices
// using a similar output to the git merge differences output.
// This was can use any given diff algorithm.
func MergeCustom(diff Algorithm, a, b []string) []string {
	if diff == nil {
		diff = DefaultDiff()
	}
	path := diff(comparable.NewString(a, b))

	const (
		startChange  = "<<<<<<<<"
		middleChange = "========"
		endChange    = ">>>>>>>>"
	)

	result := make([]string, 0, path.Total()+path.Count()*2+1)
	aIndex, bIndex := 0, 0

	prevState := step.Equal
	path.Read(func(stepType step.Type, count int) {
		switch stepType {
		case step.Equal:
			switch prevState {
			case step.Added:
				result = append(result, endChange)
			case step.Removed:
				result = append(result, middleChange)
				result = append(result, endChange)
			}
			for i := count - 1; i >= 0; i-- {
				result = append(result, a[aIndex])
				aIndex++
				bIndex++
			}

		case step.Added:
			switch prevState {
			case step.Equal:
				result = append(result, startChange)
				result = append(result, middleChange)
			case step.Removed:
				result = append(result, middleChange)
			}
			for i := count - 1; i >= 0; i-- {
				result = append(result, b[bIndex])
				bIndex++
			}

		case step.Removed:
			switch prevState {
			case step.Equal:
				result = append(result, startChange)
			case step.Added:
				result = append(result, endChange)
				result = append(result, startChange)
			}
			for i := count - 1; i >= 0; i-- {
				result = append(result, a[aIndex])
				aIndex++
			}
		}
		prevState = stepType
	})

	switch prevState {
	case step.Added:
		result = append(result, endChange)
	case step.Removed:
		result = append(result, middleChange)
		result = append(result, endChange)
	}
	return result
}
