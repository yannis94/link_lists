package queue

import (
	"testing"
)

type book struct {
    author string
    title string
    pages int
}

func getBookQueue() *Queue {
    queue := BuildQueue()

    book1 := &book{ author: "Victor Hugo", title: "Les miserables", pages: 289, }
    book2 := &book{ author: "W. Shakespeare", title: "Romeo & Juliette", pages: 212, }
    book3 := &book{ author: "Edgar Allan Poe", title: "Sherlock Holmes", pages: 333, }
    book4 := &book{ author: "Moliere", title: "Les fourberies de Scapin", pages: 105, }
    book5 := &book{ author: "Emile Zola", title: "Therese Raquin", pages: 453, }

    queue.Append(book1)
    queue.Append(book2)
    queue.Append(book3)
    queue.Append(book4)
    queue.Append(book5)

    return queue
}

func TestAppend(t *testing.T) {
    testQ := BuildQueue()

    book1 := &book{ author: "Victor Hugo", title: "Les miserables", pages: 289, }
    book2 := &book{ author: "W. Shakespeare", title: "Romeo & Juliette", pages: 212, }

    testQ.Append(book1)
    testQ.Append(book2)

    if testQ.Length != 2 {
        t.Errorf("The queue should have a length of 2 and it is %d.", testQ.Length)
    }

}

func TestPrepend(t *testing.T) {
    testQ := getBookQueue()

    originalLength := testQ.Length

    newBook := &book{ author: "Sun Tzu", title: "War art", pages: 178, }
    testQ.Prepend(newBook)

    if testQ.Length != originalLength+1 {
        t.Errorf("The length should be increment by one.")
    }
}

func TestInsertAt(t *testing.T) {
    testQ := getBookQueue()
    origLength := testQ.Length

    newBook := &book{ author: "Sun Tzu", title: "War art", pages: 178, }

    err := testQ.InsertAt(30, newBook)

    if err == nil {
        t.Errorf("This should raise an error, attempt to insert at index out of range.")
    }

    testQ.InsertAt(3, newBook)

    if testQ.Length != origLength+1 {
        t.Errorf("The length should be increment by one.")
    }
}

func TestRemoveAt(t *testing.T) {
    testQ := getBookQueue()

    initialLength := testQ.GetLength()

    _, err := testQ.RemoveAt(29)

    if err == nil {
        t.Errorf("This should raise an error, attempt to insert at index out of range.")
    }

    testQ.RemoveAt(2)

    if testQ.Length != initialLength {
        t.Errorf("The length should be %d and it is %d.", initialLength, testQ.Length)
    }
}
