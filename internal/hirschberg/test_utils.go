// +build !release

package hirschberg

// UseReduce indicates if the equal padding edges should be checked
// at each step of the algorithm or not.
//
// This is exposed in non-release so that the reduction add-on
// to the Hirschberg algorithm can be tested and benchmarks.
func (h *hirschberg) UseReduce(useReduce bool) {
	h.useReduce = useReduce
}
