/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    res := []int{}
    solution(root,&res)
    return res
}
func solution(root *TreeNode, res*[]int){
    if root == nil{
        return
    }
    
    *res = append(*res, root.Val)
    solution(root.Left,res)
    solution(root.Right,res)
}