package week1

import (
	"fmt"
	"sync"
)

// space : O(n)
// insert, search, delete: O(log n)
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(key int, value int, useKey bool) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	var treeNode *TreeNode = &TreeNode{key, value, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode, useKey)
	}
}

func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode, useKey bool) {
	// use value to order
	if useKey == false {
		// insert to left side
		if newTreeNode.value < rootNode.value {
			if rootNode.leftNode == nil {
				rootNode.leftNode = newTreeNode
			} else {
				insertTreeNode(rootNode.leftNode, newTreeNode, useKey)
			}
		} else {
			// insert to right side
			if rootNode.rightNode == nil {
				rootNode.rightNode = newTreeNode
			} else {
				insertTreeNode(rootNode.rightNode, newTreeNode, useKey)
			}
		}
	}

	// use key to order
	if useKey == true {
		if newTreeNode.key < rootNode.key {
			if rootNode.leftNode == nil {
				rootNode.leftNode = newTreeNode
			} else {
				insertTreeNode(rootNode.leftNode, newTreeNode, useKey)
			}
		} else {
			if rootNode.rightNode == nil {
				rootNode.rightNode = newTreeNode
			} else {
				insertTreeNode(rootNode.rightNode, newTreeNode, useKey)
			}
		}
	}
}

func printTreeNode(treeNode *TreeNode) {
	fmt.Println("printTreeNode")
	if treeNode != nil {
		fmt.Println("Value: ", treeNode.value)
		fmt.Printf("TreeNode Left")
		printTreeNode(treeNode.leftNode)
		fmt.Printf("TreeNode Right")
		printTreeNode(treeNode.rightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

// String method
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("************************************************")
	stringify(tree.rootNode, 0)
	fmt.Println("************************************************")
}

// stringify method
func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "***> "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}
