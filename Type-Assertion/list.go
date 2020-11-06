package main

import "fmt"

type printer interface {
	print() // method name without code
	discount(ratio float64)
}

//type list []*game
type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We are waiting for delivery")
		return
	}
	for _, it := range l {
		it.print()
	}

}

func (l list) discount(ratio float64) {
	for _, it := range l {
		g, isGame := it.(interface{ discount(float64) })
		if !isGame {
			continue
		}
		g.discount(ratio)
	}
}
