package work

import (
	"fmt"
	"time"
)

func for_loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for {
	// 	fmt.Printf("working")
	// 	time.Sleep(2 * time.Second)
	// }

	var i int
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(300 * time.Millisecond)
	}

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			continue
		case 3:
			continue
		case 8:
			// 28行目につけた名前を指定することで、正確に抜けることができる。
			break loop
		default:
			fmt.Printf("%v ", i)
		}
	}
	fmt.Printf("\n")

	items := []item{
		{price: 10.},
		{price: 20.},
		{price: 30.},
	}

	// これは値がコピーされるだけで、値は変更されない。
	for _, i := range items {
		i.price *= 1.1
	}
	fmt.Printf("%+v\n", items)

	// インデックスに直接アクセスした場合は、値が変更される。
	for i := range items {
		items[i].price *= 1.1
	}
	fmt.Printf("%+v\n", items)
}

type item struct {
	price float32
}
