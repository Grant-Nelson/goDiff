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

func getAParts(cont *Container) string {
	comp := cont.comp.(*comparable.String)
	subA := make([]string, cont.aLength)
	for i := 0; i < cont.aLength; i++ {
		subA[i] = comp.AValue(cont.aAdjust(i))
	}
	return strings.Join(subA, ` `)
}

func getBParts(cont *Container) string {
	comp := cont.comp.(*comparable.String)
	subB := make([]string, cont.bLength)
	for j := 0; j < cont.bLength; j++ {
		subB[j] = comp.BValue(cont.bAdjust(j))
	}
	return strings.Join(subB, ` `)
}

func check(t *testing.T, cont *Container, expA, expB string) {
	resultA, resultB := getAParts(cont), getBParts(cont)
	if (resultA != expA) || (resultB != expB) {
		t.Error(fmt.Sprint(
			"Unexpected resulting container:\n",
			"Container:  ", cont.aOffset, ", ", cont.aLength, ", ", cont.bOffset, ", ", cont.bLength, ", ", cont.reverse, "\n",
			"Result A:   ", resultA, " => ", expA, "\n",
			"Result B:   ", resultB, " => ", expB, "\n"))
	}
	return
}

func reduceCheck(t *testing.T, cont *Container, expA, expB string, expBefore, expAfter int) *Container {
	sub, before, after := cont.Reduce()
	resultA, resultB := getAParts(sub), getBParts(sub)
	if (before != expBefore) || (after != expAfter) || (resultA != expA) || (resultB != expB) {
		t.Error(fmt.Sprint(
			"Unexpected results from Reduce:\n",
			"Original: ", cont.aOffset, ", ", cont.aLength, ", ", cont.bOffset, ", ", cont.bLength, ", ", cont.reverse, "\n",
			"Reduces:  ", sub.aOffset, ", ", sub.aLength, ", ", sub.bOffset, ", ", sub.bLength, ", ", sub.reverse, "\n",
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
