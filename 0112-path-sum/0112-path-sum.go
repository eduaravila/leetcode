/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    return solution(root, targetSum)
    
}

func solution(root *TreeNode, targetSum int) bool{
    if root != nil && targetSum - root.Val == 0 && root.Left == nil && root.Right == nil {        
        return true
    }    
    if root == nil { 
        return false
    }
    targetSum -= root.Val
    return solution(root.Left, targetSum) || solution(root.Right, targetSum)
} 