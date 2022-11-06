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

// return the last Node
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node = linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// traversed and checked to see whether the "property" value is equal to parameter.
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// adds a node with nodeProperty after node with property
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWith *Node = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
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
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	// fmt.Println(linkedList.headNode.property)
	linkedList.IterateList()
}
