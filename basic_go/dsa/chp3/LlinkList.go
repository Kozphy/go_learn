package chp3

import "fmt"

// Node class
type Node struct {
	property int
	nextNode *Node
}

// LinkList class
type LinkedList struct {
	headNode *Node
}

// AddToHead method
func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = node
}

// IterateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func Exec_linkList() {
	var linkedList LinkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	fmt.Println(linkedList.headNode.property)
	linkedList.IterateList()
}
