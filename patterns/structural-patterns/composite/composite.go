// 同样还是实现的是文件夹
// 不同之处是为了代码实现方便，统一了 File 和 Folder 的结构。为此 File 实现了add(Inode)方法，不进行任何操作。
package main

import "fmt"

type Inode interface {
	add(Inode)
	print(int)
}

type nodeType uint

const (
	fileType nodeType = iota
	folderType
)

func intent(level int, t nodeType) string {
	s := ""
	for i := 0; i <= level - 1; i++ {
		s = s + "  "
	}
	if t == folderType {
		s = s + "->"
	} else {
		s = s + "  "
	}
	return s
}

type File struct {
	name string // file name
}

func (f *File) print(level int) {
	fmt.Println(intent(level, fileType)+f.name)
}

func (f *File) add(_ Inode) {}

type Folder struct {
	name string  // folder name
	inodes []Inode
}

func (f *Folder) print(level int) {
	fmt.Println(intent(level, folderType)+f.name)
	for _, indoe := range f.inodes {
		indoe.print(level+1)
	}
}

func (f *Folder) add(inode Inode) {
	f.inodes = append(f.inodes, inode)
}

func NewFile(name string) *File{
	return &File{name: name}
}

func NewFolder(name string) *Folder{
	return &Folder{name: name}
}

func main() {
	docFolder := NewFolder("document")
	goFolder := NewFolder("Go")
	goFolder.add(NewFile("Best Practice in Go.pdf"))
	goFolder.add(NewFile("Programming in Go: Creating Applications for the 21st Century.epub"))

	designPatternFolder := NewFolder("Design Pattern")
	designPatternFolder.add(NewFile("Design Pattern.pdf"))

	docFolder.add(goFolder)
	docFolder.add(designPatternFolder)

	docFolder.print(0)
}

// ->document
//  ->Go
//      Best Practice in Go.pdf
//      Programming in Go: Creating Applications for the 21st Century.epub
//  ->Design Pattern
//      Design Pattern.pdf
