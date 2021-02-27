package wagner

import (
	"fmt"
	"testing"

	"github.com/Grant-Nelson/goDiff/comparable"
	"github.com/Grant-Nelson/goDiff/internal/container"
)

func Test_CostMatrix(t *testing.T) {
	w := New(-1).(*wagner)
	cont := container.New(comparable.NewChar(`kitten`, `sitting`))
	w.allocateMatrix(cont.ALength() * cont.BLength())
	for k := 0; k < len(w.costs); k++ {
		w.costs[k] = k
	}

	fmt.Println(w.costsString(cont.ALength(), cont.BLength()))

	w.setCosts(cont)
	fmt.Println(w.costsString(cont.ALength(), cont.BLength()))
	t.Fail()
}
