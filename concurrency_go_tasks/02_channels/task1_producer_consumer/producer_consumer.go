package producerconsumer

import (
	"fmt"
	"io"
	"strconv"
	"sync"
)

// Run запускает продюсера, который отправляет числа от 1 до 10, и консюмера,
// который выводит их в writer. Используйте небуферизованный канал и ожидание
// завершения горутин.
func Run(w io.Writer) {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range channel {
			fmt.Fprintln(w, strconv.Itoa(i+1))
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 10 {
			channel <- i
		}
		close(channel)
	}()
	wg.Wait()
}
