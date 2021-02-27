package comparable

var _ Comparable = (*Char)(nil)

// Char is a comparable for two strings.
type Char struct {
	a string
	b string
}

// NewChar constructs a new strings comparable.
//
// WARNING: This does not handle escaped utf-8 sequences,
// this uses Go's normal len and indexing. For full unicode
// support use the Runes comparable or one of the others.
func NewChar(a, b string) *Char {
	return &Char{
		a: a,
		b: b,
	}
}

// ALength is the length of the first list being compared.
func (comp *Char) ALength() int {
	return len(comp.a)
}

// BLength is the length of the second list being compared.
func (comp *Char) BLength() int {
	return len(comp.b)
}

// Equals determines if the entries in the two given indices are equal.
func (comp *Char) Equals(aIndex, bIndex int) bool {
	return comp.a[aIndex] == comp.b[bIndex]
}

// AValue gets the value from the A source at the given index.
func (comp *Char) AValue(aIndex int) byte {
	return comp.a[aIndex]
}

// BValue gets the value from the B source at the given index.
func (comp *Char) BValue(bIndex int) byte {
	return comp.b[bIndex]
}
