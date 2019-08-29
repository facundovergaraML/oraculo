package introduction

import (
	"fmt"
)

func Run1() {
	numbers1()
	alphabets1()
}

func numbers1() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
}

func alphabets1() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
	}
}
