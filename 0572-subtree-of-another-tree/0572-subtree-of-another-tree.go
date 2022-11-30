/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil{
        return false
    }
    if subRoot == nil{
        return true
    }
    if isEqual(root,subRoot){
        return true
    }
    return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)
}

func isEqual(a,b *TreeNode)bool{

    if (a == nil && b!= nil )|| (b == nil && a != nil ){
        return false
    }
    if a == nil && b == nil{
        return true
    }

    if a.Val != b.Val {
        return false
    }
    return isEqual(a.Right,b.Right) && isEqual(a.Left,b.Left)
}

