package week1

func create_binary_search_tree(useKey bool) *BinarySearchTree {
	var tree = &BinarySearchTree{}
	datas := []int{4, 2, 7, 1, 3, 6, 9}
	for pos, data := range datas {
		tree.InsertElement(pos, data, useKey)
	}
	return tree
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.leftNode
	root.leftNode = root.rightNode
	root.rightNode = tmp

	invertTree(root.leftNode)
	invertTree(root.rightNode)
	return root
}

func Execute_invertTree() {
	// optionl parameter
	var useKey bool = false
	BTree := create_binary_search_tree(useKey)
	BTree.String()
	invertTree(BTree.rootNode)
	BTree.String()

}
