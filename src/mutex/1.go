package mutex

import (
	"fmt"
	"runtime"
	"sync"
)

//variable compartida
var counter = 0

func Run1() {
	const gs = 100
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go counterFunc(&wg, &m)
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("count:", counter)
}

func counterFunc(wg *sync.WaitGroup, m *sync.Mutex) {
	defer func() {
		m.Unlock()
		wg.Done()
	}()

	m.Lock()
	v := counter
	v++
	counter = v
}
