package pipelinectx

import (
	"context"
	"sync"
)

// Run строит конвейер из двух стадий: удвоение и суммирование.
// Конвейер должен останавливаться, если ctx отменён.
// Возвращает итоговую сумму и ошибку контекста при отмене.
func Run(ctx context.Context, nums []int) (int, error) {
	wg := sync.WaitGroup{}
	in := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(in)
		select {
		case <-ctx.Done():
			return
		default:
		}
		for _, n := range nums {
			select {
			case in <- n:
			case <-ctx.Done():
				return
			}
		}
	}()

	doubleOut := make(chan int)
	go func() {
		defer wg.Done()
		defer close(doubleOut)
		select {
		case <-ctx.Done():
			return
		default:
		}
		for n := range in {
			select {
			case doubleOut <- n * 2:
			case <-ctx.Done():
				return
			}
		}
	}()

	sum := 0
	for n := range doubleOut {
		sum += n
	}
	wg.Wait()
	return sum, ctx.Err()
}
