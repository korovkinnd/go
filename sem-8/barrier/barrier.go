//go:build !solution

package barrier

import "sync"

type Barrier interface {
	Wait()
	Done()
}

type barrier struct {
	mu      sync.Mutex
	cond    *sync.Cond
	parties int
	arrived int
	round   uint64
}

func New(n int) Barrier {
	b := &barrier{parties: n}
	b.cond = sync.NewCond(&b.mu)

	return b
}

func (b *barrier) Wait() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.arrived++
	if b.arrived >= b.parties {
		b.release()
		return
	}

	round := b.round
	for round == b.round {
		b.cond.Wait()
	}
}

func (b *barrier) Done() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.parties--
	if b.arrived > 0 && b.arrived >= b.parties {
		b.release()
	}
}

func (b *barrier) release() {
	b.arrived = 0
	b.round++
	b.cond.Broadcast()
}
