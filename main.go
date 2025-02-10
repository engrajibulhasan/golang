package main

import (
	"fmt"
	"hello_world/integers"
	"hello_world/iterations"
)

func FindPrime(n int) []int {
	if n <= 2 {
		return nil
	}

	sieveList := make([]bool, n+1)
	for x := 2; x < n; x++ {
		sieveList[x] = true
	}
	for curr := 2; curr < n; curr++ {
		if sieveList[curr] {
			for multiple := curr * curr; multiple < n; multiple += curr {
				sieveList[multiple] = false
			}
		}
	}

	var primeList []int
	for curr := 0; curr < n; curr++ {
		if sieveList[curr] {
			primeList = append(primeList, curr)
		}
	}

	return primeList
}

func main(){

	fmt.Println("Hello world")
	fmt.Println(integers.Subs(5,2) )
	fmt.Println(iterations.Repeat("a") )

	// Array class

	// x := []int{1, 2, 3, 4, 5}
	// y := []int{1, 2, 3, 4, 5}
	// z := []int{1, 2, 3, 4, 5, 6}
	// s := []string{"a", "b", "c"}
	// fmt.Println(slices.Equal(x, y)) // prints true
	// fmt.Println(slices.Equal(x, z)) // prints false
	// fmt.Println(slices.Equal(x, s)) // does not compile


	// var value = []int{} // both are ok
	// value = append(x, 36) // appending in a nil slice
	
	// var value = []int{65, 43, 90}
	// value = append(x, 36) // this works too
	
	// // This creates an `int` slice with a length of 3 and a capacity of 3.
	x := make([]int, 5,5) 
	x = append(x, 20)
	fmt.Println(x)
 
	
	
	fmt.Println(FindPrime(50))
	// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47]
	
	// //This creates an `int` slice with a length of 5 and a capacity of 10
	// x := make([]int, 5, 10)
	
	// // slice with zero length but a capacity that’s greater than zero:
	// x := make([]int, 0, 10)
}