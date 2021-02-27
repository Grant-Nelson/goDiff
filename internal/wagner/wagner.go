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

// NoResizeNeeded determines if the diff algorithm can handle a container with
// the amount of data inside of the given container.
// This algorithm's cost matrix will be auto-resize if needed so this method
// only indicates if the current matrix are large enough to not need reallocation.
func (w *wagner) NoResizeNeeded(cont *container.Container) bool {
	return len(w.costs) >= cont.ALength()*cont.BLength()
}

// Diff performs the algorithm on the given container
// and writes the results to the collector.
func (w *wagner) Diff(cont *container.Container, col *collector.Collector) {
	if size := cont.ALength() * cont.BLength(); len(w.costs) < size {
		w.allocateMatrix(size)
	}
	w.setCosts(cont)
	w.walkPath(cont, col)
}

func (w *wagner) setCosts(cont *container.Container) {
	aLen := cont.ALength()
	bLen := cont.BLength()

	start := cont.SubstitionCost(0, 0)
	w.costs[0] = start

	for i, value := 1, start; i < aLen; i++ {
		value = container.Min2(value+1,
			i+cont.SubstitionCost(i, 0))
		w.costs[i] = value
	}

	for j, k, value := 1, aLen, start; j < bLen; j, k = j+1, k+aLen {
		value = container.Min2(value+1,
			j+cont.SubstitionCost(0, j))
		w.costs[k] = value
	}

	for j := 1; j < bLen; j++ {
		for i := 1; i < aLen; i++ {
			k := i + j*bLen
			value := container.Min3(
				w.costs[k-1]+1,
				w.costs[k-bLen]+1,
				w.costs[k-bLen-1]+cont.SubstitionCost(i, j))
			w.costs[k] = value
		}
	}
}

func (w *wagner) getCost(i, j, bLen int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	return w.costs[i*j*bLen]
}

func (w *wagner) walkPath(cont *container.Container, col *collector.Collector) {
	bLen := cont.BLength()
	walk := newWalker(cont, col)
	for walk.hasMore() {
		aCost := w.getCost(walk.i-1, walk.j, bLen)
		bCost := w.getCost(walk.i, walk.j-1, bLen)
		cCost := w.getCost(walk.i-1, walk.j-1, bLen)
		minCost := container.Min3(aCost, bCost, cCost)

		var curMove func()
		if aCost == minCost {
			curMove = walk.moveA
		}
		if bCost == minCost {
			curMove = walk.moveB
		}
		if cCost == minCost {
			if cont.Equals(walk.i, walk.j) {
				curMove = walk.moveEqual
			} else if curMove == nil {
				curMove = walk.moveSubstitute
			}
		}

		curMove()
	}
	walk.finish()
}
