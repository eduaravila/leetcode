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
    res := solution(root)
    return getMax(res[1],res[0])
}

func solution(root *TreeNode)[]int{
    if root == nil{
        return []int{0,0}
    }
    
    a := solution(root.Left)
    b := solution(root.Right)
    
    return []int{root.Val + b[1] + a[1], getMax(getMax(a[1]+b[0],b[1]+ a[0]),getMax(b[0]+ a[0],a[1] + b[1]) ) }
}