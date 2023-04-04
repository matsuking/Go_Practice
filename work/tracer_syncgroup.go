

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	// 無名関数の前にキーワードをつけることで、goroutineとして扱うことができる。
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("goroutine invoked.")
// 	}()
// 	wg.Wait()
// 	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
// 	fmt.Println("main func finish")
// }

func tracer_syncgroup() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		// f.Close()をerrに代入して、errがnilではない場合の条件式
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()

	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU cores:", runtime.NumCPU())

	// task(ctx, "Task1")
	// task(ctx, "Task2")
	// task(ctx, "Task3")

	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "Task1")
	go cTask(ctx, &wg, "Task2")
	go cTask(ctx, &wg, "Task3")
	wg.Wait()

	s := []int{1, 2, 3}
	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main func finish")
}

// 同期処理
func task(ctx context.Context, name string) {
	// startRegionでチェーンつなぎにすると、最後のEnd()は遅延する。
	defer trace.StartRegion(ctx, name).End()
	time.Sleep(time.Second)
	fmt.Println(name)
}

// 非同期処理
func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}
