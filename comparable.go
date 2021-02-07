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
