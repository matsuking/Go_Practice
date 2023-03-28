package work

import (
	"fmt"
	"unsafe"
)

const secret = "abc"

type Os int

const (
	// 上から順に連番
	Mac Os = iota + 1
	Windows
	Linux
)

var (
	// i=0, s="", b=falseで初期化
	i int
	s string
	b bool
)

func pointer() {
	// fmt.Println("hello world.")
	// sl := []int{1, 2, 3}
	// if len(sl) > 2 {
	// 	fmt.Println("unreachable code")
	// }

	// 	godotenv.Load()
	// 	fmt.Println(os.Getenv("GO_ENV"))
	// 	fmt.Println(calculator.Offset)
	// 	fmt.Println(calculator.Sum(1, 2))
	// 	fmt.Println(calculator.Multiply(1, 2))

	var i int = 2
	fmt.Println(i)

	l := 1
	ui := uint16(2)
	fmt.Println(l)
	// %vは値をそのまま、%Tはデータ型を表示
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", l, ui)

	f := 1.23456
	s := "hello"
	b := true
	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("s: %[1]v %[1]T\n", s)
	fmt.Printf("b: %[1]v %[1]T\n", b)

	pi, title := 3.14, "GO"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	fmt.Printf("Mac: %v Windows: %v Linux: %v\n", Mac, Windows, Linux)

	var ui1 uint16
	// ui1のaddressを表示
	fmt.Printf("memory address of ui1: %p\n", &ui1)

	var ui2 uint16
	fmt.Printf("memory address of ui2: %p\n", &ui2)
	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("memory address of ui1(dereference): %v\n", *p1)

	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n", ui1)

	ok, result := true, "A"
	fmt.Printf("memory address of result: %p\n", &result)
	if ok {
		// このスコープ内だけで値の置き換え。 result = "B"とすると、値の置き換えが行われる。
		// 最初に定義したresultとは別のメモリー領域で定義される。
		result := "B"
		fmt.Printf("memory address of result: %p\n", &result)
		println(result)
	} else {
		result := "C"
		println(result)
	}
	println(result)
}
