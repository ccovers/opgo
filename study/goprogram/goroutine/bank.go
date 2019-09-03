package main

import (
	"fmt"
	"time"
)

type Withdrawal struct {
	Amount int
	Done   chan<- bool
}

var draw = make(chan Withdrawal)
var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case drawal := <-draw:
			if drawal.Amount <= balance {
				balance -= drawal.Amount
				drawal.Done <- true
			} else {
				drawal.Done <- false
			}
		}
	}
}

func withDraw(amount int) bool {
	done := make(chan bool)
	draw <- Withdrawal{
		Amount: amount,
		Done:   done,
	}
	br := <-done
	return br
}

func main() {
	go teller()
	go Deposit(49)

	time.Sleep(1 * time.Second)

	fmt.Println(Balance())
	r := withDraw(50)
	if r {
		fmt.Println("取款成功")
	} else {
		fmt.Println("取款失败")
	}
}
