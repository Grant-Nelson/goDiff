package godiff

import (
	"fmt"
	"testing"

	"github.com/Grant-Nelson/goDiff/comparable"
)

const (
	billNyeA = `The most serious problem facing humankind is climate change. ` +
		`All of these people breathing and burning our atmosphere has led to an ` +
		`extraordinarily dangerous situation. I hope next generation will emerge ` +
		`and produce technology, regulations, and a worldview that enable as many ` +
		`of us as possible to live happy healthy lives.`

	billNyeB = `The meaning of life is pretty clear: Living things strive to ` +
		`pass their genes into the future. The claim that we would not have ` +
		`morals or ethics without religion is extraordinary. Animals in nature ` +
		`seem to behave in moral ways without organized religion.`
)

func Benchmark_DiffComparison(b *testing.B) {
	const groups = 4
	for i := 0; i < groups; i++ {
		inputA := billNyeA[:len(billNyeA)*(i+1)/groups]
		for j := 0; j < groups; j++ {
			inputB := billNyeB[:len(billNyeB)*(j+1)/groups]
			comp := comparable.NewChar(inputA, inputB)

			b.Run(fmt.Sprintf(`Hirschberg-NoReduce-%dx%d`, len(inputA), len(inputB)),
				func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						hirschbergDiff(-1, false)(comp)
					}
				})

			b.Run(fmt.Sprintf(`Hirschberg-UseReduce-%dx%d`, len(inputA), len(inputB)),
				func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						hirschbergDiff(-1, true)(comp)
					}
				})

			b.Run(fmt.Sprintf(`Wagner-%dx%d`, len(inputA), len(inputB)),
				func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						wagnerDiff(-1)(comp)
					}
				})

			b.Run(fmt.Sprintf(`Hybrid-NoReduce-100-%dx%d`, len(inputA), len(inputB)),
				func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						hybridDiff(-1, false, 100)(comp)
					}
				})

			b.Run(fmt.Sprintf(`Hybrid-UesReduce-100-%dx%d`, len(inputA), len(inputB)),
				func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						hybridDiff(-1, true, 100)(comp)
					}
				})

			b.Run(fmt.Sprintf(`Hybrid-NoReduce-500-%dx%d`, len(inputA), len(inputB)),
				func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						hybridDiff(-1, false, 500)(comp)
					}
				})

			b.Run(fmt.Sprintf(`Hybrid-UseReduce-500-%dx%d`, len(inputA), len(inputB)),
				func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						hybridDiff(-1, true, 500)(comp)
					}
				})
		}
	}
}
