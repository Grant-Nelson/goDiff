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
		if b < c {
			return b
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func EndCase(cont *Container, col *collector.Collector) bool {
	aLen := cont.ALength()
	if aLen <= 1 {
		aEdge(cont, col)
		return true
	}

	bLen := cont.BLength()
	if bLen <= 1 {
		bEdge(cont, col)
		return true
	}

	return false
}

// aEdge handles when at the edge of the A source subset in the given container.
func aEdge(cont *Container, col *collector.Collector) {
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
func bEdge(cont *Container, col *collector.Collector) {
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
