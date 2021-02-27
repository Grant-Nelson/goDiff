// +build !release,!prod

package wagner

import (
	"fmt"
	"strings"
)

func (w *wagner) costsString(aLen, bLen int) string {
	parts := make([]string, bLen)
	for i := 0; i < aLen; i++ {
		maxWidth := 0
		for j := 0; j < bLen; j++ {
			width := len(fmt.Sprint(w.costs[i+j*aLen]))
			if width > maxWidth {
				maxWidth = width
			}
		}
		for j := 0; j < bLen; j++ {
			parts[j] = fmt.Sprintf(`%s %*d`, parts[j], maxWidth, w.costs[i+j*aLen])
		}
	}
	return `[` + strings.Join(parts, "\n ") + `]`
}
