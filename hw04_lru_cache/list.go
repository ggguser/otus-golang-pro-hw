package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Items []*ListItem
}

func (l *list) Len() int {
	return len(l.Items)
}

func (l *list) Front() *ListItem {
	return l.Items[0]
}

func (l *list) Back() *ListItem {
	return l.Items[len(l.Items)-1]
}

func NewList() List {
	return new(list)
}
