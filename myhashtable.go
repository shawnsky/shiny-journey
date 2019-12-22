package main

import (
	"crypto/md5"
	"fmt"
)
type Entity struct {
	Key string
	Value string
}

type Node struct {
	Data Entity
	Next *Node
}

type Linklist struct {
	head *Node
	current *Node
}

func (list *Linklist) Append (entity Entity) {
	node := Node { Data: entity, Next: nil}
	if list.head == nil {
		list.head = &node
		list.current = list.head
	} else {
		list.current.Next = &node
		list.current = list.current.Next
	}
}

func (list *Linklist) Print () {
	p := list.head
	for p != nil {
		fmt.Printf("%s->", p.Data.Key)
		p = p.Next
	}
}

type HashTable struct {
	Table [16]Linklist
}

func GetHashCode(key string) (code int) {
	hash := md5.Sum([]byte(key))
	for _, v := range hash {
		code += int(v)
	}
	return
}

func main() {
	var list Linklist
	data1 := Entity{
		Key:   "shawn",
		Value: "12345",
	}
	data2 := Entity{
		Key:   "coin",
		Value: "67890",
	}
	data3 := Entity{
		Key:   "apple",
		Value: "55667",
	}

	list.Append(data1)
	list.Append(data2)
	list.Append(data3)

	list.Print()

}