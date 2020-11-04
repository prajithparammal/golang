package main

import "fmt"

type printer interface {
	print() // method name without code
}

//type list []*game
type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We are waiting for delivery")
		return
	}
	for _, it := range l {
		fmt.Printf("%T --->", it)
		it.print()
	}
}
