package datastructure

type LinkList interface {
    GetLength() int
    InsertAt(idx int, data interface{}) error
    RemoveAt(idx int) (interface{}, error)
    Append(data interface{}) error
    Prepend(data interface{}) error
    Get(idx int) (interface{}, error)
}
