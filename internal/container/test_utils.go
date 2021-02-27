// +build !release,!prod

package container

import (
	"fmt"

	"github.com/Grant-Nelson/goDiff/comparable"
)

// AParts gets the strings from the comparable which are
// represented by this comparable.
// This only works for String or RuneSlice comparables.
func (cont *Container) AParts() []string {
	var fetchValue func(int) string
	switch comp := cont.comp.(type) {
	case *comparable.String:
		fetchValue = comp.AValue
	case *comparable.RuneSlice:
		fetchValue = func(aIndex int) string {
			return string(comp.AValue(aIndex))
		}
	default:
		return []string{`Unexpected Comparable Type`}
	}

	parts := make([]string, cont.aLength)
	for i := 0; i < cont.aLength; i++ {
		parts[i] = fetchValue(cont.aAdjust(i))
	}
	return parts
}

// BParts gets the strings from the comparable which are
// represented by this comparable.
// This only works for String or RuneSlice comparables.
func (cont *Container) BParts() []string {
	var fetchValue func(int) string
	switch comp := cont.comp.(type) {
	case *comparable.String:
		fetchValue = comp.BValue
	case *comparable.RuneSlice:
		fetchValue = func(bIndex int) string {
			return string(comp.BValue(bIndex))
		}
	default:
		return []string{`Unexpected Comparable Type`}
	}

	parts := make([]string, cont.bLength)
	for j := 0; j < cont.bLength; j++ {
		parts[j] = fetchValue(cont.bAdjust(j))
	}
	return parts
}

// String gets the string for debugging a container.
func (cont *Container) String() string {
	return fmt.Sprint(
		cont.aOffset, ", ", cont.aLength, ", ",
		cont.bOffset, ", ", cont.bLength, ", ",
		cont.reverse)
}
