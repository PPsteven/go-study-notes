package main

import (
	"bytes"
	"fmt"
	"html/template"
)

type Template interface {
	GetTemplate() string
}

type ChristmasCard struct {
	Name string
}

func (c *ChristmasCard) GetTemplate() string {
	return `Marry Christmas, {{.Name}}`
}

type BirthdayCard struct {
	Name string
}

func (c *BirthdayCard) GetTemplate() string {
	return `Happy Birthday, {{.Name}}`
}

type Context struct {
	temp Template
}

func (c *Context) SetTemplate(t Template) {
	c.temp = t
}

func (c *Context) Execute() {
	var buffer bytes.Buffer
	t := template.Must(template.New("").Parse(c.temp.GetTemplate()))
	_ = t.Execute(&buffer, c.temp)
	fmt.Println(buffer.String())
}

func main() {
	c := &Context{}
	cc := &ChristmasCard{"Jack"}
	c.SetTemplate(cc)
	c.Execute()
	bc := &BirthdayCard{"Mary"}
	c.SetTemplate(bc)
	c.Execute()
}
