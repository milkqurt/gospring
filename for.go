package main

import "fmt"

func main() {
	fmt.Println(findMax(22, 33, 11, 44, 55))
	func() {
		fmt.Println("Hello World")
	}()
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	
}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return max
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
