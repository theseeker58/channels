package main

import (
	"fmt"
	"sync"
	//"time"
)

type result struct {
	a int
	b int
}

func (r result) allValid() bool {
	ok := r.a == 0 && r.b == 0
	return ok
}

//var wg sync.WaitGroup

func main() {
	var res = new(result)
	var wg sync.WaitGroup
	wg.Add(2)
	result1 := make(chan int)
	result2 := make(chan int)
	go checkIban(&wg, result1)
	go checkAccount(&wg, result2)
	res.a = <-result1
	res.b = <-result2
	wg.Wait()
	fmt.Println(res)
	fmt.Println("Is valid ", res.allValid())
	fmt.Println("Both goroutines returned values")
}

func checkIban(wg *sync.WaitGroup, result chan int) {
	defer wg.Done()
	//time.Sleep(10 * time.Second)
	result <- 0
	fmt.Println("checkIban completed")
}

func checkAccount(wg *sync.WaitGroup, result chan int) {
	defer wg.Done()
	result <- 0
	fmt.Println("checkAccount completed")
}
