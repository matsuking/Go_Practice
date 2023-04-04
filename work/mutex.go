import (
	"fmt"
	"time"

	"sync"
)

func mutex() {
	var wg sync.WaitGroup
	var rwMu sync.RWMutex
	var c int

	wg.Add(4)
	go write(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)

	wg.Wait()
	fmt.Println("finish")
	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// var i int

	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	// 操作中は他のところから操作できないようにする
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	// i++
	// 	i = 1
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	// i++
	// 	i = 2
	// }()
	// wg.Wait()
	// fmt.Println(i)
}

func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("read lock")
	fmt.Println(*c)
	time.Sleep(1 * time.Second)
	fmt.Println("read unlock")

}

func write(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	mu.RLock()
	defer mu.Unlock()
	fmt.Println("write lock")
	*c += 1
	time.Sleep(1 * time.Second)
	fmt.Println("write unlock")

}
