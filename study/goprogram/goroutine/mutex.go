package main

import (
	"fmt"
	"sync"
)

type Mutex struct {
	Done chan struct{}
}

func (m Mutex) Lock() {
	m.Done <- struct{}{}
}
func (m Mutex) Unlock() {
	<-m.Done
}

var (
	m = Mutex{Done: make(chan struct{}, 1)}
	//sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	m.Lock()
	defer m.Unlock()
	balance += amount
}

func Balance() int {
	m.Lock()
	defer m.Unlock()
	return balance
}

func Print() {
	fmt.Println("hello world!\n")
}

func main() {
	Deposit(50)
	fmt.Println(Balance(), 1e6)

	var once sync.Once
	once.Do(Print)
	once.Do(Print)
}
