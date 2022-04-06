package LinkList

import "fmt"

type Node struct {
	data int
	next *Node
}
type List struct {
	length   int
	head *Node
}

func InitList() *List {
	node := new(Node)
	L := new(List)
	L.head = node
	return L
}

func (list *List) IsNull() bool {
	if list.head == nil {
		return true
	}
	return false
}

func (list *List) Insert(index int, v int) {
	if index <= 0 || index > list.length {
		fmt.Println("err")
	} else {
		pre := list.head
		node := &Node{data: v}
		if index == 1 {
			node.next = pre
			list.head = node
		} else {
			for count := 1; count < index-1; count++ {
				pre = pre.next
			}
			node.next = pre.next
			pre.next = node
		}
		list.length--
	}
}

func (list *List) PushFront(v int) {
	node := &Node{data: v}
	if list.IsNull() {
		list.head = node
		list.length++
		return
	}
	node.next = list.head
	list.head = node
	list.length++
	return
}

func (list *List) PushBack(v int) {
	node := &Node{data: v}
	if list.IsNull() {
		list.head.next = node
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
	list.length++
	return
}

func (list *List) Delete(index int) {
	if index <= 0 || index > list.length {
		fmt.Println("err")
		return
	} else {
		pre := list.head
		if index == 1 {
			list.head = pre.next
		} else {
			pre := list.head
			for count := 1; count < index-1; count++ {
				pre = pre.next
			}
			pre.next = pre.next.next
		}
		list.length--
	}
}

func (list *List) Remove(v int) {
	pre := list.head
	if pre.data == v {
		list.head = pre.next
		fmt.Println("ok")
	} else {
		for pre.next != nil {
			if pre.next.data == v {
				pre.next = pre.next.next
				fmt.Println("ok")
				return
			} else {
				pre = pre.next
			}
		}
		fmt.Println("fail")
		return
	}
}

func (list *List) Locate(v int) bool {
	if list.IsNull() {
		fmt.Println("err")
		return false;
	}
	pre := list.head
	for pre != nil {
		if pre.data == v {
			return true
		}
		pre = pre.next
	}
	return false
}

func (list *List) Get(index int) int {
	if index <= 0 || index > list.length {
		fmt.Println("err")
		return -1
	}
	pre := list.head
	for j := 0; j < index; j++ {
		if pre != nil {
			pre = pre.next
		}
	}
	return pre.data
}

func (list *List) ShowList() {
	if !list.IsNull() {
		cur := list.head
		for cur.next != nil{
			fmt.Printf("%v -> ", cur.data)
			cur = cur.next
		}
		fmt.Println("nil")
	}
}
