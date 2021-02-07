package diff

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
