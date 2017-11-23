package diff

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// printCostMatrix gets the cost matrix string.
func (cm *costMatrix) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString("[[")
	for i, costRow := range cm.costs {
		if i > 0 {
			buf.WriteString("],\n [")
		}
		for j, cost := range costRow {
			if j > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprint(cost))
		}
	}
	buf.WriteString("]]\n")
	return buf.String()
}

func TestLevenshteinDistance(t *testing.T) {
	checkLP(t, "A", "A", "0")
	checkLP(t, "A", "B", "21")
	checkLP(t, "A", "AB", "01")
	checkLP(t, "A", "BA", "10")
	checkLP(t, "AB", "A", "02")
	checkLP(t, "BA", "A", "20")
	checkLP(t, "kitten", "sitting", "210002101")
	checkLP(t, "saturday", "sunday", "022021000")
	checkLP(t, "satxrday", "sunday", "0222211000")
	checkLP(t, "ABC", "ADB", "0102")
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
	checkDiff(t, "\n",
		strings.Join([]string{
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
			`be added after it.`}, "\n"),
		strings.Join([]string{
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
			`to this document.`}, "\n"),
		strings.Join([]string{
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
			`+to this document.`}, "\n"))
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

	path := Path(&stringSliceComparable{a: aParts, b: bParts})
	result := ""
	for _, step := range path {
		result += fmt.Sprint(step)
	}

	if exp != result {
		t.Error("Levenshtein Distance returned unexpected result:",
			"\n   Input A:  ", a,
			"\n   Input B:  ", b,
			"\n   Expected: ", exp,
			"\n   Result:   ", result)
	}
}

// checkDiff gets the labelled differences
func checkDiff(t *testing.T, sep, a, b, exp string) {
	result := Strings(a, b, sep)
	if exp != result {
		t.Error("PartDiff returned unexpected result:",
			"\n   Input A:  ", a,
			"\n   Input B:  ", b,
			"\n   Expected: ", exp,
			"\n   Result:   ", result)
	}
}
