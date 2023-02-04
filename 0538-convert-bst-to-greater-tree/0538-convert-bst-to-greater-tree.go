/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    var sum int
    solution(root,&sum)
    return root
}

func solution(root *TreeNode, sum *int) {
    if root == nil{
        return 
    }
    
    solution(root.Right,sum)
    *sum += root.Val
    root.Val = *sum
    solution(root.Left, sum)
    
}