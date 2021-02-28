package container

import "github.com/Grant-Nelson/goDiff/internal/collector"

// Min2 gets the minimum value of the two given values.
func Min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Min3 gets the minimum value of the three given values.
func Min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// EndCase determines if the given container is small enough to be simply added
// into the collector without any diff algorithm. This will add into the given
// collector and return true if done, otherwise it will return false.
func (cont *Container) EndCase(col *collector.Collector) bool {
	if cont.aLength <= 1 {
		cont.aEdge(col)
		return true
	}

	if cont.bLength <= 1 {
		cont.bEdge(col)
		return true
	}

	return false
}

// aEdge handles when at the edge of the A source subset in the given container.
func (cont *Container) aEdge(col *collector.Collector) {
	aLen, bLen := cont.aLength, cont.bLength

	if aLen <= 0 {
		col.InsertAdded(bLen)
		return
	}

	split := -1
	for j := 0; j < bLen; j++ {
		if cont.Equals(0, j) { // TODO: Optimise this scan.
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
func (cont *Container) bEdge(col *collector.Collector) {
	aLen, bLen := cont.aLength, cont.bLength

	if bLen <= 0 {
		col.InsertRemoved(aLen)
		return
	}

	split := -1
	for i := 0; i < aLen; i++ {
		if cont.Equals(i, 0) { // TODO: Optimise this scan.
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
