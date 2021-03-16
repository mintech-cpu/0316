package main

import (
	"fmt"
	"strconv"
)

func main() {
	t, f := true, false
	fmt.Printf("%T %v", t, t)
	fmt.Printf("%T %v", f, f)

	// intからfloatへの変更
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f", xx, xx, xx)
	// floatからintへの変更
	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d", yy, yy, yy)

	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v", i, i)

	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// var b [2]int = [2]int{100, 200}
	// fmt.Println(b)

	// スライス
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)

	// スライスの中にスライスを入れる
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

	c := make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}

	d := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}

	// キーがstringでバリュー(値)がint
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)

	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	// 値があるかの確認
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// 空のmapを作ってから値を入れる
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

}
