package main

import "fmt"

type Subject interface {
	register(o Observer)
	notifyAll()
}

type Observer interface {
	notify()
}

type Product struct {
	inStock chan struct{}
	obersvers []Observer
	notify  chan struct{}
}

func NewProduct() *Product {
	p := &Product{inStock: make(chan struct{}, 1), notify: make(chan struct{})}
	go func() {
		<- p.inStock
		p.notifyAll()
	}()
	return p
}

func (p *Product) register(o Observer) {
	p.obersvers = append(p.obersvers, o)
}

func (p *Product) notifyAll() {
	for _, o := range p.obersvers {
		o.notify()
	}
	defer close(p.notify)
}

type Customer struct {
	email string
}

func (c *Customer) notify() {
	fmt.Printf("send email to customer email: %s\n", c.email)
}

func productInStock(p *Product) {
	p.inStock <- struct{}{}
}


func main() {
	c1 := &Customer{email: "xyz@a.com"}
	c2 := &Customer{email: "abc@b.com"}
	product := NewProduct()
	product.register(c1)
	product.register(c2)
	productInStock(product)

	<- product.notify
}

//send email to customer email: xyz@a.com
//send email to customer email: abc@b.com