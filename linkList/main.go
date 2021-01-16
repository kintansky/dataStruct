package main

import "fmt"

type Node struct {
	Num  int
	Name string
	Next *Node
}

func NewNode(num int, name string) *Node {
	return &Node{
		Num:  num,
		Name: name,
		Next: nil,
	}
}

func (n *Node) ShowNode() {
	fmt.Printf("[%d:%s]", n.Num, n.Name)
}

// Link 单向链表
type Link struct {
	Head *Node // Head.Next指向第一个节点,Head只是一个没有数据的辅助节点
}

func NewLink() *Link {
	return &Link{
		Head: NewNode(-1, ""),
	}
}

// 直接尾插
func (l *Link) Add(n *Node) {
	tmp := l.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = n
}

// 考虑顺序的尾插
func (l *Link) AddByOrder(n *Node) {
	tmp := l.Head
	// 在当前位置跟后一个数据比较，如果符合退出循环修改Next指针
	for ; tmp.Next != nil; tmp = tmp.Next {
		if tmp.Next.Num == n.Num {
			fmt.Printf("Error: Num=%d already exits!\n", n.Num)
			return
		}
		if tmp.Next.Num > n.Num {
			break
		}
	}
	n.Next = tmp.Next
	tmp.Next = n
}

func (l *Link) ShowLink() {
	if l.IsEmpty() {
		return
	}
	// 打印的第一个节点应该是Head.Next
	for tmp := l.Head.Next; tmp != nil; tmp = tmp.Next {
		fmt.Printf("->")
		tmp.ShowNode()
	}
	fmt.Println()
}

func (l *Link) IsEmpty() bool {
	if l.Head.Next == nil {
		return true
	}
	return false
}

func (l *Link) Update(n *Node) {
	if l.IsEmpty() {
		fmt.Println("empty link")
		return
	}
	tmp := l.Head.Next
	found := false
	for ; tmp != nil; tmp = tmp.Next {
		if tmp.Num == n.Num {
			tmp.Name = n.Name
			found = true
			break
		}
	}
	if !found {
		fmt.Println("no such Node")
	}
}

func (l *Link) Del(num int) {
	tmp := l.Head
	// 和下一个节点进行比较，所以不用当游标tmp移到倒数第二个节点时，已经在比较最后一个节点，所以最后一个节点不用比较
	for ; tmp.Next != nil; tmp = tmp.Next {
		if tmp.Next.Num == num {
			tmp.Next = tmp.Next.Next
			break
		}
	}
}

func (l *Link) SizeOf() (n int) {
	tmp := l.Head
	for ; tmp.Next != nil; tmp = tmp.Next {
		n++
	}
	return
}

// 倒数第k个节点
func (l *Link) NoKFromTail(k int) (node *Node) {
	// 快慢指针
	tmp := l.Head
	node = l.Head
	for ; tmp.Next != nil; tmp = tmp.Next {
		if k-1 > 0 {
			k--
			continue
		}
		node = node.Next
	}
	// 如果没有这个节点的返回空
	if node == l.Head {
		node = nil
	}
	return
}

// 原地反转链表
func (l *Link) ReverseLink() {
	cur := l.Head.Next
	var (
		tmpHead  *Node
		nextNode *Node
	)
	for cur != nil {
		nextNode = cur.Next // 暂存下一个节点信息
		cur.Next = tmpHead
		tmpHead = cur
		cur = nextNode
	}
	l.Head.Next = tmpHead
}

// 反向打印
func (l *Link) ReverseShowLink(node *Node) {
	if node == nil {
		return
	}
	l.ReverseShowLink(node.Next)
	fmt.Printf("<-[%d:%s]", node.Num, node.Name)
}

// 合并l1、l2，譬如l1,l2都是正序Link
func CombineOrderLink(l1, l2 *Link, positiveOrder bool) (newLink *Link) {
	newLink = NewLink()
	curNewLink := newLink.Head
	curL1 := l1.Head.Next
	curL2 := l2.Head.Next
	for curL1 != nil && curL2 != nil {
		// 根据输入的正反序进行指向
		if (curL1.Num <= curL2.Num && positiveOrder) || (curL1.Num >= curL2.Num && !positiveOrder) {
			curNewLink.Next = curL1     // 不能使用add方法添加，否则会每次都遍历已添加的节点
			if curL1.Num == curL2.Num { // 相等的情况另一link也继续移到下一个节点
				curL2 = curL2.Next
			}
			curL1 = curL1.Next
		} else {
			curNewLink.Next = curL2
			curL2 = curL2.Next
		}
		curNewLink = curNewLink.Next
	}
	// 剩下的节点添加回去
	if curL1 != nil {
		curNewLink.Next = curL1
	} else if curL2 != nil { // curl2 也得判断nil
		curNewLink.Next = curL2
	}
	return
}

func main() {
	l := NewLink()
	fmt.Println(l.SizeOf())
	l.AddByOrder(NewNode(0, "user0"))
	l.ShowLink()
	fmt.Println(l.SizeOf())
	for i := 1; i < 5; i++ {
		l.Add(NewNode(i*3, fmt.Sprintf("user%d", i*3)))
	}
	l.ShowLink()
	fmt.Println(l.SizeOf())

	l.AddByOrder(NewNode(5, "user5"))
	l.AddByOrder(NewNode(1, "user1"))
	l.AddByOrder(NewNode(13, "user13"))
	l.ShowLink()

	l.Update(NewNode(1, "user_mod_1"))
	l.ShowLink()
	l.Update(NewNode(0, "user_mod_0"))
	l.ShowLink()
	l.Update(NewNode(13, "user_mod_13"))
	l.ShowLink()
	l.Update(NewNode(14, "user_mod_14"))
	l.ShowLink()

	l.Del(14)
	l.ShowLink()
	l.Del(5)
	l.ShowLink()
	l.Del(0)
	l.ShowLink()
	fmt.Println(l.NoKFromTail(1))
	fmt.Println(l.NoKFromTail(2))
	fmt.Println(l.NoKFromTail(6))
	fmt.Println(l.NoKFromTail(8))
	// l.NoKFromTail(1)

	l.ReverseLink()
	l.ShowLink()
	l.ReverseShowLink(l.Head.Next)
	fmt.Println()

	l1 := NewLink()
	l2 := NewLink()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			l2.AddByOrder(NewNode(i, fmt.Sprintf("node_%d", i)))
			continue
		}
		l1.AddByOrder(NewNode(i, fmt.Sprintf("node_%d", i)))
	}
	l1.AddByOrder(NewNode(20, "node_20"))
	for i := 12; i < 15; i++ {
		l2.AddByOrder(NewNode(i, fmt.Sprintf("node_%d", i)))
	}
	l1.ShowLink()
	l2.ShowLink()
	cl := CombineOrderLink(l1, l2, true)
	cl.ShowLink()

	l1.ReverseLink()
	l2.ReverseLink()
	cl = CombineOrderLink(l1, l2, false)
	cl.ShowLink()
}
