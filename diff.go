package godiff

import (
	"github.com/Grant-Nelson/goDiff/comparable"
	"github.com/Grant-Nelson/goDiff/internal/collector"
	"github.com/Grant-Nelson/goDiff/internal/container"
	"github.com/Grant-Nelson/goDiff/internal/hirschberg"
	"github.com/Grant-Nelson/goDiff/internal/wagner"
	"github.com/Grant-Nelson/goDiff/step"
)

type (
	// Results are the result from a diff.
	Results interface {

		// Count is the number of steps in this diff.
		Count() int

		// Total is the total number of parts represented by this diff.
		// The total sum of all the counts in each step.
		Total() int

		// Read will read the steps to take for this diff.
		Read(hndl step.PathCallback)
	}

	// Algorithm is an instance of a diff algorithm configuration which can be used
	// multiple times for different input. This can help reduce memory pressure by
	// reusing already allocated buffers.
	Algorithm func(comp comparable.Comparable) Results
)

// defaultWagnerThreshold is the point at which the algorithms switch from Hirschberg
// to Wagner. When both length of the comparable are smaller than this value Wagner
// is used. The Wagner matrix will never be larger than this value of entries.
// If this is less than 4 the Wagner algorithm will not be used.
const defaultWagnerThreshold = 500

// check that the collector can be used as the resulting diff.
var _ Results = (*collector.Collector)(nil)

// hirschbergDiff creates a new hirschberg algorithm instance for performing a diff.
func hirschbergDiff(length int, useReduce bool) Algorithm {
	h := hirschberg.New(nil, length, useReduce)
	return func(comp comparable.Comparable) Results {
		col := collector.New()
		cont := container.New(comp)
		h.Diff(cont, col)
		col.Finish()
		return col
	}
}

func wagnerDiff(size int) Algorithm {
	w := wagner.New(size)
	return func(comp comparable.Comparable) Results {
		col := collector.New()
		cont := container.New(comp)
		w.Diff(cont, col)
		col.Finish()
		return col
	}
}

// Diff will perform a diff on the given comparable information.
func Diff(comp comparable.Comparable) Results {
	return wagnerDiff(-1)(comp)
}
