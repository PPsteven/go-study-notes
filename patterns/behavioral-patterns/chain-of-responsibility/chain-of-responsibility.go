package main

import "fmt"

type Handler interface {
	Handle()
	SetNext(Handler)
}

type EndHandler struct {}

func (h *EndHandler) Handle() {
	fmt.Println("end...")
}

func (h *EndHandler) SetNext(Handler) {}

type handler struct {
	name string
	next Handler
}

func (h *handler) Handle() {
	fmt.Printf("%s start processing\n", h.name)
	h.next.Handle()
}

func (h *handler) SetNext(handler Handler) {
	h.next = handler
}

func NewHandler(name string) *handler {
	return &handler{name: name, next: &EndHandler{}}
}

type Chain struct {
	handlers []Handler
}

func (c *Chain) Load(handlers ...Handler) {
	c.handlers = handlers
	for i := 1; i < len(handlers); i++ {
		c.handlers[i-1].SetNext(c.handlers[i])
	}
}

func (c *Chain) Excute() {
	c.handlers[0].Handle()
}

func main() {
	h1 := NewHandler("h1")
	h2 := NewHandler("h2")
	h3 := NewHandler("h3")

	chain := &Chain{}
	chain.Load(h1, h2, h3)
	chain.Excute()
}

//h1 start processing
//h2 start processing
//h3 start processing
//end...