package datastructure

import (
	"errors"
)

type data interface {
}

type Node struct {
    Data data
    prev *Node
    next *Node
}

type LinkList struct {
    Length int
    head *Node
    tail *Node
}

func BuildQueue() *LinkList {
    return &LinkList{ Length: 0, }
}

func (l *LinkList) GetLength() int {
    return l.Length
}

func (l *LinkList) Append(data data) error {
    newNode := &Node{ Data: data, next: &Node{}, prev: &Node{}}

    if l.Length == 0 {
        l.head = newNode
        l.tail = newNode
        l.Length++
        return nil
    }

    lastNode := l.tail

    lastNode.next = newNode
    newNode.prev = lastNode

    l.tail = newNode
    l.Length++

    return nil
}

func (l *LinkList) Prepend(data data) error {
    newNode := &Node{ Data: data, next: &Node{}, prev: &Node{}}
    
    if l.Length == 0 {
        l.head = newNode
        l.tail = newNode
        l.Length++
        return nil
    }

    l.head.prev = newNode
    newNode.next = l.head
    l.head = newNode

    l.Length++

    return nil
}

func (l *LinkList) Get(index int) (data, error) {
    var currentNode *Node

    if l.Length < index {
        return currentNode.Data, errors.New("Index out of range.")
    }

    currentNode = l.head

    for i:=0; i<l.Length; i++ {
        if i == index {
            break
        }
        currentNode = currentNode.next
    }

    return currentNode.Data, nil
}

func (l *LinkList) InsertAt(index int, data data) error {
    if index > l.Length {
        return errors.New("Index out of range")
    } else if index == 0 {
        l.Prepend(data)
        return nil
    } else if index == l.Length {
        l.Append(data)
        return nil
    }

    newNode := &Node{ Data: data, next: &Node{}, prev: &Node{}}

    currentNode := l.head

    for i:=0; i<l.Length; i++ {
        if i == index {
            newNode.prev = currentNode.prev
            newNode.next = currentNode
            currentNode.prev.next = newNode
            currentNode.prev = newNode
            break
        }
        currentNode = currentNode.next
    }

    l.Length++
    return nil
}

func (l *LinkList) RemoveAt(index int) (data, error) {
    var deletedNode *Node

    if index > l.Length {
        return deletedNode.Data, errors.New("Index out of range.")
    }

    currentNode := l.head

    for i:=0; i<l.Length; i++ {
        if i == index {

            if i == 0 {
                l.head = currentNode.next
            } else if i == l.Length - 1 {
                l.tail = currentNode.prev
            }

            currentNode.prev.next = currentNode.next
            currentNode.next.prev = currentNode.prev
            deletedNode = currentNode
            deletedNode.prev = nil
            deletedNode.next = nil
            break
        }

        currentNode = currentNode.next
    }

    l.Length--
    return deletedNode.Data, nil
}
