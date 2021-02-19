package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
循环打印dog fish cat 100次
*/
func main() {
	wg := sync.WaitGroup{}
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	catChan := make(chan struct{}, 1)
	wg.Add(3)
	dogChan <- struct{}{}
	go run(&wg, dogChan, fishChan, func() {
		fmt.Println("dog")
	})

	go run(&wg, fishChan, catChan, func() {
		fmt.Println("fish")
	})

	go run(&wg, catChan, dogChan, func() {
		fmt.Println("cat")
	})
	wg.Wait()
}

func run(wg *sync.WaitGroup, curChan, nextChan chan struct{}, f func()) {
	defer wg.Done()
	var Count uint64
	for {
		if Count > 100 {
			return
		}
		<-curChan
		f()
		atomic.AddUint64(&Count, 1)
		nextChan <- struct{}{}
	}
}
