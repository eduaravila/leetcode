/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root != nil && targetSum - root.Val == 0 && root.Left == nil && root.Right == nil {        
        return true
    }    
    if root == nil { 
        return false
    }
    targetSum -= root.Val
    return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
    
}
