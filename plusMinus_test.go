package godiff

import "testing"

func Test_PlusMinus_Lines(t *testing.T) {
	checkSlices(t, PlusMinus(exampleA, exampleB), lines(
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
		`+to this document.`))
}
