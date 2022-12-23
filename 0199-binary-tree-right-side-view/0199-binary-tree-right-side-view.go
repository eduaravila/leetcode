/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil{
        return []int{}
    }
    res := []int{}
    queue := []*TreeNode{root}
    for len(queue) > 0{
        n := len(queue)-1
        res = append(res,queue[n].Val)
        
        for _,element := range queue{
            if element.Left != nil{
                queue = append(queue,element.Left)
            }
            if element.Right != nil{
                queue = append(queue,element.Right)
            }
        }
        queue = queue[n+1:]
    }
    return res
}