package pool

import (
	"sync"
	_ "sync"
)

// RunPool обрабатывает задачи параллельно в заданном количестве воркеров
// и возвращает сумму результатов.
func RunPool(jobs []int, workers int) int {
	if workers < 1 {
		workers = 1
	}
	sum := 0
	ch := make(chan int)
	result := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for range workers {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for num := range ch {
					result <- num
				}
			}()
		}
		wg.Wait()
		close(result)
	}()
	go func() {
		for _, job := range jobs {
			ch <- job
		}
		close(ch)
	}()
	for range jobs {
		sum += <-result
	}
	return sum
}
