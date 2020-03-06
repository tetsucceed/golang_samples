package lesson7

import (
	"fmt"
)

func Pipe() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; ; x++ {
			if x > 10000 {
				close(naturals)
				return
			}
			naturals <- x
		}
	}()

	// возведение в квадрат
	go func() {
		for {
			x := <-naturals
			if x > 10000 {
				close(squares)
				return
			}
			squares <- x * x
		}
	}()

	// печать
	for {
		b := <-squares
		if b > 10000 {
			break
		}
		fmt.Println(<-squares)
	}
}
