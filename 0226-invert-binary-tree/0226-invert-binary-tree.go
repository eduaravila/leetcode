/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil{
        return root
    }
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {        
        element := queue[0]
        queue = queue[1:]
        if element.Left != nil {
            queue= append(queue,element.Left)
        }
        if element.Right!= nil{
            queue= append(queue,element.Right)
        }
        element.Left, element.Right= element.Right,element.Left
    }
    return root
}