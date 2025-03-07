package hw04lrucache

type list struct {
	firstItem  *ListItem
	lastItem   *ListItem
	countItems int
}

func NewList() ListInterface {
	return &list{}
}

func (l *list) Len() int {
	return l.countItems
}

func (l *list) Front() *ListItem {
	return l.firstItem
}

func (l *list) Back() *ListItem {
	return l.lastItem
}

func (l *list) PushFront(value interface{}) *ListItem {
	if l.isListEmpty() {
		pushedItem := l.pushInEmptyList(value)
		return pushedItem
	}

	pushedItem := ListItem{}
	pushedItem.Value = value
	pushedItem.Prev = nil
	pushedItem.Next = l.firstItem

	l.firstItem = &pushedItem
	l.firstItem.Next.Prev = &pushedItem

	l.countItems++
	return &pushedItem
}

func (l *list) PushBack(value interface{}) *ListItem {
	if l.isListEmpty() {
		pushedItem := l.pushInEmptyList(value)
		return pushedItem
	}

	pushedItem := ListItem{}
	pushedItem.Value = value
	pushedItem.Next = nil
	pushedItem.Prev = l.lastItem

	l.lastItem = &pushedItem
	l.lastItem.Prev.Next = &pushedItem

	l.countItems++
	return &pushedItem
}

func (l *list) Remove(toRemoveItem *ListItem) {
	if l.countItems == 1 {
		l.clearList()
		return
	}

	isFirstItemInList := toRemoveItem == l.firstItem
	if isFirstItemInList {
		l.removeFirstItemInList()
		return
	}

	isLastItemInList := toRemoveItem == l.lastItem
	if isLastItemInList {
		l.removeLastItemInList()
		return
	}

	l.removeMiddleItemInList(toRemoveItem)
}

func (l *list) MoveToFront(moveItem *ListItem) {
	if moveItem == l.firstItem {
		return
	}

	if moveItem == l.lastItem {
		preLastItem := l.lastItem.Prev
		preLastItem.Next = nil
		l.lastItem = preLastItem

		moveItem.Prev = nil
		moveItem.Next = l.firstItem

		l.firstItem = moveItem
		l.firstItem.Next.Prev = moveItem
		return
	}
	moveItem.Prev.Next = moveItem.Next
	moveItem.Next.Prev = moveItem.Prev

	moveItem.Prev = nil
	moveItem.Next = l.firstItem

	l.firstItem = moveItem
	l.firstItem.Next.Prev = moveItem
}

func (l *list) clearList() {
	l.firstItem = nil
	l.lastItem = nil
	l.countItems = 0
}

func (l *list) removeFirstItemInList() {
	secondItem := l.firstItem.Next
	secondItem.Prev = nil
	l.firstItem = secondItem
	l.countItems--
}

func (l *list) removeLastItemInList() {
	preLastItem := l.lastItem.Prev
	preLastItem.Next = nil
	l.lastItem = preLastItem
	l.countItems--
}

func (l *list) removeMiddleItemInList(toRemoveItem *ListItem) {
	toRemoveItem.Prev.Next = toRemoveItem.Next
	toRemoveItem.Next.Prev = toRemoveItem.Prev
	l.countItems--
}

func (l *list) isListEmpty() bool {
	return l.countItems == 0
}

func (l *list) pushInEmptyList(value interface{}) *ListItem {
	pushedItem := ListItem{}
	pushedItem.Value = value
	pushedItem.Prev = nil
	pushedItem.Next = nil

	l.firstItem = &pushedItem
	l.lastItem = &pushedItem
	l.countItems++
	return &pushedItem
}
