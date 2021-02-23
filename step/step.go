package step

type (
	// Type is the steps of the levenshtein path.
	Type int

	// PathCallback is the function signature for calling back steps in the path.
	PathCallback func(step Type, count int)
)

const (
	// Equal indicates A and B entries are equal.
	Equal Type = iota

	// Added indicates A was added.
	Added

	// Removed indicates A was removed.
	Removed
)

const (
	// RemoveCost gives the cost to remove A at the given index.
	RemoveCost = 1

	// AddCost gives the cost to add B at the given index.
	AddCost = 1

	// SubstitionCost gives the substition cost for replacing A with B at the given indices.
	SubstitionCost = 2

	// EqualCost gives the cost for A and B being equal.
	EqualCost = 0
)
