package godiff

import (
	"testing"

	"github.com/Grant-Nelson/goDiff/comparable"
	"github.com/Grant-Nelson/goDiff/internal/collector"
	"github.com/Grant-Nelson/goDiff/step"
)

func Test_Merge_Lines(t *testing.T) {
	checkSlices(t, Merge(exampleA, exampleB), lines(
		`<<<<<<<<`,
		`========`,
		`This is an important`,
		`notice! It should`,
		`therefore be located at`,
		`the beginning of this`,
		`document!`,
		``,
		`>>>>>>>>`,
		`This part of the`,
		`document has stayed the`,
		`same from version to`,
		`version.  It shouldn't`,
		`be shown if it doesn't`,
		`change.  Otherwise, that`,
		`would not be helping to`,
		`<<<<<<<<`,
		`compress the size of the`,
		`changes.`,
		``,
		`This paragraph contains`,
		`text that is outdated.`,
		`It will be deleted in the`,
		`near future.`,
		`========`,
		`compress anything.`,
		`>>>>>>>>`,
		``,
		`It is important to spell`,
		`<<<<<<<<`,
		`check this dokument. On`,
		`========`,
		`check this document. On`,
		`>>>>>>>>`,
		`the other hand, a`,
		`misspelled word isn't`,
		`the end of the world.`,
		`Nothing in the rest of`,
		`this paragraph needs to`,
		`be changed. Things can`,
		`be added after it.`,
		`<<<<<<<<`,
		`========`,
		``,
		`This paragraph contains`,
		`important new additions`,
		`to this document.`,
		`>>>>>>>>`))
}

func Test_Merge_MoreCases(t *testing.T) {
	checkSlices(t, Merge(lines(
		`sameA`,
		`removedA`,
		`sameB`,
		`sameC`,
		`removedC`,
		`sameD`,
	), lines(
		`sameA`,
		`sameB`,
		`AddedB`,
		`sameC`,
		`AddedC`,
		`sameD`,
	)), lines(
		`sameA`,
		`<<<<<<<<`,
		`removedA`,
		`========`,
		`>>>>>>>>`,
		`sameB`,
		`<<<<<<<<`,
		`========`,
		`AddedB`,
		`>>>>>>>>`,
		`sameC`,
		`<<<<<<<<`,
		`removedC`,
		`========`,
		`AddedC`,
		`>>>>>>>>`,
		`sameD`,
	))

	checkSlices(t, Merge(lines(
		`sameA`,
		`removedA`,
	), lines(
		`sameA`,
	)), lines(
		`sameA`,
		`<<<<<<<<`,
		`removedA`,
		`========`,
		`>>>>>>>>`,
	))

	checkSlices(t, Merge(lines(
		`sameA`,
	), lines(
		`sameA`,
		`addedA`,
	)), lines(
		`sameA`,
		`<<<<<<<<`,
		`========`,
		`addedA`,
		`>>>>>>>>`,
	))
}

func Test_Merge_EdgeCases(t *testing.T) {
	checkSlices(t, Merge(lines(
		`sameA`,
		`removedA`,
	), lines(
		`sameA`,
		`AddedA`,
	)), lines(
		`sameA`,
		`<<<<<<<<`,
		`removedA`,
		`========`,
		`AddedA`,
		`>>>>>>>>`,
	))

	// Normally remove is first but check if added is first.
	checkSlices(t, MergeCustom(func(comp comparable.Comparable) Results {
		col := collector.New()
		col.ForcePush(step.Removed, 1)
		col.ForcePush(step.Added, 1)
		col.ForcePush(step.Equal, 1)
		col.Finish()
		return col
	}, lines(
		`sameA`,
		`removedA`,
	), lines(
		`sameA`,
		`AddedA`,
	)), lines(
		`sameA`,
		`<<<<<<<<`,
		`========`,
		`AddedA`,
		`>>>>>>>>`,
		`<<<<<<<<`,
		`removedA`,
		`========`,
		`>>>>>>>>`,
	))
}
