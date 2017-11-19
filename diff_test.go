package diff

import (
	"fmt"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	checkLP(t, "A", "A", "0")
	checkLP(t, "A", "B", "1")
	checkLP(t, "A", "AB", "02")
	checkLP(t, "A", "BA", "20")
	checkLP(t, "AB", "A", "03")
	checkLP(t, "BA", "A", "30")
	checkLP(t, "kitten", "sitting", "1000102")
	checkLP(t, "saturday", "sunday", "03301000")
	checkLP(t, "satxrday", "sunday", "01133000")
	checkLP(t, "ABC", "ADB", "0203")
}

func TestPartDiff(t *testing.T) {
	checkDiff(t,
		"cat,dog,pig",
		"cat,horse,dog",
		" cat,+horse, dog,-pig")
	checkDiff(t,
		"mike,ted,mark,jim",
		"ted,mark,bob,bill",
		"-mike, ted, mark,-jim,+bob,+bill")
	checkDiff(t,
		"k,i,t,t,e,n",
		"s,i,t,t,i,n,g",
		"-k,+s, i, t, t,-e,+i, n,+g")
	checkDiff(t,
		"s,a,t,u,r,d,a,y",
		"s,u,n,d,a,y",
		" s,-a,-t, u,-r,+n, d, a, y")
	checkDiff(t,
		"s,a,t,x,r,d,a,y",
		"s,u,n,d,a,y",
		" s,-a,-t,-x,-r,+u,+n, d, a, y")
	checkDiff(t,
		"func A() int,{,return 10,},,func C() int,{,return 12,}",
		"func A() int,{,return 10,},,func B() int,{,return 11,},,func C() int,{,return 12,}",
		" func A() int, {, return 10, }, ,+func B() int,+{,+return 11,+},+, func C() int, {, return 12, }")
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

	_, path := levenshteinDistance(aParts, bParts, len(aParts), len(bParts))
	result := ""
	for _, step := range path {
		result += fmt.Sprint(step)
	}

	if exp != result {
		t.Fatal("Levenshtein Distance returned unexpected result:",
			"\n   Input A:  ", a,
			"\n   Input B:  ", b,
			"\n   Expected: ", exp,
			"\n   Result:   ", result)
	}
}

// checkDiff gets the labelled differences
func checkDiff(t *testing.T, a, b, exp string) {
	result := PartDiff(a, b, ",")
	if exp != result {
		t.Fatal("PartDiff returned unexpected result:",
			"\n   Input A:  ", a,
			"\n   Input B:  ", b,
			"\n   Expected: ", exp,
			"\n   Result:   ", result)
	}
}
