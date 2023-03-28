package work

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func funcDefer() {
	// 複数defer文がある場合は、下から上に実行される。
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}

// 複数returnしたい場合は、()でくくる。
func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		// errors.Newでエラー文を作成
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}

func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}

	return out
}

// 第1引数は無名関数
func addExt(f func(file string) string, name string) {
	fmt.Println(f(name))
}

func multipley() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}

func countUp() func(int) int {
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

func function_closure() {
	funcDefer()
	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(trimExtension(files...))
	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	// 無名関数
	i := 1
	func(i int) {
		fmt.Println(i)
	}(i) // 最後に（）をつけると、即座に実行される

	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))

	f2 := func(file string) string {
		return file + ".csv"
	}
	addExt(f2, "file1")

	f3 := multipley()
	fmt.Println(f3(2))

	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(2)
		fmt.Printf("%v\n", v)
	}
}
