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

func getDepth(root *TreeNode, level int)int{        
    if root == nil{
        return level
    }    
    return getMax(getDepth(root.Left,level+1),getDepth(root.Right,level+1))
}

func diameterOfBinaryTree(root *TreeNode) int {    
    var max int
    getLevel(root,&max)
    return max
}

func getLevel(root *TreeNode, max *int){        
    if root==nil{
        return
    }
    l:= getDepth(root.Left,0)
    r:= getDepth(root.Right,0)    
    *max = getMax(*max, l+r)
    getLevel(root.Left,max)
    getLevel(root.Right,max)
}