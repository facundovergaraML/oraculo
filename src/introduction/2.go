package introduction

import (
	"fmt"
)

func Run2() {
	go numbers2()
	go alphabets2()
}

func numbers2() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
}

func alphabets2() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
	}
}
