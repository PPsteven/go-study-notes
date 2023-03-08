package main

import (
	"fmt"
	"github.com/goombaio/namegenerator"
	"sync"
)

var nameGenerator = namegenerator.NewNameGenerator(2023)

var enemyPool = sync.Pool{
	New: func() interface{} {
		return new(Enemy)
	},
}

type Enemy struct {
	Name string
}

func NewEnemy() *Enemy {
	e := enemyPool.Get().(*Enemy)
	e.Name = nameGenerator.Generate()
	fmt.Printf("new an enemy: %s(addr: %p)\n", e.Name, e)
	return e
}

func (e *Enemy) Killed() {
	enemyPool.Put(e)
}

func main() {
	e1 := NewEnemy()
	e1.Killed()
	e2 := NewEnemy()
	e3 := NewEnemy()
	e2.Killed()
	e3.Killed()
}
// Output
//new an enemy: nameless-waterfall(addr: 0xc000098220)
//new an enemy: throbbing-fog(addr: 0xc000098220)
//new an enemy: dry-paper(addr: 0xc0000982a0)
