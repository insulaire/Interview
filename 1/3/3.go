package main

import "sync"

func main() {

	answer1()
	answer2()
}

var res []int
var mutex sync.Mutex

//Golang中除了加Mutex锁以外还有哪些方式安全读写共享变量？

//Mutex 方式
func answer1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(item int) {
			mutex.Lock()
			res = append(res, item)
			wg.Done()
			defer mutex.Unlock()
		}(i)
	}
	wg.Wait()

}

//Channel 方式
var c = make(chan interface{}, 1)

func answer2() {
	wg := sync.WaitGroup{}
	c <- "lock"
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(item int) {
			<-c
			res = append(res, item)
			wg.Done()
			c <- "lock"
		}(i)
	}
	wg.Wait()
	defer close(c)
}
