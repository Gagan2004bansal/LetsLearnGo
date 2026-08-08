package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id, "is running")
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}
