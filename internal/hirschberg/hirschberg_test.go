package hirschberg

import (
	"fmt"
	"testing"

	"github.com/Grant-Nelson/goDiff/comparable"
	"github.com/Grant-Nelson/goDiff/internal/collector"
	"github.com/Grant-Nelson/goDiff/internal/container"
)

func Test_Hirschberg_NoReduce(t *testing.T) {
	checkAll(t, New(nil, -1, false))
}

func Test_Hirschberg_UseReduce(t *testing.T) {
	checkAll(t, New(nil, -1, true))
}

func Test_Hirschberg_Hybrid(t *testing.T) {
	d := New(New(nil, 6, false), -1, false)
	check(t, d, `kitten kitten kitten`, `sitting sitting sitting`,
		`-1 +1 =3 -1 +1 =1 +1 =1 -1 +1 =3 -1 +1 =1 +1 =1 -1 +1 =3 -1 +1 =1 +1`)
	check(t, d, `saturday saturday saturday`, `sunday sunday sunday`,
		`=1 -2 =1 -1 +1 =5 -2 =1 -1 +1 =5 -2 =1 -1 +1 =3`)
	check(t, d, `satxrday satxrday satxrday`, `sunday sunday sunday`,
		`=1 -4 +2 =5 -4 +2 =5 -4 +2 =3`)
}

func Test_Stack(t *testing.T) {
	s := NewStack()
	intEqual(t, countNodes(s.top), 0, `top count`)
	intEqual(t, countNodes(s.graveyard), 0, `graveyard count`)
	boolEqual(t, s.NotEmpty(), false, `not empty`)

	s.Push(nil, 1)
	intEqual(t, countNodes(s.top), 1, `top count`)
	intEqual(t, countNodes(s.graveyard), 0, `graveyard count`)
	boolEqual(t, s.NotEmpty(), true, `not empty`)

	s.Push(nil, 2)
	intEqual(t, countNodes(s.top), 2, `top count`)
	intEqual(t, countNodes(s.graveyard), 0, `graveyard count`)
	boolEqual(t, s.NotEmpty(), true, `not empty`)

	s.Push(nil, 3)
	intEqual(t, countNodes(s.top), 3, `top count`)
	intEqual(t, countNodes(s.graveyard), 0, `graveyard count`)
	boolEqual(t, s.NotEmpty(), true, `not empty`)

	_, remainder := s.Pop()
	intEqual(t, remainder, 3, `popped remainder`)
	intEqual(t, countNodes(s.top), 2, `top count`)
	intEqual(t, countNodes(s.graveyard), 1, `graveyard count`)
	boolEqual(t, s.NotEmpty(), true, `not empty`)

	_, remainder = s.Pop()
	intEqual(t, remainder, 2, `popped remainder`)
	intEqual(t, countNodes(s.top), 1, `top count`)
	intEqual(t, countNodes(s.graveyard), 2, `graveyard count`)
	boolEqual(t, s.NotEmpty(), true, `not empty`)

	_, remainder = s.Pop()
	intEqual(t, remainder, 1, `popped remainder`)
	intEqual(t, countNodes(s.top), 0, `top count`)
	intEqual(t, countNodes(s.graveyard), 3, `graveyard count`)
	boolEqual(t, s.NotEmpty(), false, `not empty`)

	_, remainder = s.Pop()
	intEqual(t, remainder, 0, `popped remainder`)
	intEqual(t, countNodes(s.top), 0, `top count`)
	intEqual(t, countNodes(s.graveyard), 3, `graveyard count`)
	boolEqual(t, s.NotEmpty(), false, `not empty`)

	s.Push(nil, 5)
	intEqual(t, countNodes(s.top), 1, `top count`)
	intEqual(t, countNodes(s.graveyard), 2, `graveyard count`)
	boolEqual(t, s.NotEmpty(), true, `not empty`)
}

func countNodes(node *stackNode) int {
	count := 0
	for ; node != nil; count++ {
		node = node.prev
	}
	return count
}

func checkAll(t *testing.T, d container.Diff) {
	check(t, d, `A`, `A`, `=1`)
	check(t, d, `A`, `B`, `-1 +1`)
	check(t, d, `A`, `AB`, `=1 +1`)
	check(t, d, `A`, `BA`, `+1 =1`)
	check(t, d, `AB`, `A`, `=1 -1`)
	check(t, d, `BA`, `A`, `-1 =1`)
	check(t, d, `kitten`, `sitting`, `-1 +1 =3 -1 +1 =1 +1`)
	check(t, d, `saturday`, `sunday`, `=1 -2 =1 -1 +1 =3`)
	check(t, d, `satxrday`, `sunday`, `=1 -4 +2 =3`)
	check(t, d, `ABC`, `ADB`, `=1 +1 =1 -1`)
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

// checks the levenshtein distance algorithm
func check(t *testing.T, d container.Diff, a, b, exp string) {
	col := collector.New()
	cont := container.New(comparable.NewChar(a, b))
	d.Diff(cont, col)
	col.Finish()
	if result := col.String(); exp != result {
		t.Error("Hirschberg returned unexpected result:",
			"\n   Input A:  ", a,
			"\n   Input B:  ", b,
			"\n   Expected: ", exp,
			"\n   Result:   ", result)
	}
}
