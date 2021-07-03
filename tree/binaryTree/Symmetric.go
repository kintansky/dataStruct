package main

func isSymmetric(root *Node) bool {
	return dfsSymmetric(root, root)
}

func dfsSymmetric(leftTreeRoot, rightTreeRoot *Node) (b bool) {
	if leftTreeRoot == nil && rightTreeRoot == nil {
		b = true
		return
	}
	if leftTreeRoot != nil && rightTreeRoot != nil && leftTreeRoot.num == rightTreeRoot.num {
		b = true
	} else {
		b = false
		return
	}
	symmetricOuter := dfsSymmetric(leftTreeRoot.left, rightTreeRoot.right) // （外侧）左子树的左节点与右子树的右节点一致
	symmetricInner := dfsSymmetric(leftTreeRoot.right, rightTreeRoot.left) // （内侧）左子树的右节点与右子树的左节点一致
	return symmetricOuter && symmetricInner
}
