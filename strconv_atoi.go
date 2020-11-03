package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage = `
	Usage: main.go <2> <10>
	Pass two number to find prime number within the range`
)

func main() {
	var (
		prime   []int
		IsPrime bool
	)
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println(usage)
		return
	}

	num1, err := strconv.Atoi(arg[0])
	if err != nil {
		fmt.Printf("%q is not vaild\n", arg[0])
		return
	}

	num2, err := strconv.Atoi(arg[1])
	if err != nil {
		fmt.Printf("%q is not vaild\n", arg[1])
		return
	}

	for i := num1; i <= num2; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				IsPrime = false

			}
		}
		if IsPrime == true {
			prime = append(prime, i)
		}
		IsPrime = true
	}

	fmt.Printf("Prime number within the range [%d - %d]\n%v", num1, num2, prime)
}
