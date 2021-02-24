package comparable

var _ Comparable = (*String)(nil)

// String is a comparable for two string slices.
type String struct {
	a []string
	b []string
}

// NewString constructs a new string slice comparable.
func NewString(a, b []string) *String {
	return &String{
		a: a,
		b: b,
	}
}

// ALength is the length of the first list being compared.
func (comp *String) ALength() int {
	return len(comp.a)
}

// BLength is the length of the second list being compared.
func (comp *String) BLength() int {
	return len(comp.b)
}

// Equals determines if the entries in the two given indices are equal.
func (comp *String) Equals(aIndex, bIndex int) bool {
	return comp.a[aIndex] == comp.b[bIndex]
}

// AValue gets the value from the A source at the given index.
func (comp *String) AValue(aIndex int) string {
	return comp.a[aIndex]
}

// BValue gets the value from the B source at the given index.
func (comp *String) BValue(bIndex int) string {
	return comp.b[bIndex]
}
