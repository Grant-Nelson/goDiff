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

// ForcePush pushes a new step into the collection
// without otherwise effecting the collector.
// This is used to test very specific edge cases which may or may not
// occur, but should be protected against, when processing results of a diff.
func (c *Collector) ForcePush(step step.Type, count int) {
	c.push(step, count)
}
