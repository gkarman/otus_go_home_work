package hw04lrucache

type Key string

type lruCache struct {
	capacity int
	queue    ListInterface
	items    map[Key]*ListItem
}

func NewCache(capacity int) CacheInterface {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	v, isExists := c.items[key]
	if isExists {
		v.Value = value
		c.queue.PushFront(value)
		return true
	}

	listItem := c.queue.PushFront(value)
	listItem.Key = key
	c.items[key] = listItem

	if c.queue.Len() > c.capacity {
		oldestKey := c.queue.Back().Key
		delete(c.items, oldestKey)
		c.queue.Remove(c.queue.Back())
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	v, isExists := c.items[key]
	if isExists {
		return v.Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
}
