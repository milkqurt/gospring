package main

import (
	"fmt"
	"reflect"
)

func main() {
	message1 := "test"
	message1 = "test2"
	var message2 string = "test3"
	const message3 string = "test4"

	fmt.Println(message1)
	fmt.Println(reflect.TypeOf(message2))
	fmt.Println("________________________________")

	var str string
	var number int8
	// uint + числа
	// int8 по байтам
	var b bool
	byt := []byte("asd")
	//byte 0 to 250 asci table
	var by2 byte = 62
	var run rune = 'a'
	a1, b1, c1 := 1, 2, 3
	a1, b1 = b1, a1
	num7, str5 := 7, "string"

	fmt.Println(str)
	fmt.Println(number)
	fmt.Println(b)
	fmt.Println(byt)
	fmt.Printf("%c\n", by2)
	fmt.Println(run)
	fmt.Println(a1, b1, c1)
	fmt.Println(num7, str5)
}
