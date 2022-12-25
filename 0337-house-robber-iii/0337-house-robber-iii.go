/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMax(a,b int)int{
    if a > b {
        return a
    }    
    return b
}

func rob(root *TreeNode) int {
    res := solution(root,0,0)
    return getMax(res[0],res[1])
}

func solution(root *TreeNode, sum, prev int)[]int{
    if root == nil{
        return []int{0,0}
    }
    
    l := solution(root.Left, sum, root.Val+sum)
    r := solution(root.Right, sum, root.Val+sum)
    
    return []int{root.Val + l[1] + r[1], getMax(l[1],l[0]) + getMax(r[0],r[1])}
}