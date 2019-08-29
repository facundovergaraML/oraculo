package introduction

import (
	"fmt"
	"sync"
)

func Run3() {
	var wg sync.WaitGroup
	wg.Add(2)

	go numbers3(&wg)
	go alphabets3(&wg)

	wg.Wait()
}

func numbers3(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
}

func alphabets3(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
	}
}
