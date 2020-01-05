package diff

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

// StringSlicePath gets the difference path for the two given string slices.
func StringSlicePath(a, b []string) []StepGroup {
	return Path(&stringSliceComparable{a: a, b: b})
}

// PlusMinus gets the labelled difference between the two slices.
// It formats the results by prepending a "+" to new strings in [b],
// a "-" for any to removed strings from [a], and " " if the strings are the same.
func PlusMinus(a, b []string) []string {
	result := []string{}
	aIndex, bIndex := 0, 0
	path := StringSlicePath(a, b)
	for _, step := range path {
		switch step.Step {
		case EqualStep:
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, " "+a[aIndex])
				aIndex++
				bIndex++
			}
		case AddedStep:
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, "+"+b[bIndex])
				bIndex++
			}
		case RemovedStep:
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, "-"+a[aIndex])
				aIndex++
			}
		}
	}
	return result
}

// Merge gets the labelled difference between the two slices
// using a similar output to the git merge differences output.
func Merge(a, b []string) []string {
	result := []string{}
	aIndex, bIndex := 0, 0
	path := StringSlicePath(a, b)

	const startChange = "<<<<<<<<"
	const middleChange = "========"
	const endChange = ">>>>>>>>"

	prevState := EqualStep
	for _, step := range path {
		switch step.Step {
		case EqualStep:
			switch prevState {
			case AddedStep:
				result = append(result, endChange)
			case RemovedStep:
				result = append(result, middleChange)
				result = append(result, endChange)
			}
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, a[aIndex])
				aIndex++
				bIndex++
			}

		case AddedStep:
			switch prevState {
			case EqualStep:
				result = append(result, startChange)
				result = append(result, middleChange)
			case RemovedStep:
				result = append(result, middleChange)
			}
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, b[bIndex])
				bIndex++
			}

		case RemovedStep:
			switch prevState {
			case EqualStep:
				result = append(result, startChange)
			case AddedStep:
				result = append(result, middleChange)
			}
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, a[aIndex])
				aIndex++
			}
		}
		prevState = step.Step
	}

	switch prevState {
	case AddedStep:
		result = append(result, endChange)
	case RemovedStep:
		result = append(result, middleChange)
		result = append(result, endChange)
	}
	return result
}
