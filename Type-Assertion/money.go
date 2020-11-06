package main

import "fmt"

//A type and its methods should be in the same package.

type money float64

func (m money) string() string {
	return fmt.Sprintf("$%.2f", m)
}
