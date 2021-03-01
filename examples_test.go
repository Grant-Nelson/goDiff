package godiff

import (
	"fmt"
	"strings"

	"github.com/Grant-Nelson/goDiff/comparable"
	"github.com/Grant-Nelson/goDiff/step"
)

func ExamplePlusMinus() {
	original := "Shopping List:\n" +
		"Eggs\n" +
		"Bacon\n" +
		"Apples\n" +
		"Oranges\n" +
		"Milk"
	changed := "Shopping List:\n" +
		"Yogurt\n" +
		"Apples\n" +
		"Oranges\n" +
		"Bananas\n" +
		"Milk\n" +
		"Tea"

	originalLines := strings.Split(original, "\n")
	changedLines := strings.Split(changed, "\n")
	diffLines := PlusMinus(originalLines, changedLines)
	diff := strings.Join(diffLines, "\n")

	fmt.Println(diff)
	// Output: Shopping List:
	// -Eggs
	// -Bacon
	// +Yogurt
	//  Apples
	//  Oranges
	// +Bananas
	//  Milk
	// +Tea
}

func ExampleMerge() {
	original := "Shopping List:\n" +
		"Eggs\n" +
		"Bacon\n" +
		"Apples\n" +
		"Oranges\n" +
		"Milk"
	changed := "Shopping List:\n" +
		"Yogurt\n" +
		"Apples\n" +
		"Oranges\n" +
		"Bananas\n" +
		"Milk\n" +
		"Tea"

	originalLines := strings.Split(original, "\n")
	changedLines := strings.Split(changed, "\n")
	diffLines := Merge(originalLines, changedLines)
	diff := strings.Join(diffLines, "\n")

	fmt.Println(diff)
	// Output: Shopping List:
	// <<<<<<<<
	// Eggs
	// Bacon
	// ========
	// Yogurt
	// >>>>>>>>
	// Apples
	// Oranges
	// <<<<<<<<
	// ========
	// Bananas
	// >>>>>>>>
	// Milk
	// <<<<<<<<
	// ========
	// Tea
	// >>>>>>>>
}

func ExampleDiff() {
	original := "Shopping List:\n" +
		"Eggs\n" +
		"Bacon\n" +
		"Apples\n" +
		"Oranges\n" +
		"Milk"
	changed := "Shopping List:\n" +
		"Yogurt\n" +
		"Apples\n" +
		"Oranges\n" +
		"Bananas\n" +
		"Milk\n" +
		"Tea"

	originalLines := strings.Split(original, "\n")
	changedLines := strings.Split(changed, "\n")
	result := Diff(comparable.NewString(originalLines, changedLines))

	fmt.Println("Count:", result.Count())
	fmt.Println("Total:", result.Total())
	result.Read(func(step step.Type, count int) {
		fmt.Println(step, count)
	})

	// Output: Count: 7
	// Total: 9
	// = 1
	// - 2
	// + 1
	// = 2
	// + 1
	// = 1
	// + 1
}
