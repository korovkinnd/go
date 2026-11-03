//go:build !solution

package oncecache

import "sync"

type Cache interface {
	Get(key string) (any, bool)
}

type entry struct {
	once  sync.Once
	value any
}

type cache struct {
	load    func(string) any
	mu      sync.Mutex
	entries map[string]*entry
}

func New(load func(string) any) Cache {
	return &cache{
		load:    load,
		entries: make(map[string]*entry),
	}
}

func (c *cache) Get(key string) (any, bool) {
	c.mu.Lock()
	e, ok := c.entries[key]
	if !ok {
		e = &entry{}
		c.entries[key] = e
	}
	c.mu.Unlock()

	e.once.Do(func() {
		e.value = c.load(key)
	})

	return e.value, true
}
