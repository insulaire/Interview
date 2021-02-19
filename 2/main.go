package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Client struct {
	Id int
}

func (this *Client) Run() {
	fmt.Println(this.Id)
}

type Balancer struct {
	client []*Client
	rand   *rand.Rand
	l      uint64
	c      uint64
}

func NewBalancer() *Balancer {
	return &Balancer{
		client: []*Client{},
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
func (this *Balancer) AddClient(client *Client) {
	this.client = append(this.client, client)
	atomic.AddUint64(&this.l, 1)
}

func (this *Balancer) GetRandClient() *Client {
	if len(this.client) == 0 {
		return nil
	}
	this.rand.Seed(time.Now().UnixNano())
	//fmt.Println(n)
	return this.client[rand.Intn(len(this.client))]
}

func (this *Balancer) GetPollClient() *Client {
	if len(this.client) == 0 {
		return nil
	}

	if atomic.CompareAndSwapUint64(&this.c, this.l-1, 0) {
		return this.client[0]
	} else {
		n := atomic.AddUint64(&this.c, 1)
		return this.client[n]
	}
}

func main() {
	b := NewBalancer()
	b.AddClient(&Client{Id: 1})
	b.AddClient(&Client{Id: 2})
	b.AddClient(&Client{Id: 3})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//c := b.GetPollClient()
			c := b.GetRandClient()
			c.Run()
		}()
	}

	wg.Wait()
}
