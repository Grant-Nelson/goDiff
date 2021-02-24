package hirschberg

import "../container"

type (
	// stackNode is a node in a stack of containers.
	stackNode struct {

		// cont is the container for this node.
		cont *container.Container

		// remainder is any equals left over from the reduction.
		remainder int

		// prev is the node previous to this node in the stack.
		prev *stackNode
	}

	// Stack is a stack of containers.
	Stack struct {

		// top is the top of the stack.
		top *stackNode
	}
)

// NewStack creates a new stack of containers.
func NewStack() *Stack {
	return &Stack{
		top: nil,
	}
}

// NotEmpty indicates if this stack is not empty or is empty.
func (s *Stack) NotEmpty() bool {
	return s.top != nil
}

// Push a new container onto this stack.
func (s *Stack) Push(cont *container.Container, remainder int) {
	s.top = &stackNode{
		cont:      cont,
		remainder: remainder,
		prev:      s.top,
	}
}

// Pop a container off the stack, will be nil if empty.
func (s *Stack) Pop() (*container.Container, int) {
	node := s.top
	if node != nil {
		s.top = node.prev
		return node.cont, node.remainder
	}
	return nil, 0
}
