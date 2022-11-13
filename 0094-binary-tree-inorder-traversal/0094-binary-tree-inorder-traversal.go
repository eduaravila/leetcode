/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    solution(root,&res)
    return res
}

func solution(root *TreeNode, res *[]int){
    if root == nil{
        return
    }        
    solution(root.Left,res)
    *res = append(*res,root.Val)
    solution(root.Right,res)
}