package main

import "fmt"

type Plant struct {
	name string
	mediator Mediator
}

func (p *Plant) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Printf("Plant %s arrival blocked, waiting\n", p.name)
		return
	}
	fmt.Printf("Plant %s arrived\n", p.name)
}

func (p *Plant) depart() {
	fmt.Printf("Plant %s is leaving\n", p.name)
	p.mediator.notifyAboutDeparture()
}

// 中间人
type Mediator interface {
	canArrive(*Plant)	bool // 能否降落
	notifyAboutDeparture()   // 起飞
}

type ControlTower struct {
	isPlatformFree bool
}

func (t *ControlTower) canArrive(*Plant) bool {
	if t.isPlatformFree {
		t.isPlatformFree = false
		return true
	}
	return false
}

func (t *ControlTower) notifyAboutDeparture() {
	if !t.isPlatformFree {
		t.isPlatformFree = true
	}
}

func NewControlTower() *ControlTower{
	return &ControlTower{isPlatformFree: true}
}

func main() {
	mediator := NewControlTower()
	plant1 := &Plant{mediator: mediator, name: "A"}
	plant2 := &Plant{mediator: mediator, name: "B"}

	plant1.arrive()
	plant2.arrive()
	plant1.depart()
	plant2.arrive()
}

// Plant A arrived
// Plant B arrival blocked, waiting
// Plant A is leaving
// Plant B arrived