package main

import (
	"crypto/md5"
	"fmt"
)
type Entry struct {
	Key string
	Value string
}

type Node struct {
	Data Entry
	Next *Node
}

type Linklist struct {
	head *Node
	current *Node
}

func (list *Linklist) Append (entity Entry) {
	node := Node { Data: entity, Next: nil}
	if list.head == nil {
		list.head = &node
		list.current = list.head
	} else {
		list.current.Next = &node
		list.current = list.current.Next
	}
}

func (list *Linklist) Exist (key string) bool {
	p := list.head
	for p != nil {
		if p.Data.Key == key {
			return true
		}
		p = p.Next
	}
	return false
}

func (list *Linklist) Print () {
	p := list.head
	for p != nil {
		fmt.Printf("%s->", p.Data.Key)
		p = p.Next

	}
	fmt.Print("\n")
}

type HashTable struct {
	buckets [16]Linklist
}

func (table *HashTable) Put (key string, value string) {
	entry := Entry{key,value}
	pos := GetHashCode(key) % 16
	if table.buckets[pos].Exist(key) {
		return
	}
	table.buckets[pos].Append(entry)
}

func (table *HashTable) Get (key string) (value string) {
	pos := GetHashCode(key) % 16
	list := table.buckets[pos]
	p := list.head
	for p != nil {
		if p.Data.Key == key {
			value = p.Data.Value
			return
		}
	}
	return
}

func GetHashCode(key string) (code int) {
	hash := md5.Sum([]byte(key))
	for _, v := range hash {
		code += int(v)
	}
	return
}

func main() {
	var ht HashTable
	ht.Put("c", "xt123")
	ht.Put("t", "y456")
	ht.Put("r", "pj789")
	fmt.Println(ht.Get("r"))
}