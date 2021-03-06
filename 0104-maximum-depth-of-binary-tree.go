// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: 3
// Example 2:

// Input: root = [1,null,2]
// Output: 2

// Constraints:

// The number of nodes in the tree is in the range [0, 104].
// -100 <= Node.val <= 100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	maxDepth := 0
	if root == nil {
		return maxDepth
	}

	curLevel := []*TreeNode{root}

	for len(curLevel) > 0 {
		maxDepth++
		tmp := []*TreeNode{}
		for i := range curLevel {
			if curLevel[i].Left != nil {
				tmp = append(tmp, curLevel[i].Left)
			}

			if curLevel[i].Right != nil {
				tmp = append(tmp, curLevel[i].Right)
			}
		}
		curLevel = tmp
	}

	return maxDepth
}