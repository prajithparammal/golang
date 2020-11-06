package main

import "fmt"

type game struct {
	title string
	price money
}

func (g *game) print() {
	//fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64) {
	//g.price *= (1 - ratio)
	g.price *= money(1 - ratio)

}
