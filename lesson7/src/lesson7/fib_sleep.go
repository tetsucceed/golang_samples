package lesson7

import (
	"fmt"
	"time"
)

func DoSpinner() {
	go spinner(50 * time.Millisecond)
	const n = 42
	time.Sleep(10 * time.Second)
	//fibN := fibonacci(n)
	//fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
