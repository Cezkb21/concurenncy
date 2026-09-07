package pipeline

func Run(nums []int) int {
	in := make(chan int)
	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()
	squareOut := make(chan int)
	go func() {
		for n := range in {
			squareOut <- n * n
		}
		close(squareOut)
	}()
	doubleOut := make(chan int)
	go func() {
		for n := range squareOut {
			doubleOut <- n * 2
		}
		close(doubleOut)
	}()
	sum := 0
	for n := range doubleOut {
		sum += n
	}
	return sum
}
