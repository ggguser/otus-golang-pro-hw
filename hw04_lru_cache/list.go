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

type list struct {
	count int
	front *ListItem
	back  *ListItem
}

// Len длина списка
func (l *list) Len() int {
	return l.count
}

// Front первый элемент списка
func (l *list) Front() *ListItem {
	if l.empty() {
		return nil
	}
	return l.front
}

// Back последний элемент списка
func (l *list) Back() *ListItem {
	if l.empty() {
		return nil
	}
	return l.back
}

// PushFront добавить значение в начало
func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Prev:  nil,
		Next:  l.Front(),
	}
	l.front.Prev = newItem
	l.front = newItem
	l.back = newItem
	if l.empty() {
		l.back = newItem
	}
	l.count++
	return newItem
}

// PushBack добавить значение в конец
func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Prev:  l.Back(),
		Next:  nil,
	}
	l.back.Next = newItem
	l.front = newItem
	if l.empty() {
		l.front = newItem
	}
	l.count++
	return newItem
}

// Remove удалить элемент
func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = l.front.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = l.back.Prev
	}
	l.count--
}

// MoveToFront переместить элемент в начало
func (l *list) MoveToFront(i *ListItem) {
	if l.count == 1 {
		return
	}
	l.Remove(i)
	l.front.Prev = i
	i.Next = l.front
	l.count++
}

func (l *list) empty() bool {
	return l.Len() <= 0
}

func NewList() List {
	return new(list)
}
