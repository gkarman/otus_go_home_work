package hw04lrucache

type ListItem struct {
	Value interface{}
	Key   Key
	Prev  *ListItem
	Next  *ListItem
}
