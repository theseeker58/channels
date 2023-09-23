package main

import (
	"fmt"
	"sync"
)

type outcome struct {
	app  string
	code int
	msg  string
}

type overallResult struct {
	a outcome
	b outcome
}

func (r overallResult) allValid() bool {
	ok := r.a.code == 0 && r.b.code == 0
	return ok
}

func main() {
	var res = new(overallResult)
	var wg sync.WaitGroup
	wg.Add(2)
	result := make(chan outcome)
	go checkIban(&wg, result)
	go checkAccount(&wg, result)
	res.a, res.b = <-result, <-result
	wg.Wait()
	close(result)
	fmt.Println(res)
	fmt.Println("Is valid", res.allValid())
	fmt.Println("Both goroutines returned values")
}

func checkIban(wg *sync.WaitGroup, result chan outcome) {
	defer wg.Done()
	returnValue := outcome{app: "a", code: 0}
	result <- returnValue
	fmt.Println("checkIban completed")
}

func checkAccount(wg *sync.WaitGroup, result chan outcome) {
	defer wg.Done()
	//returnValue := outcome{app: "b", code: 0}
	returnValue := outcome{app: "b", code: 1, msg: "Conto inesistente"}
	result <- returnValue
	fmt.Println("checkAccount completed")
}
