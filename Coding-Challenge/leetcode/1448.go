/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, v int) int {
	if node == nil {
		return 0
	}
	curMax := max(node.Val, v)
	if node.Val >= v {
		return 1 + dfs(node.Left, curMax) + dfs(node.Right, curMax)
	}
	return dfs(node.Left, curMax) + dfs(node.Right, curMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}