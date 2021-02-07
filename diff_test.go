package diff

import (
	"fmt"
	"strings"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	checkLP(t, "A", "A", "0x1")
	checkLP(t, "A", "B", "2x1, 1x1")
	checkLP(t, "A", "AB", "0x1, 1x1")
	checkLP(t, "A", "BA", "1x1, 0x1")
	checkLP(t, "AB", "A", "0x1, 2x1")
	checkLP(t, "BA", "A", "2x1, 0x1")
	checkLP(t, "kitten", "sitting", "2x1, 1x1, 0x3, 2x1, 1x1, 0x1, 1x1")
	checkLP(t, "saturday", "sunday", "0x1, 2x2, 0x1, 2x1, 1x1, 0x3")
	checkLP(t, "satxrday", "sunday", "0x1, 2x4, 1x2, 0x3")
	checkLP(t, "ABC", "ADB", "0x1, 1x1, 0x1, 2x1")
}

func TestPartDiff(t *testing.T) {
	checkDiff(t, ",",
		"cat,dog,pig",
		"cat,horse,dog",
		" cat,+horse, dog,-pig")
	checkDiff(t, ",",
		"mike,ted,mark,jim",
		"ted,mark,bob,bill",
		"-mike, ted, mark,-jim,+bob,+bill")
	checkDiff(t, ",",
		"k,i,t,t,e,n",
		"s,i,t,t,i,n,g",
		"-k,+s, i, t, t,-e,+i, n,+g")
	checkDiff(t, ",",
		"s,a,t,u,r,d,a,y",
		"s,u,n,d,a,y",
		" s,-a,-t, u,-r,+n, d, a, y")
	checkDiff(t, ",",
		"s,a,t,x,r,d,a,y",
		"s,u,n,d,a,y",
		" s,-a,-t,-x,-r,+u,+n, d, a, y")
	checkDiff(t, ",",
		"func A() int,{,return 10,},,func C() int,{,return 12,}",
		"func A() int,{,return 10,},,func B() int,{,return 11,},,func C() int,{,return 12,}",
		" func A() int, {, return 10, }, ,+func B() int,+{,+return 11,+},+, func C() int, {, return 12, }")
}

func TestFormatting(t *testing.T) {
	checkAll(t,
		lines(
			`This part of the`,
			`document has stayed the`,
			`same from version to`,
			`version.  It shouldn't`,
			`be shown if it doesn't`,
			`change.  Otherwise, that`,
			`would not be helping to`,
			`compress the size of the`,
			`changes.`,
			``,
			`This paragraph contains`,
			`text that is outdated.`,
			`It will be deleted in the`,
			`near future.`,
			``,
			`It is important to spell`,
			`check this dokument. On`,
			`the other hand, a`,
			`misspelled word isn't`,
			`the end of the world.`,
			`Nothing in the rest of`,
			`this paragraph needs to`,
			`be changed. Things can`,
			`be added after it.`),
		lines(
			`This is an important`,
			`notice! It should`,
			`therefore be located at`,
			`the beginning of this`,
			`document!`,
			``,
			`This part of the`,
			`document has stayed the`,
			`same from version to`,
			`version.  It shouldn't`,
			`be shown if it doesn't`,
			`change.  Otherwise, that`,
			`would not be helping to`,
			`compress anything.`,
			``,
			`It is important to spell`,
			`check this document. On`,
			`the other hand, a`,
			`misspelled word isn't`,
			`the end of the world.`,
			`Nothing in the rest of`,
			`this paragraph needs to`,
			`be changed. Things can`,
			`be added after it.`,
			``,
			`This paragraph contains`,
			`important new additions`,
			`to this document.`),
		lines(
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
			`+to this document.`),
		lines(
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
			`========`,
			`compress anything.`,
			`>>>>>>>>`,
			``,
			`<<<<<<<<`,
			`This paragraph contains`,
			`text that is outdated.`,
			`It will be deleted in the`,
			`near future.`,
			``,
			`========`,
			`>>>>>>>>`,
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

func lines(ln ...string) []string {
	return ln
}

// checks the levenshtein distance algorithm
func checkLP(t *testing.T, a, b, exp string) {
	aParts := []string{}
	for _, part := range a {
		aParts = append(aParts, string([]rune{part}))
	}

	bParts := []string{}
	for _, part := range b {
		bParts = append(bParts, string([]rune{part}))
	}

	path := Path(newStrSliceComp(aParts, bParts))
	parts := []string{}
	for _, step := range path {
		parts = append(parts, fmt.Sprint(step.Type, `x`, step.Count))
	}
	result := strings.Join(parts, `, `)

	if exp != result {
		t.Error("Levenshtein Distance returned unexpected result:",
			"\n   Input A:  ", a,
			"\n   Input B:  ", b,
			"\n   Expected: ", exp,
			"\n   Result:   ", result)
	}
}

// checkDiff gets the labelled differences for PlusMinus
func checkDiff(t *testing.T, sep, a, b, exp string) {
	aParts := strings.Split(a, sep)
	bParts := strings.Split(b, sep)
	resultParts := PlusMinus(aParts, bParts)
	result := strings.Join(resultParts, sep)
	if exp != result {
		t.Error("PartDiff returned unexpected result:",
			"\n   Input A:  ", a,
			"\n   Input B:  ", b,
			"\n   Expected: ", exp,
			"\n   Result:   ", result)
	}
}

// checkAll gets the labelled differences for different formatting.
func checkAll(t *testing.T, a, b, expPlusMinus, expMerge []string) {
	checkSlices(t, PlusMinus(a, b), expPlusMinus)
	checkSlices(t, Merge(a, b), expMerge)
}

// checkSlices checks that the given slices are the same.
func checkSlices(t *testing.T, result, exp []string) {
	resultStr := strings.Join(result, "\n")
	expStr := strings.Join(exp, "\n")
	if expStr != resultStr {
		t.Error("Unexpected result:",
			"\n   Expected: ", expStr,
			"\n   Result:   ", resultStr)
	}
}
