package container

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Grant-Nelson/goDiff/comparable"
)

func Test_Min2(t *testing.T) {
	checkMin2(t, 0, 0, 0)
	checkMin2(t, 0, 2, 0)
	checkMin2(t, 2, 0, 0)
	checkMin2(t, 2, 2, 2)
	checkMin2(t, 42, -42, -42)
	checkMin2(t, -42, 42, -42)
}

func Test_Min3(t *testing.T) {
	checkMin3(t, 0, 0, 0, 0)
	checkMin3(t, 0, 0, 2, 0)
	checkMin3(t, 0, 2, 0, 0)
	checkMin3(t, 2, 0, 0, 0)
	checkMin3(t, 4, 2, 0, 0)
	checkMin3(t, 2, 4, 0, 0)
	checkMin3(t, 4, 0, 2, 0)
	checkMin3(t, 2, 0, 4, 0)
	checkMin3(t, 0, 4, 2, 0)
	checkMin3(t, 0, 2, 4, 0)
	checkMin3(t, 2, 2, 2, 2)
	checkMin3(t, 42, -42, 0, -42)
	checkMin3(t, -42, 0, 42, -42)
	checkMin3(t, 0, 42, -42, -42)
}

func Test_Equals(t *testing.T) {
	cont := newCont(`cat`, `kitten`)
	check(t, cont, `cat`, `kitten`)
	intEqual(t, cont.ALength(), 3, `ALength`)
	intEqual(t, cont.BLength(), 6, `BLength`)
	boolEqual(t, cont.Equals(0, 0), false, `Equal(0, 0)`)
	boolEqual(t, cont.Equals(2, 2), true, `Equal(2, 2)`)
	boolEqual(t, cont.Equals(2, 3), true, `Equal(2, 3)`)
	boolEqual(t, cont.Equals(1, 2), false, `Equal(1, 2)`)
	intEqual(t, cont.SubstitionCost(0, 0), SubstitionCost, `SubstitionCost(0, 0)`)
	intEqual(t, cont.SubstitionCost(2, 2), EqualCost, `SubstitionCost(2, 2)`)
}

func Test_Equals_Reversed(t *testing.T) {
	cont := reverse(newCont(`cat`, `kitten`))
	check(t, cont, `tac`, `nettik`)
	intEqual(t, cont.ALength(), 3, `ALength`)
	intEqual(t, cont.BLength(), 6, `BLength`)
	boolEqual(t, cont.Equals(0, 0), false, `Equal(0, 0)`)
	boolEqual(t, cont.Equals(0, 2), true, `Equal(0, 2)`)
	boolEqual(t, cont.Equals(0, 3), true, `Equal(0, 3)`)
	boolEqual(t, cont.Equals(1, 2), false, `Equal(1, 2)`)
	intEqual(t, cont.SubstitionCost(0, 0), SubstitionCost, `SubstitionCost(0, 0)`)
	intEqual(t, cont.SubstitionCost(0, 2), EqualCost, `SubstitionCost(2, 2)`)
}

func Test_Sub(t *testing.T) {
	cont := newCont(`abcdef`, `ghi`)
	check(t, cont, `abcdef`, `ghi`)
	subCheck(t, cont, 0, 3, 0, 3, false, `abc`, `ghi`)
	subCheck(t, cont, 1, 4, 1, 3, false, `bcd`, `hi`)
	subCheck(t, cont, 0, 3, 0, 3, true, `cba`, `ihg`)
	subCheck(t, cont, 2, 5, 1, 3, true, `edc`, `ih`)
}

func Test_Sub_Reversed(t *testing.T) {
	cont := reverse(newCont(`abcdef`, `ghi`))
	check(t, cont, `fedcba`, `ihg`)
	subCheck(t, cont, 0, 3, 0, 3, false, `fed`, `ihg`)
	subCheck(t, cont, 1, 4, 1, 3, false, `edc`, `hg`)
	subCheck(t, cont, 0, 3, 0, 3, true, `def`, `ghi`)
	subCheck(t, cont, 2, 5, 1, 3, true, `bcd`, `gh`)
}

func Test_Reduce(t *testing.T) {
	reduceCheck(t, newCont(`abc`, `abc`), ``, ``, 3, 0)
	reduceCheck(t, newCont(`abc`, `def`), `abc`, `def`, 0, 0)
	reduceCheck(t, newCont(`abc`, `aef`), `bc`, `ef`, 1, 0)
	reduceCheck(t, newCont(`abc`, `dec`), `ab`, `de`, 0, 1)
	reduceCheck(t, newCont(`abc`, `ac`), `b`, ``, 1, 1)
	reduceCheck(t, newCont(`ac`, `abc`), ``, `b`, 1, 1)
	reduceCheck(t, newCont(`abcd`, `acd`), `b`, ``, 1, 2)
	reduceCheck(t, newCont(`abcd`, `abd`), `c`, ``, 2, 1)
	reduceCheck(t, newCont(`abc`, ``), `abc`, ``, 0, 0)
	reduceCheck(t, newCont(``, `abc`), ``, `abc`, 0, 0)
}

func Test_Reduce_Reversed(t *testing.T) {
	reduceCheck(t, reverse(newCont(`abc`, `abc`)), ``, ``, 0, 3)
	reduceCheck(t, reverse(newCont(`abc`, `def`)), `cba`, `fed`, 0, 0)
	reduceCheck(t, reverse(newCont(`abc`, `aef`)), `cb`, `fe`, 0, 1)
	reduceCheck(t, reverse(newCont(`abc`, `dec`)), `ba`, `ed`, 1, 0)
	reduceCheck(t, reverse(newCont(`abc`, `ac`)), `b`, ``, 1, 1)
	reduceCheck(t, reverse(newCont(`ac`, `abc`)), ``, `b`, 1, 1)
	reduceCheck(t, reverse(newCont(`abcd`, `acd`)), `b`, ``, 2, 1)
	reduceCheck(t, reverse(newCont(`abcd`, `abd`)), `c`, ``, 1, 2)
	reduceCheck(t, reverse(newCont(`abc`, ``)), `cba`, ``, 0, 0)
	reduceCheck(t, reverse(newCont(``, `abc`)), ``, `cba`, 0, 0)
}

func checkMin2(t *testing.T, a, b, exp int) {
	if result := Min2(a, b); result != exp {
		t.Error(fmt.Sprint(
			"Unexpected minimum value:",
			"\n   Input:  ", a, ", ", b,
			"\n   Result: ", result, " => ", exp))
	}
}

func checkMin3(t *testing.T, a, b, c, exp int) {
	if result := Min3(a, b, c); result != exp {
		t.Error(fmt.Sprint(
			"Unexpected minimum value:",
			"\n   Input:  ", a, ", ", b, ", ", c,
			"\n   Result: ", result, " => ", exp))
	}
}

func boolEqual(t *testing.T, value, exp bool, msg string) {
	if value != exp {
		t.Error(fmt.Sprint("Unexpected boolean value:",
			"\n   Message:  ", msg,
			"\n   Value:    ", value,
			"\n   Expected: ", exp))
	}
}

func intEqual(t *testing.T, value, exp int, msg string) {
	if value != exp {
		t.Error(fmt.Sprint("Unexpected integer value:",
			"\n   Message:  ", msg,
			"\n   Value:    ", value,
			"\n   Expected: ", exp))
	}
}

func strEqual(t *testing.T, value, exp, msg string) {
	if value != exp {
		t.Error(fmt.Sprint("Unexpected string value:",
			"\n   Message:  ", msg,
			"\n   Value:    ", value,
			"\n   Expected: ", exp))
	}
}

func newCont(a, b string) *Container {
	return New(comparable.NewChar(a, b))
}

func reverse(c *Container) *Container {
	return c.Sub(0, c.aLength, 0, c.bLength, !c.reverse)
}

func check(t *testing.T, cont *Container, expA, expB string) {
	resultA := strings.Join(cont.AParts(), ``)
	resultB := strings.Join(cont.BParts(), ``)
	if (resultA != expA) || (resultB != expB) {
		t.Error(fmt.Sprint(
			"Unexpected resulting container:",
			"\n   Container:  ", cont,
			"\n   Result A:   ", resultA, " => ", expA,
			"\n   Result B:   ", resultB, " => ", expB))
	}
	return
}

func subCheck(t *testing.T, cont *Container, aLow, aHigh, bLow, bHigh int, reverse bool, expA, expB string) *Container {
	sub := cont.Sub(aLow, aHigh, bLow, bHigh, reverse)
	resultA := strings.Join(sub.AParts(), ``)
	resultB := strings.Join(sub.BParts(), ``)
	if (resultA != expA) || (resultB != expB) {
		t.Error(fmt.Sprint(
			"Unexpected results from Reduce:",
			"\n   Original: ", cont,
			"\n   Sub:      ", sub,
			"\n   A Parts:  ", resultA, " => ", expA,
			"\n   B Parts:  ", resultB, " => ", expB))
	}
	return sub
}

func reduceCheck(t *testing.T, cont *Container, expA, expB string, expBefore, expAfter int) *Container {
	sub, before, after := cont.Reduce()
	resultA := strings.Join(sub.AParts(), ``)
	resultB := strings.Join(sub.BParts(), ``)
	if (before != expBefore) || (after != expAfter) || (resultA != expA) || (resultB != expB) {
		t.Error(fmt.Sprint(
			"Unexpected results from Reduce:",
			"\n   Original: ", cont,
			"\n   Reduces:  ", sub,
			"\n   A Parts:  ", resultA, " => ", expA,
			"\n   B Parts:  ", resultB, " => ", expB,
			"\n   Before:   ", before, " => ", expBefore,
			"\n   After:    ", after, " => ", expAfter))
	}
	return sub
}
