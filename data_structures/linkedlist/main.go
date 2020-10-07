package main

import "log"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		log.Printf("length:%d::data:%d ", l.length, toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

//TODO: handle removing duplicates that aren't head
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		l.deleteWithValue(value)
		return
	}

	previousToDelete := l.head
	for previousToDelete.next != nil {
		if previousToDelete.next.next == nil {
			return
		}

		if previousToDelete.next.data == value {
			previousToDelete.next = previousToDelete.next.next
			l.length--
		}
		previousToDelete = previousToDelete.next
	}
}

func main() {
	list := linkedList{}
	//for _, node := range []node{node{data:1}, node{data:1}, node{data:2}, node{data:3}, node{data:5}}{
	//log.Printf("prepending %d\n", node.data)
	//list.prepend(&node)
	//}
	list.prepend(&node{data: 1})
	list.prepend(&node{data: 1})
	list.prepend(&node{data: 2})
	list.prepend(&node{data: 3})
	list.prepend(&node{data: 5})
	log.Println(list)
	list.printListData()
	log.Println("deleting 8")
	list.deleteWithValue(8)
	list.printListData()
	log.Println("deleting 3")
	list.deleteWithValue(3)
	list.printListData()
	log.Println("deleting 1")
	list.deleteWithValue(1)
	list.printListData()
	log.Println(list)
}
