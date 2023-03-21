package main

import "fmt"

type Editor struct {}

func (e *Editor) copy() {
	fmt.Println("copy text")
}

func (e *Editor) paste() {
	fmt.Println("paste text")
}

type Command interface {
	execute()
	// undo()
}
type CopyCommand struct {
	recevier *Editor
}

func (c *CopyCommand) execute() {
	c.recevier.copy()
}

type PasteCommand struct {
	recevier *Editor
}

func (c *PasteCommand) execute() {
	c.recevier.paste()
}

func main() {
	editor := &Editor{}
	copyCmd := &CopyCommand{recevier: editor}
	pasteCmd := &PasteCommand{recevier: editor}
	copyCmd.execute()
	pasteCmd.execute()
}
