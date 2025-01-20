package main

import (
	"errors"
	"fmt"
	"log"
)

//var a,b,c int

func main() {
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)
	a = 15
	print(a, b, c)
	var name string = "Alice"
	fmt.Println(sayMessage(name))
	fmt.Println(sayStringInt(name, a))
	message, res, err := checkAge(a)
	if err != nil {
		log.Fatal(err)
	}
	// _ ignore
	fmt.Println(message, res)
}

func print(a int, b int, c int) {
	fmt.Println(a, b, c)
}

func sayMessage(name string) string {
	return "Hello " + name
}

func sayStringInt(name string, age int) string {
	result := fmt.Sprintf("Hello %s, this age %d", name, age)
	return result
}

func checkAge(age int) (string, bool, error) {
	if age >= 18 && age <= 24 {
		response := "Yes 18"
		return response, true, nil
	} else if age >= 25 && age <= 35 {
		response := "Yes 25 or 35"
		return response, true, nil
	} else {
		response := "No 18"
		return response, false, errors.New("Invalid Age")
	}
}
