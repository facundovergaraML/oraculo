package atomic

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64
const gs = 100

func Run1() {
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go counterFunc(&wg)
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count:", counter)
}

func counterFunc(wg *sync.WaitGroup) {
	defer wg.Done()

	//recibe un puntero a la variable y un delta(cuanto se le quiere sumar a esa variable)
	atomic.AddInt64(&counter, 1)
	fmt.Println(atomic.LoadInt64(&counter))
}

