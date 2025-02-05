package hw04lrucache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var debug_flag = true

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("one element", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		debug(l)

		require.Equal(t, 1, l.Len())
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 1, l.Back().Value)
	})

	t.Run("test PushFront count", func(t *testing.T) {
		l := NewList()

		l.PushFront(3)
		l.PushFront(2)
		l.PushFront(1)

		require.Equal(t, 3, l.Len())
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 3, l.Back().Value)
	})

	t.Run("test PushFront links", func(t *testing.T) {
		l := NewList()

		l.PushFront(3)
		l.PushFront(2)
		l.PushFront(1)

		require.Equal(t, l.Front(), l.Front().Next.Prev)
		require.Equal(t, l.Back(), l.Back().Prev.Next)
	})

	t.Run("test PushBack", func(t *testing.T) {
		l := NewList()
		debug(l)
		l.PushBack(1)
		debug(l)
		l.PushBack(2)
		debug(l)

		require.Equal(t, 2, l.Len())
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 2, l.Back().Value)
	})

	t.Run("test Remove middle", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		debug(l)
		l.PushBack(20) // [10, 20]
		debug(l)
		l.PushBack(30) // [10, 20, 30]
		debug(l)
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		debug(l)
		require.Equal(t, 2, l.Len())
		require.Equal(t, l.Back(), l.Front().Next)
		require.Equal(t, l.Front(), l.Back().Prev)
	})

	t.Run("test MoveToFront back to front — order", func(t *testing.T) {
		l := NewList()

		l.PushFront(3)
		l.PushFront(2)
		l.PushFront(1)

		l.MoveToFront(l.Back())
		require.Equal(t, []any{3, 1, 2}, getItems(l))
		require.Equal(t, 2, l.Back().Value)

	})

	t.Run("test MoveToFront front to front — order", func(t *testing.T) {
		l := NewList()

		l.PushFront(3)
		l.PushFront(2)
		l.PushFront(1)

		l.MoveToFront(l.Front())
		require.Equal(t, []any{1, 2, 3}, getItems(l))
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		debug(l)
		l.PushBack(20) // [10, 20]
		debug(l)
		l.PushBack(30) // [10, 20, 30]
		debug(l)
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		debug(l)
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
				// 0 -> [40, 10, 30]
				// 2 -> [60, 40, 10, 30, 50]
				// 4 -> [80, 60, 40, 10, 30, 50, 70]

			} else {
				l.PushBack(v)
				// 1 -> [40, 10, 30, 50]
				// 3 -> [60, 40, 10, 30, 50, 70]
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		require.Equal(t, []any{80, 60, 40, 10, 30, 50, 70}, getItems(l))

		l.MoveToFront(l.Back()) // [70, 80, 60, 40, 10, 30, 50]
		require.Equal(t, []any{70, 80, 60, 40, 10, 30, 50}, getItems(l))

		require.Equal(t, 70, l.Front().Value)
		require.Equal(t, 50, l.Back().Value)

		require.Equal(t, []any{70, 80, 60, 40, 10, 30, 50}, getItems(l))
	})
}

// PrintList выдаёт слайс всех элементов списка
func getItems(dll *DoublyLinkedList) []any {
	items := make([]any, 0, dll.Len())
	for i := dll.Front(); i != nil; i = i.Next {
		items = append(items, i.Value)
	}
	return items
}

func debug(dll *DoublyLinkedList) {
	if debug_flag {
		items := getItems(dll)
		fmt.Printf("%v — %d\n", items, dll.count)
	}
}
