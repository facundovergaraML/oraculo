package racecondition

import (
	"fmt"
	"runtime"
	"sync"
)

//variable compartida
var counter = 0

//cantidad de gorutines
const gs = 100

func Run1() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())

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
	v := counter
	runtime.Gosched()
	v++
	counter = v
	wg.Done()
}
