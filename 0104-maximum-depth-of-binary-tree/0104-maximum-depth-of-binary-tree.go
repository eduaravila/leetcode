/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    queue := []*TreeNode{root} 
    
    var depth int
    for len(queue) > 0{
        n := len(queue)
        depth++
        for _, element := range queue{
            if element.Left != nil{
                queue = append(queue,element.Left)
            }
            if element.Right != nil{
                queue = append(queue,element.Right)
            }
        }
        queue = queue[n:]
    }
    return depth
}

