package diff

// StringSlicePath gets the difference path for the two given string slices.
func StringSlicePath(a, b []string) []Step {
	return Path(newStrSliceComp(a, b))
}

// RuneSlicesPath gets the difference path for the two given runes slices.
func RuneSlicesPath(a, b [][]rune) []Step {
	return Path(newRuneSliceComp(a, b))
}

// Path gets the difference path for the given comparable.
func Path(comp Comparable) []Step {
	path := []Step{}
	WalkPath(comp, func(step Step) {
		path = append(path, step)
	})
	return path
}

// PlusMinus gets the labelled difference between the two slices.
// It formats the results by prepending a "+" to new strings in [b],
// a "-" for any to removed strings from [a], and " " if the strings are the same.
func PlusMinus(a, b []string) []string {
	result := []string{}
	aIndex, bIndex := 0, 0
	path := StringSlicePath(a, b)
	for _, step := range path {
		switch step.Type {
		case Equal:
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, " "+a[aIndex])
				aIndex++
				bIndex++
			}
		case Added:
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, "+"+b[bIndex])
				bIndex++
			}
		case Removed:
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

	const (
		startChange  = "<<<<<<<<"
		middleChange = "========"
		endChange    = ">>>>>>>>"
	)

	prevState := Equal
	for _, step := range path {
		switch step.Type {
		case Equal:
			switch prevState {
			case Added:
				result = append(result, endChange)
			case Removed:
				result = append(result, middleChange)
				result = append(result, endChange)
			}
			for i := step.Count - 1; i >= 0; i-- {
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
			for i := step.Count - 1; i >= 0; i-- {
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
			for i := step.Count - 1; i >= 0; i-- {
				result = append(result, a[aIndex])
				aIndex++
			}
		}
		prevState = step.Type
	}

	switch prevState {
	case Added:
		result = append(result, endChange)
	case Removed:
		result = append(result, middleChange)
		result = append(result, endChange)
	}
	return result
}
