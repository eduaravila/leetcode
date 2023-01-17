/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    vals := []int{}
    solution(root,&vals)
    return vals[k-1]
}

func solution(root *TreeNode , vals *[]int) {
    if root == nil{
        return 
    }
    
    solution(root.Left,vals)
    *vals = append(*vals,root.Val)
    solution(root.Right,vals)
}