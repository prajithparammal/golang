package main

func main() {
	var (
		harrypotter = book{"Harry Potter", 500}
		mahabarth   = book{"Mahabaratham", 650.50}
		pubg        = game{"PUBG", 200}
		cod         = game{"Call of Duty", 300}
		rubics      = puzzle{"rubis cube", 25}
	)

	//go will pass the value of a book automatactially to these methods, methods are functions inside a type.
	//Method belongs to Type and function below to package.
	// to change input parameter in a function, we use pointers
	// you can attach methods to almost any type

	var store list
	store = append(store, &pubg, &cod, mahabarth, harrypotter, rubics)
	store.print()

}
