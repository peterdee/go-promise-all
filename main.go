package main

import (
	"fmt"
	"sync"
	"time"
)

var array = []int32{1, 2, 3, 4}
var group sync.WaitGroup

func handler(value int32) {
	defer group.Done()
	time.Sleep(time.Second * 5)
	fmt.Println("done working with", value)
}

func main() {
	start := time.Now()
	for _, value := range array {
		group.Add(1)
		go handler(value)
	}
	group.Wait()

	fmt.Println("all done in", time.Since(start))
}
