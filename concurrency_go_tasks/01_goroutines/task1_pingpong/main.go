package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// PingPong должен запускать две горутины "ping" и "pong",
// которые поочередно выводят строки пять раз каждая.
// Реализуйте синхронизацию через каналы и ожидание завершения.
func PingPong(w io.Writer) {
	wg := sync.WaitGroup{}
	ping := make(chan struct{})
	pong := make(chan struct{})
	wg.Add(2)

	go func() {
		for range 5 {
			<-ping
			fmt.Fprintln(w, "pong")
			pong <- struct{}{}
		}
		wg.Done()
	}()

	go func() {
		for range 5 {
			fmt.Fprintln(w, "ping")
			ping <- struct{}{}
			<-pong
		}
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	PingPong(os.Stdout)
}
