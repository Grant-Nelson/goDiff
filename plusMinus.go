package diff

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
