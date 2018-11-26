package main

import "fmt"

func main() {

	userData := map[string]int{
		"age": 40,
	}

	if val, ok := userData["age"]; ok {
		fmt.Printf("User age is %v", val)
	}
}
