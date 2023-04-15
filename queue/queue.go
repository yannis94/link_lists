package queue

type data interface {
}

type Node struct {
    Data data
    next *Node
}

type Queue struct {
    Length int
    head *Node
    tail *Node
}

func BuildQueue() *Queue {
    return &Queue{ Length: 0, }
}

func (q *Queue) Enqueue(data data) {
    newNode := &Node{ Data: data, }

    if q.Length == 0 {
        q.head = newNode
    } else {
        q.tail.next = newNode
    }

    q.tail = newNode
    q.Length++
}

func (q *Queue) Dequeue() data {
    if q.Length == 0 {
        return nil
    }

    node := q.head
    q.head = node.next
    node.next = nil

    q.Length--

    return node.Data
}

func (q *Queue) Peek() data {
    if q.Length == 0 {
        return nil
    }

    return q.head.Data
}
