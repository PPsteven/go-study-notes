package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

// Store is Product, all kind of Product do the same thing
// redis, mysql both have Open() method
// car, plane both have Transport() method
type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	MemoryStorage
)

func NewStore(t StorageType) (Store, error) {
	switch t {
	case MemoryStorage:
		return newMemoryStorage( /*...*/ ), nil
	case DiskStorage:
		return newDiskStorage( /*...*/ ), nil
	default:
		return nil, errors.New("invalid storage")
	}
}

var memoryStore map[string]string

// MemoryStore is Concrete Product A
type MemoryStore struct {
	buf  *bytes.Buffer
	name string
	l    sync.Mutex
}

func (m *MemoryStore) Open(name string) (io.ReadWriteCloser, error) {
	m.l.Lock()
	defer m.l.Unlock()
	if memoryStore == nil {
		memoryStore = make(map[string]string)
	}
	m.name = name
	text, _ := memoryStore[name]
	m.buf = bytes.NewBufferString(text)
	return m, nil
}

func (m *MemoryStore) Read(p []byte) (n int, err error) {
	return m.buf.Read(p)
}

func (m *MemoryStore) Write(p []byte) (n int, err error) {
	return m.buf.Write(p)
}

func (m *MemoryStore) Close() error {
	m.l.Lock()
	defer m.l.Unlock()
	memoryStore[m.name] = m.buf.String()
	m.buf.Reset()
	return nil
}

// Concrete Creator
func newMemoryStorage() *MemoryStore{
	return &MemoryStore{}
}

// DiskStore is Concrete Product B
type DiskStore struct {}

func (d *DiskStore) Open(name string) (io.ReadWriteCloser, error) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Concrete Creator
func newDiskStorage() *DiskStore {
	return &DiskStore{}
}

func main() {
	s, _ := NewStore(MemoryStorage)
	f, _ := s.Open("file")
	_, _ = f.Write([]byte("data"))
	_ = f.Close()

	p := make([]byte, 4)
	f, _ = s.Open("file")
	_, _ = f.Read(p)
	fmt.Println(string(p)) // data
}