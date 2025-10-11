package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(wg *sync.WaitGroup, massage string) {
	defer wg.Done()
	fmt.Println(massage)
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)

		wg.Add(1)
		go doPrint(&wg, data)
	}
	wg.Wait()
}
