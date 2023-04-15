package queue

import "testing"

type book struct {
    author string
    title string
    pages int
}

func getTestingQueue() *Queue {
    queue := BuildQueue()

    book1 := &book{ author: "Victor Hugo", title: "Les miserables", pages: 289, }
    book2 := &book{ author: "W. Shakespeare", title: "Romeo & Juliette", pages: 212, }
    book3 := &book{ author: "Edgar Allan Poe", title: "Sherlock Holmes", pages: 333, }
    book4 := &book{ author: "Moliere", title: "Les fourberies de Scapin", pages: 105, }
    book5 := &book{ author: "Emile Zola", title: "Therese Raquin", pages: 453, }

    queue.Enqueue(book1)
    queue.Enqueue(book2)
    queue.Enqueue(book3)
    queue.Enqueue(book4)
    queue.Enqueue(book5)
    
    return queue
}

func TestEnqueue(t *testing.T) {
    testQ := getTestingQueue()
    firstLength := testQ.Length

    testBook := &book{ author: "John Doe", title: "This book does not exist", pages: 60, }
    testQ.Enqueue(testBook)

    if testQ.Length != firstLength+1 {
        t.Errorf("Queue's length should be %d and not %d.", firstLength, testQ.Length)
    }
}

func TestDequeue(t *testing.T) {
    testQ := getTestingQueue()
    firstLength := testQ.Length

    testQ.Dequeue()

    if testQ.Length != firstLength-1 {
        t.Errorf("Queue's length should be %d and not %d.", firstLength-1, testQ.Length)
    }

}
