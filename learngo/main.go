package main

import "fmt"

func main() {
	nico := map[string]string{"name": "aaa", "age": "20"}
	for _, value := range nico {
		fmt.Println(value)
	}
}
