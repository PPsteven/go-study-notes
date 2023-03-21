package main

import "fmt"

type Item interface {
	String() string
}

type Iterator interface {
	hasNext() bool
	getNext() *Item
}

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() Item {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

func NewUserIterator(users ...*User) *UserIterator {
	return &UserIterator{users: users}
}

type User struct {
	name string
}

func (u *User) String() string {
	return fmt.Sprintf("User is %s", u.name)
}

func main() {
	user1 := &User{"a"}
	user2 := &User{"b"}

	iter := NewUserIterator(user1, user2)
	for iter.hasNext() {
		fmt.Println(iter.getNext())
	}
}

// User is a
// User is b
