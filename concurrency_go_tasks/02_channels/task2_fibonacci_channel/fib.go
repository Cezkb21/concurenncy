package fibonacci

// Fib возвращает канал, из которого можно читать первые n чисел Фибоначчи.
func Fib(n int) <-chan int {
	ch := make(chan int)
	a, b := 1, 1
	go func() {
		if n > 0 {
			ch <- 0
		}
		if n > 1 {
			ch <- 1
		}
		if n > 2 {
			ch <- 1
		}
		for i := 3; i < n; i++ {
			ch <- a + b
			a, b = b, a+b
		}
		close(ch)
	}()

	return ch
}
