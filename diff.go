package diff

import (
	"./comparable"
	"./internal/collector"
	"./internal/hirschberg"
)

type (
	// StepType is the steps of the levenshtein path.
	StepType int

	// PathCallback is the function signature for calling back steps in the path.
	PathCallback func(step StepType, count int)

	// Results are the result from a diff.
	Results interface {

		// Count is the number of steps in this diff.
		Count() int

		// Total is the total number of parts represented by this diff.
		// The total sum of all the counts in each step.
		Total() int

		// Read will read the steps to take for this diff.
		Read(hndl PathCallback)
	}
)

const (
	// Equal indicates A and B entries are equal.
	Equal = StepType(collector.Equal)

	// Added indicates A was added.
	Added = StepType(collector.Added)

	// Removed indicates A was removed.
	Removed = StepType(collector.Removed)

	// defaultWagnerThreshold is the point at which the algorithms switch from Hirschberg
	// to Wagner. When both length of the comparable are smaller than this value Wagner
	// is used. The Wagner matrix will never be larger than this value of entries.
	// If this is less than 4 the Wagner algorithm will not be used.
	defaultWagnerThreshold = 500
)

// check that the collector can be used as the resulting diff.
var _ Results = (*collector.Collector)(nil)

// Diff will perform a diff on the given comparable information.
func Diff(comp comparable.Comparable) Results {
	col := collector.New()
	h := hirschberg.New(nil)
	h.Diff(comp, col)
	return col
}
