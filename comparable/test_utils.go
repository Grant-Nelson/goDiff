// +build !release,!prod

package comparable

import "fmt"

// APart gets the A string for debugging the value at the comparable index.
// This should only be used in unit-tests.
func APart(comp Comparable, aIndex int) string {
	switch comp := comp.(type) {
	case *Char:
		return string(comp.AValue(aIndex))
	case *Integer:
		return fmt.Sprintf(`%d`, comp.AValue(aIndex))
	case *Interface:
		return fmt.Sprintf(`%v`, comp.AValue(aIndex))
	case *Runes:
		return string([]rune{comp.AValue(aIndex)})
	case *RuneSlice:
		return string(comp.AValue(aIndex))
	case *String:
		return comp.AValue(aIndex)
	default:
		return `?`
	}
}

// BPart gets the B string for debugging the value at the comparable index.
// This should only be used in unit-tests.
func BPart(comp Comparable, bIndex int) string {
	switch comp := comp.(type) {
	case *Char:
		return string(comp.BValue(bIndex))
	case *Integer:
		return fmt.Sprintf(`%d`, comp.BValue(bIndex))
	case *Interface:
		return fmt.Sprintf(`%v`, comp.BValue(bIndex))
	case *Runes:
		return string([]rune{comp.BValue(bIndex)})
	case *RuneSlice:
		return string(comp.BValue(bIndex))
	case *String:
		return comp.BValue(bIndex)
	default:
		return `?`
	}
}
