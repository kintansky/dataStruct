package main

func maxDepth(root *Node) (maxDeep int) {
	dfs(root, 0, &maxDeep)
	return
}

func dfs(node *Node, subTreeDeep int, maxDeep *int) {
	if node == nil { // 如果当前节点空，返回即可
		return
	}
	subTreeDeep++               // 非空，更新当前子树高度为+1
	if subTreeDeep > *maxDeep { // 更新树的最大高度
		*maxDeep = subTreeDeep
	}
	// 以当前节点为根，继续遍历其子树高度
	dfs(node.left, subTreeDeep, maxDeep)
	dfs(node.right, subTreeDeep, maxDeep)
}

func maxDepth2(root *Node) (maxDeep int) {
	maxDeep = findDepth(root)
	return
}

func findDepth(node *Node) (maxDeep int) {
	if node == nil {
		return
	}
	leftDeep := findDepth(node.left)
	rightDeep := findDepth(node.right)
	if leftDeep > rightDeep {
		maxDeep = leftDeep + 1
	} else {
		maxDeep = rightDeep + 1
	}
	return
}
