package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct {
}

var instance *Singleton

func getInstance() *Singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &Singleton{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}
