package stack

import "testing"

type book struct {
    author string
    title string
    pages int
}

func  buildStack() Stack {
    stack := Stack{ Length: 0,}

    //fake data, my stack will contain books
    b1 := book{author: "Shakespare", title: "Romeo and Juliette", pages: 235,}
    b2 := book{author: "Sun Tzu", title: "War art", pages: 127,}
    b3 := book{author: "Alan A. A. Donovan", title: "The Go book", pages: 302,}

    stack.Push(b1)
    stack.Push(b2)
    stack.Push(b3)

    return stack
}

func TestPeek(t *testing.T) {
    stack := buildStack()

    peek := stack.Peek()

    testBook, _ := peek.(book)

    if testBook.title != "The Go book" {
        t.Errorf("The title should be \"the Go book\" and it is %s.", testBook.title)
    }
}

func TestPush(t *testing.T) {
    stack := buildStack()

    myBook := book{author: "Edgare Allan Poe", title: "Sherlock Holmes: Shadow game", pages: 302,}
    stack.Push(myBook)

    val := stack.Peek()
    stack.Push(myBook)

    b, ok := val.(book)

    if !ok {
        t.Errorf("The returned value is not a book : %s", b.author)
    } else if b.title != myBook.title {
        t.Errorf("The title should be %s and it is %s", myBook.title, b.title)
    }
}

func TestPop(t *testing.T) {
    stack := &Stack{ Length: 0,}

    for i:=0; i<stack.Length + 5; i++ {
        stack.Pop()
    }

    if stack.Length < 0 {
        t.Errorf("The stack length shouldn't be lower than zero, but it is %d.", stack.Length)
    }
}
