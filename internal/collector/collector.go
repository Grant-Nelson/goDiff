package collector

type (
	// StepType is the steps of the levenshtein path.
	StepType int

	// PathCallback is the function signature for calling back steps in the path.
	PathCallback func(step StepType, count int)

	// stepNode is a continuous group of step types stored in the collector.
	stepNode struct {

		// Step is the type for this group.
		Step StepType

		// Count is the number of the given type in the group.
		Count int

		// Next is the step which occurs after this step, or nil when done.
		Next *stepNode
	}

	// Collector collects, groups, and reverses the steps taken
	// by a diff algorithm in reverse order.
	Collector struct {

		// head is the most recent node added to the collection.
		// Will be nil if there are no results yet.
		head *stepNode

		// count is the number of steps that have been collected.
		count int

		// total is the total number of parts represented by this collection.
		// The total sum of all the counts in each step.
		total int

		// addedRun is the current amount of consecutive Added parts.
		addedRun int

		// removeRun is the current amount of consecutive Removed parts.
		removedRun int

		// equalRun is the current amount of consecutive Equal parts.
		equalRun int
	}
)

const (
	// Equal indicates A and B entries are equal.
	Equal StepType = iota

	// Added indicates A was added.
	Added

	// Removed indicates A was removed.
	Removed
)

// New creates a new collector.
func New() *Collector {
	return &Collector{
		head:       nil,
		count:      0,
		total:      0,
		addedRun:   0,
		removedRun: 0,
		equalRun:   0,
	}
}

// push pushes a new step into the collection.
func (c *Collector) push(step StepType, count int) {
	c.head = &stepNode{
		Step:  step,
		Count: count,
		Next:  c.head,
	}
	c.count++
	c.total += count
}

// pushAdd pushes an Added step if there is any Added parts currently collected.
func (c *Collector) pushAdded() {
	if c.addedRun > 0 {
		c.push(Added, c.addedRun)
		c.addedRun = 0
	}
}

// pushRemove pushes an Removed step if there is any Removed parts currently collected.
func (c *Collector) pushRemoved() {
	if c.removedRun > 0 {
		c.push(Removed, c.removedRun)
		c.removedRun = 0
	}
}

// pushEqual pushes an Add step if there is any Add parts currently collected.
func (c *Collector) pushEqual() {
	if c.equalRun > 0 {
		c.push(Equal, c.equalRun)
		c.equalRun = 0
	}
}

// InsertAdded inserts new Added parts into this collection.
// This is expected to be inserted in reverse order from the expected result.
func (c *Collector) InsertAdded(count int) {
	if count > 0 {
		c.pushEqual()
		c.addedRun += count
	}
}

// InsertRemoved inserts new Removed parts into this collection.
// This is expected to be inserted in reverse order from the expected result.
func (c *Collector) InsertRemoved(count int) {
	if count > 0 {
		c.pushEqual()
		c.removedRun += count
	}
}

// InsertEqual inserts new Equal parts into this collection.
// This is expected to be inserted in reverse order from the expected result.
func (c *Collector) InsertEqual(count int) {
	if count > 0 {
		c.pushAdded()
		c.pushRemoved()
		c.equalRun += count
	}
}

// Finish inserts any remaining parts which haven't been inserted yet.
func (c *Collector) Finish() {
	c.pushAdded()
	c.pushRemoved()
	c.pushEqual()
}

// Count is the number of steps that have been collected.
// Finish should be called prior to this being used.
func (c *Collector) Count() int {
	return c.count
}

// Total is the total number of parts represented by this collection.
// The total sum of all the counts in each step.
// Finish should be called prior to this being used.
func (c *Collector) Total() int {
	return c.total
}

// Read will read the collected steps in the expected result order,
// reversed from the order that it was inserted.
func (c *Collector) Read(hndl PathCallback) {
	node := c.head
	for node != nil {
		hndl(node.Step, node.Count)
		node = node.Next
	}
}