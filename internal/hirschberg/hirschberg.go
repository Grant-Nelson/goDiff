package hirschberg

import (
	"github.com/Grant-Nelson/goDiff/internal/collector"
	"github.com/Grant-Nelson/goDiff/internal/container"
)

// hirschberg will perform a hybrid Hirschberg/Wagner diff on the given comparable.
// The base algorithm is a Hirschberg's algorithm (https://en.wikipedia.org/wiki/Hirschberg%27s_algorithm)
// used to divide the problem space until the threshold is reached to switch to Wagner.
type hirschberg struct {
	scores    *scores
	hybrid    container.Diff
	useReduce bool
}

// New creates a new Hirschberg diff algorithm.
//
// This allows for an optional diff to use when possible to hybrid the algorithm, to not use
// the optional diff pass in nil. The hybrid is used if it has enough memory preallocated,
// WillResize returns true, otherwise Hirschberg will continue to divide the space until
// the hybrid can be used without causing it to reallocate memory.
//
// The given length is the initial score vector size. If the vector is too small it will be
// reallocated to the larger size. Use -1 to not preallocate the vectors.
func New(hybrid container.Diff, length int) container.Diff {
	return &hirschberg{
		scores:    newScores(length),
		hybrid:    hybrid,
		useReduce: true,
	}
}

// WillResize determines if the diff algorithm can handle a container with
// the amount of data inside of the given container.
// This algorithm's score vectors will be auto-resize if needed so this method
// only indicates if the current vectors are large enough to not need reallocation.
func (h *hirschberg) WillResize(cont *container.Container) bool {
	return len(h.scores.back) >= cont.BLength()+1
}

// Diff performs the algorithm on the given container
// and writes the results to the collector.
func (h *hirschberg) Diff(cont *container.Container, col *collector.Collector) {
	stack := NewStack()
	stack.Push(cont, 0)

	for stack.NotEmpty() {
		cur, remainder := stack.Pop()
		col.InsertEqual(remainder)
		if cur == nil {
			continue
		}

		if h.useReduce {
			var before, after int
			cur, before, after = cur.Reduce()
			col.InsertEqual(after)
			stack.Push(nil, before)
		}

		bLen := cur.BLength()
		if bLen <= 1 {
			bEdge(cur, col)
			continue
		}

		aLen := cur.ALength()
		if aLen <= 1 {
			aEdge(cur, col)
			continue
		}

		if (h.hybrid != nil) && h.hybrid.WillResize(cur) {
			h.hybrid.Diff(cur, col)
			continue
		}

		aMid, bMid := h.scores.Split(cur)
		stack.Push(cur.Sub(0, aMid, 0, bMid, false), 0)
		stack.Push(cur.Sub(aMid, aLen, bMid, bLen, false), 0)
	}
}

// aEdge handles when at the edge of the A source subset in the given container.
func aEdge(cont *container.Container, col *collector.Collector) {
	aLen := cont.ALength()
	bLen := cont.BLength()

	if aLen <= 0 {
		col.InsertAdded(bLen)
		return
	}

	split := -1
	for j := 0; j < bLen; j++ {
		if cont.Equals(0, j) {
			split = j
			break
		}
	}

	if split < 0 {
		col.InsertAdded(bLen)
		col.InsertRemoved(1)
	} else {
		col.InsertAdded(bLen - split - 1)
		col.InsertEqual(1)
		col.InsertAdded(split)
	}
}

// bEdge Handles when at the edge of the B source subset in the given container.
func bEdge(cont *container.Container, col *collector.Collector) {
	aLen := cont.ALength()
	bLen := cont.BLength()

	if bLen <= 0 {
		col.InsertRemoved(aLen)
		return
	}

	split := -1
	for i := 0; i < aLen; i++ {
		if cont.Equals(i, 0) {
			split = i
			break
		}
	}

	if split < 0 {
		col.InsertAdded(1)
		col.InsertRemoved(aLen)
	} else {
		col.InsertRemoved(aLen - split - 1)
		col.InsertEqual(1)
		col.InsertRemoved(split)
	}
}
