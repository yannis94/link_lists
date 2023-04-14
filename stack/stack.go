package stack

type Node struct {
	Value interface{}
	prev  *Node
}

type Stack struct {
	Length int
	head   *Node
}

func (s *Stack) Peek() interface{} {
	return s.head.Value
}

func (s *Stack) Push(item interface{}) {
	var node Node = Node{Value: item}
	if s.Length != 0 {
		node.prev = s.head
	}
	s.head = &node
	s.Length += 1
}

func (s *Stack) Pop() interface{} {
	if s.Length == 0 {
		return nil
	}

	poped := *s.head
	s.head = poped.prev
	s.Length -= 1

	return poped.Value
}
