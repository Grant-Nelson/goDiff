// +build !release,!prod

package container

import (
	"fmt"

	"github.com/Grant-Nelson/goDiff/comparable"
)

// AAdjust gets the A index adjusted by the container's condition.
func (cont *Container) AAdjust(aIndex int) int {
	if cont.reverse {
		return cont.aLength - 1 - aIndex + cont.aOffset
	}
	return aIndex + cont.aOffset
}

// BAdjust gets the B index adjusted by the container's condition.
func (cont *Container) BAdjust(bIndex int) int {
	if cont.reverse {
		return cont.bLength - 1 - bIndex + cont.bOffset
	}
	return bIndex + cont.bOffset
}

// AParts gets the strings from the comparable which are
// represented by this comparable.
// This only works for String or RuneSlice comparables.
func (cont *Container) AParts() []string {
	parts := make([]string, cont.aLength)
	for i := 0; i < cont.aLength; i++ {
		parts[i] = comparable.APart(cont.comp, cont.AAdjust(i))
	}
	return parts
}

// BParts gets the strings from the comparable which are
// represented by this comparable.
// This only works for String or RuneSlice comparables.
func (cont *Container) BParts() []string {
	parts := make([]string, cont.bLength)
	for j := 0; j < cont.bLength; j++ {
		parts[j] = comparable.BPart(cont.comp, cont.BAdjust(j))
	}
	return parts
}

// String gets the string for debugging a container.
func (cont *Container) String() string {
	return fmt.Sprintf(`%d, %d, %d, %d, %t`,
		cont.aOffset, cont.aLength, cont.bOffset, cont.bLength, cont.reverse)
}
