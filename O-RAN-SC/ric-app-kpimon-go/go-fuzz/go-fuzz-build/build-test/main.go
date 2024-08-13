package main

import "fmt"

func testSwitch(x int) {
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}
}

func testTypeSwitch(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case bool:
	default:
		fmt.Println("Unknown")
	}
}

func testLoops() {
	for i := 0; i < 5; i++ {
		fmt.Println("Loop", i)
	}

	numbers := []int{1, 2, 3}
	for _, num := range numbers {
		fmt.Println("Range", num)
	}

	goto end // goto end jumps over declaration of

	for i := 0; i < 5; i++ {
		fmt.Println("Loop", i)
	}

end:
}

func testIfElse(x int) {
	if x < 0 {
		fmt.Println("Negative")
	} else if x == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Positive")
	}
}

func testGoto() {
	i := 0
loop:
	if i < 5 {
		fmt.Println("Goto", i)
		i++
		goto loop
	}
}

func main() {
	testSwitch(1)
	testTypeSwitch("hello")
	testLoops()
	testIfElse(10)
	testGoto()
}
