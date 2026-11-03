//go:build !solution

package condbuffer

import "sync"

type Buffer interface {
	Put(x any)
	Get() any
	Len() int
}

type buffer struct {
	mu       sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond
	items    []any
	capacity int
}

func New(capacity int) Buffer {
	b := &buffer{capacity: capacity}
	b.notFull = sync.NewCond(&b.mu)
	b.notEmpty = sync.NewCond(&b.mu)

	return b
}

func (b *buffer) Put(x any) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == b.capacity {
		b.notFull.Wait()
	}

	b.items = append(b.items, x)
	b.notEmpty.Signal()
}

func (b *buffer) Get() any {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.notEmpty.Wait()
	}

	x := b.items[0]
	b.items[0] = nil
	b.items = b.items[1:]
	b.notFull.Signal()

	return x
}

func (b *buffer) Len() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return len(b.items)
}
