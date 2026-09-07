package debounce

import "time"

// Debounce принимает значения и отдаёт только последнее после паузы d.
func Debounce(d time.Duration, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		var timer *time.Timer

		last, ok := <-in
		if !ok {
			return
		}
		timer = time.NewTimer(d)
		defer timer.Stop()
		for {
			if timer == nil {
				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					last = v
					timer = time.NewTimer(d)
				}
			} else {
				select {
				case v, ok := <-in:
					if !ok {
						if timer != nil {
							<-timer.C
							out <- last
						} else {
							out <- last
						}
						return
					}
					last = v
					if timer != nil {
						timer.Stop()
					}
					timer = time.NewTimer(d)
				case <-timer.C:
					out <- last
					timer = nil
				}
			}
		}
	}()

	return out
}
