package comparable

var _ Comparable = (*Integer)(nil)

// Integer is a comparable for two integer slices.
type Integer struct {
	a []int
	b []int
}

// NewInteger constructs a new integer slice comparable.
func NewInteger(a, b []int) *Integer {
	return &Integer{
		a: a,
		b: b,
	}
}

// ALength is the length of the first list being compared.
func (comp *Integer) ALength() int {
	return len(comp.a)
}

// BLength is the length of the second list being compared.
func (comp *Integer) BLength() int {
	return len(comp.b)
}

// Equals determines if the entries in the two given indices are equal.
func (comp *Integer) Equals(aIndex, bIndex int) bool {
	return comp.a[aIndex] == comp.b[bIndex]
}

// AValue gets the value from the A source at the given index.
func (comp *Integer) AValue(aIndex int) int {
	return comp.a[aIndex]
}

// BValue gets the value from the B source at the given index.
func (comp *Integer) BValue(bIndex int) int {
	return comp.b[bIndex]
}
