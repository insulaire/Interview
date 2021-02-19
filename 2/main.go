package main

import (
	"fmt"
	"math/rand"
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
}

func NewBalancer() *Balancer {
	return &Balancer{
		client: []*Client{},
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
func (this *Balancer) AddClient(client *Client) {
	this.client = append(this.client, client)
}

func (this *Balancer) GetClient() *Client {
	if len(this.client) == 0 {
		return nil
	}
	this.rand.Seed(time.Now().UnixNano())
	//fmt.Println(n)
	return this.client[rand.Intn(len(this.client))]
}

func main() {
	b := NewBalancer()
	b.AddClient(&Client{Id: 1})
	b.AddClient(&Client{Id: 2})
	b.AddClient(&Client{Id: 3})
	for i := 0; i < 10; i++ {
		c := b.GetClient()
		c.Run()
	}
}
