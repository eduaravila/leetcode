/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root == nil{
        return root
    }
    
    if root.Val < low{
       root = trimBST(root.Right,low,high)
        
    }else if root.Val > high{
       root = trimBST(root.Left,low,high)
    }
    
    if root == nil{
        return root
    }
    
    root.Left = trimBST(root.Left,low,high)
    root.Right = trimBST(root.Right,low,high)
    
    return root
}