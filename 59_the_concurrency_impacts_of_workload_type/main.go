package main

import (
	"io"
	"sync"
	"sync/atomic"
)

func main() {

}

func read(r io.Reader) (int, error) {
	count := 0
	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		count += task(b)
	}
	return count, nil
}

func task(b []byte) int {
	count := 0
	for _, v := range b {
		if !isWhitespace(v) {
			count++
		}
	}
	return count
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func readV2(r io.Reader) (int, error) {
	var count int64
	wg := sync.WaitGroup{}
	var n = 10 // I/O bound. the value mainly depends on extermal system
	/* n := runtime.GOMAXPROCS(0) */ // CPU-bound
	ch := make(chan []byte, n)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}
	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		ch <- b
	}
	close(ch)
	wg.Wait()

	return int(count), nil
}

/* When implementing the worker-pooling pattern, we have seen that the optimal number of goroutines in the pool */
/* depends on the workload type. If the workload executed by the workers is I/O-bound, the value mainly depends on the external system.
Conversely, if the workload is CPU-bound, the optimal number of goroutines is close to the number of available threads.
Knowing the workload type (I/O or CPU) is crucial when designing concurrent applications. */
