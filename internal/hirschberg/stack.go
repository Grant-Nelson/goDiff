package hirschberg

import "github.com/Grant-Nelson/goDiff/internal/container"

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

		// graveyard is used to keep from allocating new nodes
		// when some have been already allocated and discarded.
		graveyard *stackNode
	}
)

// NewStack creates a new stack of containers.
func NewStack() *Stack {
	return &Stack{
		top:       nil,
		graveyard: nil,
	}
}

// NotEmpty indicates if this stack is not empty or is empty.
func (s *Stack) NotEmpty() bool {
	return s.top != nil
}

// Push a new container onto this stack.
func (s *Stack) Push(cont *container.Container, remainder int) {
	if (cont != nil) || (remainder > 0) {
		if s.graveyard != nil {
			node := s.graveyard
			s.graveyard = node.prev
			node.cont = cont
			node.remainder = remainder
			node.prev = s.top
			s.top = node
		} else {
			s.top = &stackNode{
				cont:      cont,
				remainder: remainder,
				prev:      s.top,
			}
		}
	}
}

// Pop a container off the stack, will be nil if empty.
func (s *Stack) Pop() (*container.Container, int) {
	node := s.top
	if node != nil {
		s.top = node.prev
		node.prev = s.graveyard
		s.graveyard = node
		cont := node.cont
		remainder := node.remainder
		node.cont = nil
		node.remainder = 0
		return cont, remainder
	}
	return nil, 0
}
