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
type NodeMax struct{
    node *TreeNode    
    max int
}

func goodNodes(root *TreeNode) int {
    var res int
    queue := []NodeMax{{root,root.Val}}
    
    for len(queue) > 0 {
        n := len(queue)
        for _, e := range queue{
            element,max := e.node, e.max
            if element.Val >= max{
                res++
            }
            if element.Left != nil{
                queue = append(queue,NodeMax{element.Left,getMax(max,element.Left.Val)})
            }
            if element.Right != nil{
                queue = append(queue,NodeMax{element.Right,getMax(max,element.Right.Val)})
            }
        }
        queue = queue[n:]
    }
    
    return res
}