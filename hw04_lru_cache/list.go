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
	Prev  *ListItem
	Next  *ListItem
}

type DoublyLinkedList struct {
	count int
	front *ListItem
	back  *ListItem
}

// Len длина списка
func (dll *DoublyLinkedList) Len() int {
	return dll.count
}

// Front первый элемент списка
func (dll *DoublyLinkedList) Front() *ListItem {
	return dll.front
}

// Back последний элемент списка
func (dll *DoublyLinkedList) Back() *ListItem {
	return dll.back
}

// PushFront добавить значение в начало
func (dll *DoublyLinkedList) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Prev:  nil,
		Next:  dll.Front(),
	}
	if dll.isEmpty() {
		dll.back = newItem
		dll.front = newItem
	} else {
		dll.front.Prev = newItem
		dll.front = newItem
	}
	dll.count++
	return newItem
}

// PushBack добавить значение в конец
func (dll *DoublyLinkedList) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Prev:  dll.Back(),
		Next:  nil,
	}
	if dll.isEmpty() {
		dll.back = newItem
		dll.front = newItem
	} else {
		dll.back.Next = newItem
		dll.back = newItem
	}
	dll.count++
	return newItem
}

// Remove удалить элемент
func (dll *DoublyLinkedList) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		dll.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		dll.back = i.Prev
	}
	i.Prev = nil
	i.Next = nil
	dll.count--
}

// MoveToFront переместить элемент в начало
func (dll *DoublyLinkedList) MoveToFront(i *ListItem) {
	if i == dll.Front() {
		return
	}
	dll.Remove(i)
	i.Prev = nil
	i.Next = dll.front
	if dll.Front() != nil {
		dll.front.Prev = i
	}
	dll.front = i
	if dll.Back() != nil {
		dll.back = i
	}
	dll.count++
}

func (dll *DoublyLinkedList) isEmpty() bool {
	return dll.Len() <= 0
}

func NewList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}
