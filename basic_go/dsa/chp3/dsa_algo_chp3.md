# Linear Data Structures

## We will cover the following linear data structures in this chapter

- Lists
- Sets
- Tuples
- Stacks

## List

A list is a `collection of ordered elements` that are used to store list of items. Unlike array lists, these `can expand and shrink dynamically`.

## LinkList

`LinkedList` is a sequence of nodes that have properties and a reference to the next node in the sequence. It is a linear data structure that is used to store data.

They are `not stored contiguously in memory`, which makes them different arrays.

### Node class

```go
// Node class
type Node struct {
    property int
    nextNode *Node
}
```

### The LinkList class

```go
// LinkedList class
type LinkedList struct {
    headNode *Node
}
```

### AddToHead method

The `AddToHead` method adds the node to the start of the linked list.

```go
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
```
