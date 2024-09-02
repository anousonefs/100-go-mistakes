package main

import "time"

func main() {
	// we provided NewTicker with a duration of 1,000 nanoseconds = 1 microsecond.
	/* ticker := time.NewTicker(1000) */

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			println("haha")
		}
	}
}
