package container

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Grant-Nelson/goDiff/comparable"
)

func newCont(a, b string) *Container {
	return New(comparable.NewString(strings.Split(a, ` `), strings.Split(b, ` `)))
}

func check(t *testing.T, cont *Container, expA, expB string) {
	resultA := strings.Join(cont.AParts(), ` `)
	resultB := strings.Join(cont.BParts(), ` `)
	if (resultA != expA) || (resultB != expB) {
		t.Error(fmt.Sprint(
			"Unexpected resulting container:\n",
			"Container:  ", cont, "\n",
			"Result A:   ", resultA, " => ", expA, "\n",
			"Result B:   ", resultB, " => ", expB, "\n"))
	}
	return
}

func reduceCheck(t *testing.T, cont *Container, expA, expB string, expBefore, expAfter int) *Container {
	sub, before, after := cont.Reduce()
	resultA := strings.Join(sub.AParts(), ` `)
	resultB := strings.Join(sub.BParts(), ` `)
	if (before != expBefore) || (after != expAfter) || (resultA != expA) || (resultB != expB) {
		t.Error(fmt.Sprint(
			"Unexpected results from Reduce:\n",
			"Original: ", cont, "\n",
			"Reduces:  ", sub, "\n",
			"A Parts:  ", resultA, " => ", expA, "\n",
			"B Parts:  ", resultB, " => ", expB, "\n",
			"Before:   ", before, " => ", expBefore, "\n",
			"After:    ", after, " => ", expAfter, "\n"))
	}
	return sub
}

func Test_Reduce(t *testing.T) {
	reduceCheck(t, newCont(`a b c`, `a b c`), ``, ``, 3, 0)
	reduceCheck(t, newCont(`a b c`, `d e f`), `a b c`, `d e f`, 0, 0)

	reduceCheck(t, newCont(`a b c`, `a e f`), `b c`, `e f`, 1, 0)
	reduceCheck(t, newCont(`a b c`, `d e c`), `a b`, `d e`, 0, 1)

	reduceCheck(t, newCont(`a b c`, `a c`), `b`, ``, 1, 1)
	reduceCheck(t, newCont(`a c`, `a b c`), ``, `b`, 1, 1)

	reduceCheck(t, newCont(`a b c d`, `a c d`), `b`, ``, 1, 2)
	reduceCheck(t, newCont(`a b c d`, `a b d`), `c`, ``, 2, 1)

	reduceCheck(t, newCont(`a b c`, ``), `a b c`, ``, 0, 0)
	reduceCheck(t, newCont(``, `a b c`), ``, `a b c`, 0, 0)
}

// TODO: Add more tests including tests of sub and reduce/reverse
