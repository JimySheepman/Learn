/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	stackP, stackQ := make([]*TreeNode, 0), make([]*TreeNode, 0)
	stackP = append(stackP, p)
	stackQ = append(stackQ, q)

	for len(stackP) > 0 && len(stackQ) > 0 {
		tempP := stackP[0]
		stackP = stackP[1:]
		tempQ := stackQ[0]
		stackQ = stackQ[1:]

		if tempP == nil {
			if tempQ != nil {
				return false
			}
			continue
		}

		if tempQ == nil {
			if tempP != nil {
				return false
			}
			continue
		}

		if tempP.Val != tempQ.Val {
			return false
		} else {
			stackP = append(stackP, tempP.Left, tempP.Right)
			stackQ = append(stackQ, tempQ.Left, tempQ.Right)
		}
	}

	return len(stackP) == len(stackQ)
}
