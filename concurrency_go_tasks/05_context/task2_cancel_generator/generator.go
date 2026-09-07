package generator

import "context"

// Generate возвращает канал, из которого можно читать возрастающие числа,
// начиная с нуля. Генерация прекращается при отмене ctx.
func Generate(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var i int
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			select {
			case <-ctx.Done():
				return
			default:
			case out <- i:
				i++
			}
		}
	}()

	return out
}
