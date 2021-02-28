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

// DefaultWagnerThreshold is the point at which the algorithms switch from Hirschberg
// to Wagner-Fischer. When both length of the comparable are smaller than this value
// Wagner-Fischer is used. The Wagner matrix will never be larger than this value of entries.
// If this is less than 4 the Wagner algorithm will not be used.
const DefaultWagnerThreshold = 500

// check that the collector can be used as the resulting diff.
var _ Results = (*collector.Collector)(nil)

// wrap wraps an instance of a Diff into an Algorithm.
func wrap(diff container.Diff) Algorithm {
	return func(comp comparable.Comparable) Results {
		col := collector.New()
		cont := container.New(comp)
		cont, before, after := cont.Reduce()
		col.InsertEqual(after)
		if !cont.EndCase(col) {
			diff.Diff(cont, col)
		}
		col.InsertEqual(before)
		col.Finish()
		return col
	}
}

// HirschbergDiff creates a new Hirschberg algorithm instance for performing a diff.
//
// The given length is the initial score vector size. If the vector is too small it will be
// reallocated to the larger size. Use -1 to not preallocate the vectors.
// The useReduce flag indicates if the equal padding edges should be checked
// at each step of the algorithm or not.
func HirschbergDiff(length int, useReduce bool) Algorithm {
	return wrap(hirschberg.New(nil, length, useReduce))
}

// WagnerDiff creates a new Wagner-Fischer algorithm instance for performing a diff.
//
// The given size is the amount of matrix space, width * height, to preallocate
// for the Wagner-Fischer algorithm. Use -1 to not preallocate any matrix.
func WagnerDiff(size int) Algorithm {
	return wrap(wagner.New(size))
}

// HybridDiff creates a new hybrid Hirschberg with Wagner-Fischer cutoff for performing a diff.
//
// The given length is the initial score vector size of the Hirschberg algorithm. If the vector
// is too small it will be reallocated to the larger size. Use -1 to not preallocate the vectors.
// The useReduce flag indicates if the equal padding edges should be checked
// at each step of the algorithm or not.
//
// The given size is the amount of matrix space, width * height, to use for the Wagner-Fischer.
// This must be greater than 4 fo use the cutoff. The larger the size, the more memory is used
// creating the matrix but the earlier the Wagner-Fischer algorithm can take over.
func HybridDiff(length int, useReduce bool, size int) Algorithm {
	return wrap(hirschberg.New(wagner.New(size), length, useReduce))
}

// DefaultDiff creates the default diff algorithm with default configuration.
// The default is a hybrid Hirschberg with Wagner-Fischer using a reduction
// at each step and the default Wagner threshold.
func DefaultDiff() Algorithm {
	return HybridDiff(-1, true, DefaultWagnerThreshold)
}

// Diff will perform a diff on the given comparable information.
// This will use a new instance of the default diff configuration.
func Diff(comp comparable.Comparable) Results {
	return DefaultDiff()(comp)
}
