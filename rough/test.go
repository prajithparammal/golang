package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 2 // value 3
	j := 5 // value 6

	// s1[:i] ==> [1,2]
	// s1[j:] ==> [6,7,8,9,10]
	s1 = append(s1[:i], s1[j:]...)
	fmt.Println(s1)
	// so here we removed [3,4,5]. I.e, removed s1[2],s1[3],s1[4]
}
