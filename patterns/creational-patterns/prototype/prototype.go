// prototype 原型链模式
// 原型链模式需要注意的是，要完全Copy一个原型，不能仅仅拷贝指针而是值。
// 实现一个文件夹复制功能:
// 1. 实现两个类型：文件File 和 文件夹Folder
// 2. 类型均支持Copy操作
package main

import "fmt"

type Inode interface {
	print(intent int)
	clone() Inode
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

func splitFileName(path string) (string, string){
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[:i], path[i:]
		}
	}
	return path, ""
}

type File struct {
	name string // file name
}

func (f *File) print(level int) {
	fmt.Println(intent(level, fileType)+f.name)
}

func (f *File) clone() Inode {
	name, ext := splitFileName(f.name)
	return &File{name: name+"_copy"+ext}
}

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

func (f *Folder) clone() Inode {
	inodes := make([]Inode, len(f.inodes))
	for i, inode := range f.inodes {
		inodes[i] = inode.clone()
	}
	return &Folder{name: f.name+"_copy", inodes: inodes}
}

func main() {
	fd1 := &Folder{"pdf", []Inode{
		&Folder{"Go", []Inode{
			&File{"Best Practice in Go.pdf"},
			&File{"Programming in Go: Creating Applications for the 21st Century.epub"}}},
		&Folder{"Design Pattern", []Inode{
			&File{"Design Pattern.pdf"}}},
	}}

	fmt.Println("----- origin file ------")
	fd1.print(0)
	fd2 := fd1.clone()
	fmt.Println("----- copy file ------")
	fd2.print(0)
}

// ----- origin file ------
//->pdf
//  ->Go
//      Best Practice in Go.pdf
//      Programming in Go: Creating Applications for the 21st Century.epub
//  ->Design Pattern
//      Design Pattern.pdf
//----- copy file ------
//->pdf_copy
//  ->Go_copy
//      Best Practice in Go_copy.pdf
//      Programming in Go: Creating Applications for the 21st Century_copy.epub
//  ->Design Pattern_copy
//      Design Pattern_copy.pdf