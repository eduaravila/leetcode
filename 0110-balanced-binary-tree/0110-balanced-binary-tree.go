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

func isBalanced(root *TreeNode) bool {
    if root== nil{
        return true
    }    
    res := isBalanced(root.Left) && isBalanced(root.Right)
    l:= getDepth(root.Left,0)
    r:= getDepth(root.Right,0)    
    return int(math.Abs(float64(l-r))) <= 1 && res
}

func getDepth(root *TreeNode,level int)int{
    if root == nil {
        return level
    }
    return getMax(getDepth(root.Left,level+1),getDepth(root.Right,level+1))
}