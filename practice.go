package main

import (
	"fmt"
)

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}
	fmt.Println(result)

	m := map[string]int{
		"Mike":  30,
		"Nancy": 20,
		"meisa": 50,
	}
	fmt.Printf("%T %v", m, m)

	// if文
	num := 11
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// continueとbreak
	for i1 := 0; i1 < 10; i1++ {
		if i1 == 3 {
			fmt.Println("continue")
			continue
		}
		if i1 > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i1)

	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

}
