package diff

import (
	"fmt"
	"strings"
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

func ExamplePath() {

}
