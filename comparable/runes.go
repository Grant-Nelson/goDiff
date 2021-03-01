package comparable

var _ Comparable = (*Runes)(nil)

// Runes is a comparable for two runes.
type Runes struct {
	a []rune
	b []rune
}

// NewRunes constructs a new runes comparable.
func NewRunes(a, b []rune) *Runes {
	return &Runes{
		a: a,
		b: b,
	}
}

// ALength is the length of the first list being compared.
func (comp *Runes) ALength() int {
	return len(comp.a)
}

// BLength is the length of the second list being compared.
func (comp *Runes) BLength() int {
	return len(comp.b)
}

// Equals determines if the entries in the two given indices are equal.
func (comp *Runes) Equals(aIndex, bIndex int) bool {
	return comp.a[aIndex] == comp.b[bIndex]
}

// AValue gets the value from the A source at the given index.
func (comp *Runes) AValue(aIndex int) rune {
	return comp.a[aIndex]
}

// BValue gets the value from the B source at the given index.
func (comp *Runes) BValue(bIndex int) rune {
	return comp.b[bIndex]
}
