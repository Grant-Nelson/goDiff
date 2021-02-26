package wagner

import (
	"github.com/Grant-Nelson/goDiff/internal/collector"
	"github.com/Grant-Nelson/goDiff/internal/container"
)

// Wagner–Fischer algorithm
// (https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm)
type wagner struct {
	costs []int
}

// New creates a new Wagner–Fischer diff algorithm.
//
// The given size is the amount of matrix space, width * height, to preallocate
// for the Wagner-Fischer algorithm. Use -1 to not preallocate any matrix.
func New(size int) container.Diff {
	w := &wagner{}
	if size > 0 {
		w.allocateMatrix(size)
	}
	return w
}

// allocateVectors will create the slice to used for the costs matrix.
func (w *wagner) allocateMatrix(size int) {
	w.costs = make([]int, size)
}

// WillResize determines if the diff algorithm can handle a container with
// the amount of data inside of the given container.
// This algorithm's cost matrix will be auto-resize if needed so this method
// only indicates if the current matrix are large enough to not need reallocation.
func (w *wagner) WillResize(cont *container.Container) bool {
	return len(w.costs) >= cont.ALength()*cont.BLength()
}

// Diff performs the algorithm on the given container
// and writes the results to the collector.
func (w *wagner) Diff(cont *container.Container, col *collector.Collector) {

}
