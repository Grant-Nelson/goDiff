package comparable

import (
	"fmt"
	"testing"
)

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

func Test_String(t *testing.T) {
	comp := NewString(
		[]string{`a`, `b`, `c`},
		[]string{`a`, `d`})
	intEqual(t, comp.ALength(), 3, `String.ALength`)
	intEqual(t, comp.BLength(), 2, `String.BLength`)
	boolEqual(t, comp.Equals(0, 0), true, `String.Equals(0, 0)`)
	boolEqual(t, comp.Equals(0, 1), false, `String.Equals(0, 1)`)
	boolEqual(t, comp.Equals(1, 0), false, `String.Equals(1, 0)`)
	boolEqual(t, comp.Equals(1, 1), false, `String.Equals(1, 1)`)
	strEqual(t, comp.AValue(1), `b`, `String.AValue(1)`)
	strEqual(t, comp.BValue(1), `d`, `String.BValue(1)`)
}

func Test_Char(t *testing.T) {
	comp := NewChar(`abc`, `ad`)
	intEqual(t, comp.ALength(), 3, `Char.ALength`)
	intEqual(t, comp.BLength(), 2, `Char.BLength`)
	boolEqual(t, comp.Equals(0, 0), true, `Char.Equals(0, 0)`)
	boolEqual(t, comp.Equals(0, 1), false, `Char.Equals(0, 1)`)
	boolEqual(t, comp.Equals(1, 0), false, `Char.Equals(1, 0)`)
	boolEqual(t, comp.Equals(1, 1), false, `Char.Equals(1, 1)`)
	intEqual(t, int(comp.AValue(1)), 'b', `Char.AValue(1)`)
	intEqual(t, int(comp.BValue(1)), 'd', `Char.BValue(1)`)
}

func Test_Runes(t *testing.T) {
	comp := NewRunes([]rune(`abc`), []rune(`ad`))
	intEqual(t, comp.ALength(), 3, `Runes.ALength`)
	intEqual(t, comp.BLength(), 2, `Runes.BLength`)
	boolEqual(t, comp.Equals(0, 0), true, `Runes.Equals(0, 0)`)
	boolEqual(t, comp.Equals(0, 1), false, `Runes.Equals(0, 1)`)
	boolEqual(t, comp.Equals(1, 0), false, `Runes.Equals(1, 0)`)
	boolEqual(t, comp.Equals(1, 1), false, `Runes.Equals(1, 1)`)
	intEqual(t, int(comp.AValue(1)), 'b', `Runes.AValue(1)`)
	intEqual(t, int(comp.BValue(1)), 'd', `Runes.BValue(1)`)
}

func Test_Integer(t *testing.T) {
	comp := NewInteger(
		[]int{1, 2, 3},
		[]int{1, 4})
	intEqual(t, comp.ALength(), 3, `Integer.ALength`)
	intEqual(t, comp.BLength(), 2, `Integer.BLength`)
	boolEqual(t, comp.Equals(0, 0), true, `Integer.Equals(0, 0)`)
	boolEqual(t, comp.Equals(0, 1), false, `Integer.Equals(0, 1)`)
	boolEqual(t, comp.Equals(1, 0), false, `Integer.Equals(1, 0)`)
	boolEqual(t, comp.Equals(1, 1), false, `Integer.Equals(1, 1)`)
	intEqual(t, comp.AValue(1), 2, `Integer.AValue(1)`)
	intEqual(t, comp.BValue(1), 4, `Integer.BValue(1)`)
}

func Test_RuneSlice(t *testing.T) {
	comp := NewRuneSlice(
		[][]rune{[]rune(`cat`), []rune(`cats`), []rune(`ca`)},
		[][]rune{[]rune(`cat`), []rune(`dog`)})
	intEqual(t, comp.ALength(), 3, `RuneSlice.ALength`)
	intEqual(t, comp.BLength(), 2, `RuneSlice.BLength`)
	boolEqual(t, comp.Equals(0, 0), true, `RuneSlice.Equals(0, 0)`)
	boolEqual(t, comp.Equals(0, 1), false, `RuneSlice.Equals(0, 1)`)
	boolEqual(t, comp.Equals(1, 0), false, `RuneSlice.Equals(1, 0)`)
	boolEqual(t, comp.Equals(1, 1), false, `RuneSlice.Equals(1, 1)`)
	boolEqual(t, comp.Equals(2, 0), false, `RuneSlice.Equals(1, 0)`)
	strEqual(t, string(comp.AValue(1)), `cats`, `RuneSlice.AValue(1)`)
	strEqual(t, string(comp.BValue(1)), `dog`, `RuneSlice.BValue(1)`)
}

func Test_Interface_Float(t *testing.T) {
	const epsilon = 0.001
	comp := NewInterface(
		[]interface{}{1.2345, 1.23, 3.14159},
		[]interface{}{1.235, 1.2356},
		func(a, b interface{}) bool {
			af, bf := a.(float64), b.(float64)
			if af > bf {
				return af-bf < epsilon
			}
			return bf-af < epsilon
		})
	intEqual(t, comp.ALength(), 3, `Interface(Float).ALength`)
	intEqual(t, comp.BLength(), 2, `Interface(Float).BLength`)
	boolEqual(t, comp.Equals(0, 0), true, `Interface(Float).Equals(0, 0)`)
	boolEqual(t, comp.Equals(0, 1), false, `Interface(Float).Equals(0, 1)`)
	boolEqual(t, comp.Equals(1, 0), false, `Interface(Float).Equals(1, 0)`)
	boolEqual(t, comp.Equals(1, 1), false, `Interface(Float).Equals(1, 1)`)
	boolEqual(t, comp.Equals(2, 0), false, `Interface(Float).Equals(1, 0)`)
	strEqual(t, fmt.Sprint(comp.AValue(1)), `1.23`, `Interface(Float).AValue(1)`)
	strEqual(t, fmt.Sprint(comp.BValue(1)), `1.2356`, `Interface(Float).BValue(1)`)
}

type Cat struct {
	name string
}

func Test_Interface_Cat(t *testing.T) {
	cat1 := &Cat{name: `kitty`}
	cat2 := &Cat{name: `mittens`}
	cat3 := &Cat{name: `kitty`}
	comp := NewInterface(
		[]interface{}{cat1, cat2, cat3},
		[]interface{}{cat1, cat3}, nil)
	intEqual(t, comp.ALength(), 3, `Interface(Cat).ALength`)
	intEqual(t, comp.BLength(), 2, `Interface(Cat).BLength`)
	boolEqual(t, comp.Equals(0, 0), true, `Interface(Cat).Equals(0, 0)`)
	boolEqual(t, comp.Equals(0, 1), false, `Interface(Cat).Equals(0, 1)`)
	boolEqual(t, comp.Equals(1, 0), false, `Interface(Cat).Equals(1, 0)`)
	boolEqual(t, comp.Equals(1, 1), false, `Interface(Cat).Equals(1, 1)`)
	boolEqual(t, comp.Equals(2, 0), false, `Interface(Cat).Equals(2, 0)`)
	boolEqual(t, comp.Equals(2, 1), true, `Interface(Cat).Equals(2, 1)`)
	strEqual(t, fmt.Sprintf("%+v", comp.AValue(1)), `&{name:mittens}`, `Interface(Cat).AValue(1)`)
	strEqual(t, fmt.Sprintf("%+v", comp.BValue(1)), `&{name:kitty}`, `Interface(Cat).BValue(1)`)
}
