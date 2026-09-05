package scheduler

import (
	"sync"
	"time"
)

// Every запускает f каждые d и возвращает функцию для остановки.
func Every(d time.Duration, f func()) (stop func()) {
	if d < time.Millisecond*5 {
		d = time.Millisecond * 5
	}
	ticker := time.NewTicker(d)
	stopCh := make(chan struct{})
	once := sync.Once{}
	stop = func() {
		once.Do(func() {
			close(stopCh)
		})
	}
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			case <-stopCh:
				ticker.Stop()
				return
			}
		}
	}()
	return stop
}
