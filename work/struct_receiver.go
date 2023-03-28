package work

import (
	"fmt"
	"unsafe"
)

// 構造体はstructで作れる
type Task struct {
	Title    string
	Estimate int
}

func struct_receiver() {
	task1 := Task{
		Title:    "Learn Golang",
		Estimate: 3,
	}
	task1.Title = "learning Go"
	// %+[1]vで出力すると、field名も出力してくれる。
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

	var task2 Task = task1
	task2.Title = "new"
	fmt.Printf("task1: %v task2: %v\n", task1.Title, task2.Title)

	// &をつけることで先頭アドレスを取得, task1pはポインタ変数, *task1pは構造体の中身を出力
	task1p := &Task{
		Title:    "Learn concurrency",
		Estimate: 2,
	}
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	(*&task1p).Title = "Changed"
	fmt.Printf("task1p: %+v\n", *task1p)

	var task2p *Task = task1p
	task2p.Title = "Changed by Task2"
	fmt.Printf("task1: %+v\n", *task1p)
	fmt.Printf("task2: %+v\n", *task2p)

	task1.extendEstimate()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)

	(&task1).extendEstimatePointer()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)
}

// 値は変わらない。
func (task Task) extendEstimate() {
	task.Estimate += 10
}

// 値は変わる
func (taskp *Task) extendEstimatePointer() {
	taskp.Estimate += 10
}
