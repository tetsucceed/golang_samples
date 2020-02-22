package utils

import (
	"fmt"
	"math/big"
)

func Isdiv(a int) {
	if a%2 == 0 {
		fmt.Println("Number is 2k")
	} else {
		fmt.Println("Number is 2k + 1")
	}
}

func Isdiv3(a int) {
	if a%3 == 0 {
		fmt.Println("Number is divisible by 3")
	} else {
		fmt.Println("Number is not divisible by 3")
	}
}

func Fibb100() {
	var iter int32
	fib1 := big.NewInt(0)
	fib2 := big.NewInt(1)

	fmt.Println("Fibonacci 100")
	fmt.Println(fib1) // first one

	// 0 to 99, cause first was printed
	for iter = 0; iter < 99; iter++ {
		fib1.Add(fib1, fib2)
		fib1, fib2 = fib2, fib1
		fmt.Println(fib1)
	}
}

func Primes() {
	var t [100]big.Int
	var iter int32

	prime := big.NewInt(0)
	iter = 0

	for iter < 100 {
		prime.Add(prime, big.NewInt(1))
		if prime.ProbablyPrime(0) {
			t[iter].Set(prime)
			iter += 1
		} else {
			fmt.Println("it's not prime")
		}
	}

	iter = 0
	for iter = 0; iter < 100; iter++ {
		fmt.Println(t[iter].String())
	}

}
