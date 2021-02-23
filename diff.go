package diff

import (
	"./comparable"
	"./internal/collector"
	"./internal/container"
	"./internal/hirschberg"
	"./step"
)

// Results are the result from a diff.
type Results interface {

	// Count is the number of steps in this diff.
	Count() int

	// Total is the total number of parts represented by this diff.
	// The total sum of all the counts in each step.
	Total() int

	// Read will read the steps to take for this diff.
	Read(hndl step.PathCallback)
}

// defaultWagnerThreshold is the point at which the algorithms switch from Hirschberg
// to Wagner. When both length of the comparable are smaller than this value Wagner
// is used. The Wagner matrix will never be larger than this value of entries.
// If this is less than 4 the Wagner algorithm will not be used.
const defaultWagnerThreshold = 500

// check that the collector can be used as the resulting diff.
var _ Results = (*collector.Collector)(nil)

// Diff will perform a diff on the given comparable information.
func Diff(comp comparable.Comparable) Results {
	col := collector.New()
	h := hirschberg.New(nil)
	cont := container.New(comp)
	h.Diff(cont, col)
	return col
}
