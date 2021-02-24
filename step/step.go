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
