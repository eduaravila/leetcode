/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    values := []int{}
    solution(root,&values)
    return values[k-1]
}
func solution(root *TreeNode, values *[]int){
    if root == nil{
        return
    }
    if root.Left!= nil{
        solution(root.Left,values)    
    }
    *values = append(*values,root.Val)
    if root.Right != nil{
        solution(root.Right,values)
    }
}