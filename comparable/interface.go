package comparable

var _ Comparable = (*Integer)(nil)

type (
	// EqualTest is the function used to determine if
	// the two given interfaces are equal or not.
	EqualTest func(a, b interface{}) bool

	// Interface is a comparable for two interface slices.
	Interface struct {
		a  []interface{}
		b  []interface{}
		eq EqualTest
	}
)

// NewInterface constructs a new interface slice comparable.
// If the given equal test function is nil then the interfaces will
// be compared with the default equal (==).
func NewInterface(a, b []interface{}, eq EqualTest) *Interface {
	return &Interface{
		a:  a,
		b:  b,
		eq: eq,
	}
}

// ALength is the length of the first list being compared.
func (comp *Interface) ALength() int {
	return len(comp.a)
}

// BLength is the length of the second list being compared.
func (comp *Interface) BLength() int {
	return len(comp.b)
}

// Equals determines if the entries in the two given indices are equal.
func (comp *Interface) Equals(aIndex, bIndex int) bool {
	if comp.eq != nil {
		return comp.eq(comp.a[aIndex], comp.b[bIndex])
	}
	return comp.a[aIndex] == comp.b[bIndex]
}

// AValue gets the value from the A source at the given index.
func (comp *Interface) AValue(aIndex int) interface{} {
	return comp.a[aIndex]
}

// BValue gets the value from the B source at the given index.
func (comp *Interface) BValue(bIndex int) interface{} {
	return comp.b[bIndex]
}
