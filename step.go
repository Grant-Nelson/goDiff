package diff

// StepType is the steps of the levenshtein path.
type StepType int

const (
	// Equal indicates A and B entries are equal.
	Equal StepType = iota

	// Added indicates A was added.
	Added

	// Removed indicates A was removed.
	Removed
)

// Step is a continuous group of step types.
type Step struct {

	// step is the type for this group.
	Type StepType

	// count is the number of the given type in the group.
	Count int
}
