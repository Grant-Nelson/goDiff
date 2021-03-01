package godiff

import (
	"testing"
)

var (
	hirschbergPlusMinus = lines(
		`+This is an important`,
		`+notice! It should`,
		`+therefore be located at`,
		`+the beginning of this`,
		`+document!`,
		`+`,
		` This part of the`,
		` document has stayed the`,
		` same from version to`,
		` version.  It shouldn't`,
		` be shown if it doesn't`,
		` change.  Otherwise, that`,
		` would not be helping to`,
		`-compress the size of the`,
		`-changes.`,
		`-`,
		`-This paragraph contains`,
		`-text that is outdated.`,
		`-It will be deleted in the`,
		`-near future.`,
		`+compress anything.`,
		` `,
		` It is important to spell`,
		`-check this dokument. On`,
		`+check this document. On`,
		` the other hand, a`,
		` misspelled word isn't`,
		` the end of the world.`,
		` Nothing in the rest of`,
		` this paragraph needs to`,
		` be changed. Things can`,
		` be added after it.`,
		`+`,
		`+This paragraph contains`,
		`+important new additions`,
		`+to this document.`)

	// wagner is different because of differences in which
	// equal Levenstein distance paths are preferences.
	wagnerPlusMinus = lines(
		`+This is an important`,
		`+notice! It should`,
		`+therefore be located at`,
		`+the beginning of this`,
		`+document!`,
		`+`,
		` This part of the`,
		` document has stayed the`,
		` same from version to`,
		` version.  It shouldn't`,
		` be shown if it doesn't`,
		` change.  Otherwise, that`,
		` would not be helping to`,
		`-compress the size of the`,
		`-changes.`,
		`+compress anything.`,
		` `,
		`-This paragraph contains`,
		`-text that is outdated.`,
		`-It will be deleted in the`,
		`-near future.`,
		`-`,
		` It is important to spell`,
		`-check this dokument. On`,
		`+check this document. On`,
		` the other hand, a`,
		` misspelled word isn't`,
		` the end of the world.`,
		` Nothing in the rest of`,
		` this paragraph needs to`,
		` be changed. Things can`,
		` be added after it.`,
		`+`,
		`+This paragraph contains`,
		`+important new additions`,
		`+to this document.`)
)

func Test_PlusMinus_Lines(t *testing.T) {
	checkSlices(t, PlusMinus(exampleA, exampleB), hirschbergPlusMinus)
	checkSlices(t, PlusMinusCustom(DefaultDiff(), exampleA, exampleB), hirschbergPlusMinus)

	checkSlices(t, PlusMinusCustom(HirschbergDiff(-1, false), exampleA, exampleB), hirschbergPlusMinus)
	checkSlices(t, PlusMinusCustom(HirschbergDiff(-1, true), exampleA, exampleB), hirschbergPlusMinus)

	checkSlices(t, PlusMinusCustom(HybridDiff(-1, false, -1), exampleA, exampleB), hirschbergPlusMinus)
	checkSlices(t, PlusMinusCustom(HybridDiff(-1, true, -1), exampleA, exampleB), hirschbergPlusMinus)

	checkSlices(t, PlusMinusCustom(WagnerDiff(-1), exampleA, exampleB), wagnerPlusMinus)
}
