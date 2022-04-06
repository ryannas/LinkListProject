package main

import "LinkListProject/LinkList"

func main() {
	list := LinkList.InitList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.Insert(3, 4)
	list.ShowList()
}
