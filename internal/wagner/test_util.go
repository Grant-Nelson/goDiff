// +build !release,!prod

package wagner

import (
	"fmt"
	"strings"
)

func (w *wagner) maxStrWidth(i, aLen, bLen int) int {
	maxWidth := 0
	for j := 0; j < bLen; j++ {
		width := len(fmt.Sprint(w.costs[i+j*aLen]))
		if width > maxWidth {
			maxWidth = width
		}
	}
	return maxWidth
}

func (w *wagner) costsString(aLen, bLen int) string {
	parts := make([]string, bLen)
	for i := 0; i < aLen; i++ {
		maxWidth := w.maxStrWidth(i, aLen, bLen)
		for j := 0; j < bLen; j++ {
			parts[j] = fmt.Sprintf(`%*d`, maxWidth, w.costs[i+j*aLen])
		}
	}
	for i := 1; i < aLen; i++ {
		maxWidth := w.maxStrWidth(i, aLen, bLen)
		for j := 0; j < bLen; j++ {
			parts[j] = fmt.Sprintf(`%s %*d`, parts[j], maxWidth, w.costs[i+j*aLen])
		}
	}
	return `[` + strings.Join(parts, "\n ") + `]`
}
