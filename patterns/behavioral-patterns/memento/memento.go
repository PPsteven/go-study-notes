package main

import "fmt"

type Originator struct {
	money     int
	gameLevel int
}

func (e *Originator) createMemento() *Memento {
	return &Memento{money: e.money, gameLevel: e.gameLevel}
}

func (e *Originator) restoreMemento(m *Memento) {
	e.money = m.money
	e.gameLevel = m.gameLevel
}

func (e *Originator) String() string {
	return fmt.Sprintf("now game level is %d and you have $%d", e.gameLevel, e.money)
}

type Memento struct {
	money     int
	gameLevel int
}

type Saver struct {
	mementoArray []*Memento
}

func (c *Saver) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Saver) getMemento(index int) *Memento {
	return c.mementoArray[index]
}

func main() {
	saver := &Saver{}
	originator := &Originator{money: 100, gameLevel:10}
	fmt.Println(originator)
	saver.addMemento(originator.createMemento())
	fmt.Println("after serveral play, your money decrease...")
	originator.money = -30
	originator.gameLevel += 3
	fmt.Println(originator)
	fmt.Println("restore memento and play again")
	m := saver.getMemento(0)
	originator.restoreMemento(m)
	fmt.Println(originator)
}
