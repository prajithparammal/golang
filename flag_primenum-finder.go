package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		prime   []int
		IsPrime bool = true
	)
	num1 := flag.Int("num1", 0, "First Number")
	num2 := flag.Int("num2", 0, "Second Number")
	flag.Parse()

	if *num1 == 0 || *num2 == 0 {
		flag.PrintDefaults()
		return
	}

	for i := *num1; i <= *num2; i++ {
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				IsPrime = false
				break
			}
		}
		if IsPrime == true {
			prime = append(prime, i)
		}
		IsPrime = true
	}

	fmt.Printf("Prime number within the range [%d - %d]\n%v", *num1, *num2, prime)
}
