package diff

import (
	"./comparable"
	"./step"
)

// Merge gets the labelled difference between the two slices
// using a similar output to the git merge differences output.
func Merge(a, b []string) []string {
	const (
		startChange  = "<<<<<<<<"
		middleChange = "========"
		endChange    = ">>>>>>>>"
	)

	path := Diff(comparable.NewString(a, b))

	result := make([]string, 0, path.Total()+path.Count()*2)
	aIndex, bIndex := 0, 0

	prevState := step.Equal
	path.Read(func(step step.Type, count int) {
		switch step {
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
				result = append(result, middleChange)
			}
			for i := count - 1; i >= 0; i-- {
				result = append(result, a[aIndex])
				aIndex++
			}
		}
		prevState = step
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
