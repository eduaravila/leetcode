/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMax(a,b int)int{
    if a > b{
        return a
    }    
    return b
}

func maxDepth(root *TreeNode) int {
    return solution(root,0)
}
func solution(root *TreeNode, level int)int{
    if root == nil{
        return level
    }
    
    return getMax(solution(root.Left,level+1),solution(root.Right,level+1))
}