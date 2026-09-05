package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	tokens chan struct{}
	stop   chan struct{}
	once   sync.Once
}

func NewLimiter() *Limiter {
	l := &Limiter{
		tokens: make(chan struct{}, 5),
		stop:   make(chan struct{}),
	}
	for i := 0; i < 5; i++ {
		l.tokens <- struct{}{}
	}

	ticker := time.NewTicker(time.Second / 5)
	go func() {
		for {
			select {
			case <-ticker.C:
				select {
				case l.tokens <- struct{}{}:
				default:
				}
			case <-l.stop:
				ticker.Stop()
				close(l.tokens)
				return
			}
		}
	}()
	return l
}

func (l *Limiter) Allow() bool {
	select {
	case _, ok := <-l.tokens:
		return ok
	default:
		return false
	}
}

func (l *Limiter) Stop() {
	l.once.Do(func() {
		close(l.stop)
	})
}
