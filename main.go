package main

import (
	"fmt" // osの1階層下にuserがある
	"strings"
	"time"
)

// func bazz() {
// 	fmt.Println("Bazz")
// }

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println("HelloWorld!!!", time.Now())
	fmt.Println("HelloWorld")
	fmt.Println("Hello" + "World")
	fmt.Println(string("HelloWorld"[0]))
	var s string = "HelloWorld"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(strings.Contains(s, "Word")) // "Word"が含まれているか？
	fmt.Println("TestTest")

	// bazz()

	// 関数の外でも宣言できる
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// 関数の中でしか宣言できない
	xi := 1
	xf64 := 1.2
	fmt.Println(xi, xf64)
	fmt.Printf("%T", xf64)
}
