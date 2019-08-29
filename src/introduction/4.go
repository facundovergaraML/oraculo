package introduction

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func Run4() {
	wg.Add(1)
	printInfo()

	go numbers4()
	alphabets4()

	printInfo()
	wg.Wait()
}

func numbers4() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func alphabets4() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
	}
}

func printInfo(){
	fmt.Println("\n\nCPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

