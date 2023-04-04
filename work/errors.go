

import (
	"errors"
	"fmt"
	"os"
)

var ErrCustom = errors.New("Not found")

func error() {
	err01 := errors.New("something wrong")
	err02 := errors.New("something wrong")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)

	// errorのインターフェースを持っているかどうかのチェックも行っている。
	fmt.Println(err01.Error())
	fmt.Println(err01)

	// errors.Newはそれぞれ違うメモリー領域に保存される。
	fmt.Println(err01 == err02)

	err0 := fmt.Errorf("add info: %w", errors.New("Original Error"))
	fmt.Printf("%[1]p %[1]T %[1]v\n", err0)

	// errorの内容を取り出す
	fmt.Println(errors.Unwrap(err0))
	fmt.Printf("%T\n", errors.Unwrap(err0))

	err1 := fmt.Errorf("add info: %v", errors.New("Original Error"))
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)
	// %vで設定しないと、nilが返ってくる？
	fmt.Println(errors.Unwrap(err1))

	err2 := fmt.Errorf("in repository layer: %w", ErrCustom)
	fmt.Println(err2)
	err2 = fmt.Errorf("in service layer: %w", err2)
	fmt.Println(err2)

	// 1つでも一致していればtrueを返す
	if errors.Is(err2, ErrCustom) {
		fmt.Println("matched")
	}

	file := "dummy.txt"
	err3 := fileChecker(file)

	if err3 != nil {
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Printf("%v file not found\n", file)
		} else {
			fmt.Println("unknown error")
		}
	}
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close()
	return nil
}
