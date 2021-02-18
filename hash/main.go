package main

import "fmt"

type Node struct {
	id   int
	name string
	next *Node
}

func NewNode(id int, name string) *Node {
	return &Node{id: id, name: name}
}

type Link struct {
	head *Node
}

func NewLink() *Link {
	return &Link{head: NewNode(-1, "")}
}

func (l *Link) Add(node *Node) {
	tmp := l.head
	for tmp.next != nil {
		if tmp.next.id == node.id {
			fmt.Println("node already exists")
			return
		}
		if tmp.next.id > node.id {
			break
		}
		tmp = tmp.next
	}
	node.next = tmp.next
	tmp.next = node
}

func (l *Link) Del(node *Node) {
	tmp := l.head
	for tmp.next != nil {
		if tmp.next.id == node.id {
			tmp.next = tmp.next.next
			return
		}
		tmp = tmp.next
	}
	fmt.Printf("no node%#v\n", node)
}

func (l *Link) Show() {
	tmp := l.head.next
	for tmp != nil {
		fmt.Printf("[%d:%s]->", tmp.id, tmp.name)
		tmp = tmp.next
	}
	fmt.Println()
}

type HashTable struct {
	size    int
	hashArr []*Link
}

func NewHashTable(size int) *HashTable {
	hashArr := make([]*Link, size)
	for i := range hashArr {
		hashArr[i] = NewLink()
	}
	return &HashTable{
		size:    size,
		hashArr: hashArr,
	}
}

func (h *HashTable) Hash(id int) (idx int) {
	idx = id % h.size
	return
}

func (h *HashTable) Add(node *Node) {
	idx := h.Hash(node.id)
	h.hashArr[idx].Add(node)
}

func (h *HashTable) Show() {
	for i, link := range h.hashArr {
		fmt.Printf("%d:", i)
		link.Show()
	}
}

func (h *HashTable) Del(node *Node) {
	idx := h.Hash(node.id)
	h.hashArr[idx].Del(node)
}

func main() {
	hashTab := NewHashTable(5)
	for i := 0; i < 14; i++ {
		hashTab.Add(NewNode(i, fmt.Sprintf("node_%d", i)))
	}
	hashTab.Show()
	fmt.Println("del node 5")
	hashTab.Del(NewNode(5, "node_5"))
	hashTab.Show()

	fmt.Println("del node 4")
	hashTab.Del(NewNode(4, "node_4"))
	hashTab.Show()

	fmt.Println("del node 9")
	hashTab.Del(NewNode(9, "node_9"))
	hashTab.Show()

	fmt.Println("del node 20")
	hashTab.Del(NewNode(20, "node_20"))
	hashTab.Show()

	hashTab.Add(NewNode(5, "node_5"))
	hashTab.Add(NewNode(9, "node_9"))
	hashTab.Add(NewNode(4, "node_4"))
	hashTab.Add(NewNode(20, "node_20"))
	hashTab.Show()

}
