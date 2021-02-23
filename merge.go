package diff

import "../../comparable"

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

	prevState := Equal
	path.Read(func(step StepType, count int) {
		switch step {
		case Equal:
			switch prevState {
			case Added:
				result = append(result, endChange)
			case Removed:
				result = append(result, middleChange)
				result = append(result, endChange)
			}
			for i := count - 1; i >= 0; i-- {
				result = append(result, a[aIndex])
				aIndex++
				bIndex++
			}

		case Added:
			switch prevState {
			case Equal:
				result = append(result, startChange)
				result = append(result, middleChange)
			case Removed:
				result = append(result, middleChange)
			}
			for i := count - 1; i >= 0; i-- {
				result = append(result, b[bIndex])
				bIndex++
			}

		case Removed:
			switch prevState {
			case Equal:
				result = append(result, startChange)
			case Added:
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
	case Added:
		result = append(result, endChange)
	case Removed:
		result = append(result, middleChange)
		result = append(result, endChange)
	}
	return result
}
