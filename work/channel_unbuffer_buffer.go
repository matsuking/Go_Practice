
import (
	"fmt"
	"runtime"
)

func channel_unbuffer_buffer() {
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()

	// // chに書き込まれるのを待っている状態
	// fmt.Println(<-ch)
	// wg.Wait()

	ch1 := make(chan int)

	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	// 第二引数でbufferの設定
	ch2 := make(chan int, 1)
	ch2 <- 2
	ch2 <- 3
	fmt.Println(<-ch2)
}
