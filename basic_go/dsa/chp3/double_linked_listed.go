package chp3

// Node
type DoubleNode struct {
	property     int
	nextNode     *DoubleNode
	previousNode *DoubleNode
}

// Class
type DoubleLinkedList struct {
	headNode *DoubleNode
}

func (D_linkedList *DoubleLinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *DoubleNode {
	var node *DoubleNode
	var nodeWith *DoubleNode
	for node = D_linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}
