package comparable

var _ Comparable = (*RuneSlice)(nil)

// RuneSlice is a comparable for two string slices.
type RuneSlice struct {
	a [][]rune
	b [][]rune
}

// NewRuneSlice constructs a new string slice comparable.
func NewRuneSlice(a, b [][]rune) *RuneSlice {
	return &RuneSlice{
		a: a,
		b: b,
	}
}

// ALength is the length of the first list being compared.
func (comp *RuneSlice) ALength() int {
	return len(comp.a)
}

// BLength is the length of the second list being compared.
func (comp *RuneSlice) BLength() int {
	return len(comp.b)
}

// Equals determines if the entries in the two given indices are equal.
func (comp *RuneSlice) Equals(aIndex, bIndex int) bool {
	a, b := comp.a[aIndex], comp.b[bIndex]
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

// AValue gets the value from the A source at the given index.
func (comp *RuneSlice) AValue(aIndex int) []rune {
	return comp.a[aIndex]
}

// BValue gets the value from the B source at the given index.
func (comp *RuneSlice) BValue(bIndex int) []rune {
	return comp.b[bIndex]
}
