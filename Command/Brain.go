package main

import "fmt"

type Brain struct {
	isWorking bool
}

func (b *Brain) on() {
	b.isWorking = true
	fmt.Println("Turning brain on")
}

func (b *Brain) off() {
	b.isWorking = false
	fmt.Println("Turning brain off")
}
