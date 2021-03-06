package collector

import (
	"fmt"
	"testing"

	"github.com/Grant-Nelson/goDiff/step"
)

func Test_Basics(t *testing.T) {
	col := New()

	col.InsertAdded(1)
	col.InsertRemoved(1)
	col.InsertAdded(2)
	col.InsertRemoved(2)
	col.InsertEqual(3)

	col.InsertAdded(4)
	col.InsertEqual(2)
	col.InsertEqual(2)

	col.InsertRemoved(5)
	col.InsertEqual(2)
	col.InsertEqual(3)

	col.InsertRemoved(-6)
	col.InsertEqual(-6)
	col.InsertAdded(-6)

	readEqual(t, col, `not finished`)
	boolEqual(t, col.Finished(), false, `Collection.Finished`)
	col.Finish()
	boolEqual(t, col.Finished(), true, `Collection.Finished`)

	intEqual(t, col.Count(), 7, `Collection.Count`)
	intEqual(t, col.Total(), 27, `Collection.Total`)
	readEqual(t, col, `=5 -5 =4 +4 =3 -3 +3`)
}

func Test_Error(t *testing.T) {
	col := New()

	col.InsertAdded(1)
	col.InsertRemoved(1)
	col.InsertEqual(3)
	col.InsertRemoved(2)
	col.InsertAdded(2)
	col.InsertSubstitute(3)

	panicEqual(t, func() { col.Count() }, errFinishBeforeCount, `Collection.Count`)
	panicEqual(t, func() { col.Total() }, errFinishBeforeTotal, `Collection.Total`)
	panicEqual(t, func() { col.Read(nil) }, errFinishBeforeRead, `Collection.Read`)

	col.Finish()

	intEqual(t, col.Count(), 5, `Collection.Count`)
	intEqual(t, col.Total(), 15, `Collection.Total`)
	readEqual(t, col, `-5 +5 =3 -1 +1`)

	panicEqual(t, func() { col.Finish() }, errFinishAfterFinish, `Collection.Finish`)
	panicEqual(t, func() { col.InsertAdded(4) }, errInsertAfterFinish, `Collection.InsertAdded`)
	panicEqual(t, func() { col.InsertRemoved(4) }, errInsertAfterFinish, `Collection.InsertRemoved`)
	panicEqual(t, func() { col.InsertEqual(4) }, errInsertAfterFinish, `Collection.InsertEqual`)
	panicEqual(t, func() { col.InsertSubstitute(4) }, errInsertAfterFinish, `Collection.InsertSubstitute`)
}

func Test_ForcePush(t *testing.T) {
	col := New()

	col.ForcePush(step.Added, 1)
	col.ForcePush(step.Added, 2)
	col.ForcePush(step.Removed, 3)
	col.ForcePush(step.Removed, 4)
	col.ForcePush(step.Equal, 5)
	col.ForcePush(step.Equal, 6)
	col.Finish()

	readEqual(t, col, `=6 =5 -4 -3 +2 +1`)
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

func readEqual(t *testing.T, col *Collector, exp string) {
	if value := col.String(); value != exp {
		t.Error(fmt.Sprint("Unexpected collection read:",
			"\n   Value:    ", value,
			"\n   Expected: ", exp))
	}
}

func panicEqual(t *testing.T, hndl func(), exp, msg string) {
	value := func() (errMsg string) {
		defer func() {
			if r := recover(); r != nil {
				errMsg = fmt.Sprint(r)
			}
		}()

		hndl()

		return `no panic occurred`
	}()
	if value != exp {
		t.Error(fmt.Sprint("Unexpected panic message:",
			"\n   Message:  ", msg,
			"\n   Value:    ", value,
			"\n   Expected: ", exp))
	}
}
