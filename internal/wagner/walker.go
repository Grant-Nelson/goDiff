package wagner

import (
	"github.com/Grant-Nelson/goDiff/internal/collector"
	"github.com/Grant-Nelson/goDiff/internal/container"
)

type walker struct {
	col *collector.Collector
	i   int
	j   int
}

func newWalker(cont *container.Container, col *collector.Collector) *walker {
	return &walker{
		col: col,
		i:   cont.ALength() - 1,
		j:   cont.BLength() - 1,
	}
}

func (w *walker) hasMore() bool {
	return w.i >= 0 && w.j >= 0
}

func (w *walker) moveA() {
	w.i--
	w.col.InsertRemoved(1)
}

func (w *walker) moveB() {
	w.j--
	w.col.InsertAdded(1)
}

func (w *walker) moveEqual() {
	w.i--
	w.j--
	w.col.InsertEqual(1)
}

func (w *walker) moveSubstitute() {
	w.i--
	w.j--
	w.col.InsertSubstitute(1)
}

func (w *walker) finish() {
	w.col.InsertRemoved(w.i)
	w.col.InsertAdded(w.j)
}
