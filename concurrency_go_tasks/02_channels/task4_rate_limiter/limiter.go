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

var limit int = 5

func NewLimiter() *Limiter {
	limiter := &Limiter{
		tokens: make(chan struct{}, limit),
		stop:   make(chan struct{}),
	}
	for range limit {
		limiter.tokens <- struct{}{}
	}

	go func() {
		delta := time.Second / time.Duration(limit)
		ticker := time.NewTicker(delta)
		defer ticker.Stop()
		for {
			select {
			case <-limiter.stop:
				return
			case <-ticker.C:
				select {
				case limiter.tokens <- struct{}{}:
				default:
				}
			}
		}
	}()

	return limiter
}

func (l *Limiter) Allow() bool {
	select {
	case <-l.stop:
		return false
	default:
	}

	select {
	case <-l.tokens:
		return true
	default:
		return false
	}
}

func (l *Limiter) Stop() {
	l.once.Do(func() {
		close(l.stop)
	})
}
