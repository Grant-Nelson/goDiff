// +build !release,!prod

package collector

import (
	"fmt"
	"strings"

	"github.com/Grant-Nelson/goDiff/step"
)

// String gets the debug string of this collector if it has been finished.
func (c *Collector) String() string {
	if !c.Finished() {
		return `not finished`
	}

	parts := make([]string, 0, c.Count())
	c.Read(func(step step.Type, count int) {
		parts = append(parts, fmt.Sprintf(`%s%d`, step.String(), count))
	})
	return strings.Join(parts, ` `)
}
