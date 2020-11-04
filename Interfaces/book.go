package main

import "fmt"

type book struct {
	title string
	price money
}

func (b book) print() {
	// fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}

//book type should have Data(title and price ) and Behaviour(methods or functions) like
//printBook / discount etc
// func (b book) printBook() {} ; here printBook is a method of book type
// ie can call it as harrypotter.printBook() ; harrypotter is a book type variable
// A receiver is an input paramater written before a function name.
