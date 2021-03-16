package main

import "fmt"

func add(x, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x - y
}

// func cal(price, item int) int{
// 	result := price * item
// 	return result
// }
func cal(price, item int) (result int) {
	// (result int)と変数名を描いてあげることで、どんな変数を返すのかが分かりやすい
	result = price * item
	return
}

func main() {
	r1, r2 := add(110, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 3)
	fmt.Println(r3)

	// 変数にfunctionを入れる
	f1 := func(x int) {
		fmt.Println("inner", x)
	}
	f1(1)
	// 書き換え
	func(x int) {
		fmt.Println("inner", x)
	}(1)

	var f = 1.11
	ff := int(f)
	fmt.Printf("%T %v %d", ff, ff, ff)

}
