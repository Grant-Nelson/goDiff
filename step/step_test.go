package step

import (
	"fmt"
	"testing"
)

func strEqual(t *testing.T, value, exp string) {
	if value != exp {
		t.Error(fmt.Sprint("Unexpected string value:",
			"\n   Value:    ", value,
			"\n   Expected: ", exp))
	}
}

func Test_Type(t *testing.T) {
	strEqual(t, Equal.String(), `=`)
	strEqual(t, Added.String(), `+`)
	strEqual(t, Removed.String(), `-`)
	strEqual(t, ((Type)(4)).String(), `?`)
}
