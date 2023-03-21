package main

import "fmt"

type State interface {
	On(m *Machine)
	Off(m *Machine)
}

type Machine struct {
	current State
}

func (m *Machine) setCurrent(s State) {
	m.current = s
}

func (m *Machine) On() {
	m.current.On(m)
}

func (m *Machine) Off() {
	m.current.Off(m)
}

func NewMachine() *Machine {
	fmt.Printf("Machine is ready.\n")
	return &Machine{current: &Off{}}
}

type On struct {}

func (o *On) On(m *Machine) {
	fmt.Printf("already on.\n")
}

func (o *On) Off(m *Machine) {
	fmt.Printf("on to off\n")
	m.setCurrent(&Off{})
}

type Off struct {}

func (o *Off) On(m *Machine) {
	fmt.Printf("off to on\n")
	m.setCurrent(&On{})
}

func (o *Off) Off(m *Machine) {
	fmt.Printf("already off.\n")
}

func main() {
	m := NewMachine()
	m.On()
	m.On()
	m.Off()
}