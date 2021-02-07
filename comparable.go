package diff

// Comparable is a simple interface for generic difference determination.
type Comparable interface {

	// ALength is the length of the first list being compared.
	ALength() int

	// BLength is the length of the second list being compared.
	BLength() int

	// Equals determines if the entries in the two given indices are equal.
	Equals(aIndex, bIndex int) bool
}

var _ Comparable = (*strSliceComp)(nil)

// strSliceComp is a comparable for two string slices.
type strSliceComp struct {
	a []string
	b []string
}

// newStrSliceComp constructs a new string slice comparable.
func newStrSliceComp(a, b []string) *strSliceComp {
	return &strSliceComp{
		a: a,
		b: b,
	}
}

// ALength is the length of the first list being compared.
func (comp *strSliceComp) ALength() int {
	return len(comp.a)
}

// BLength is the length of the second list being compared.
func (comp *strSliceComp) BLength() int {
	return len(comp.b)
}

// Equals determines if the entries in the two given indices are equal.
func (comp *strSliceComp) Equals(aIndex, bIndex int) bool {
	return comp.a[aIndex] == comp.b[bIndex]
}

var _ Comparable = (*runeSliceComp)(nil)

// runeSliceComp is a comparable for two string slices.
type runeSliceComp struct {
	a [][]rune
	b [][]rune
}

// newRuneSliceComp constructs a new string slice comparable.
func newRuneSliceComp(a, b [][]rune) *runeSliceComp {
	return &runeSliceComp{
		a: a,
		b: b,
	}
}

// ALength is the length of the first list being compared.
func (comp *runeSliceComp) ALength() int {
	return len(comp.a)
}

// BLength is the length of the second list being compared.
func (comp *runeSliceComp) BLength() int {
	return len(comp.b)
}

// Equals determines if the entries in the two given indices are equal.
func (comp *runeSliceComp) Equals(aIndex, bIndex int) bool {
	a, b := comp.a[aIndex], comp.b[bIndex]
	if len(a) == len(b) {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
	}
	return true
}
