package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"aaa", "bbb"}
	nico := person{name: "nico", age: 12, favFood: favFood}

	fmt.Println(nico.favFood)
}
