package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {
	go teller()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(100)
		fmt.Println("=", Balance())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(200)
		fmt.Println("=", Balance())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := Withdraw(200)
		if !res {
			fmt.Println("取款失败")
		}
	}()
	wg.Wait()
	b := Balance()
	fmt.Println(b)
}
