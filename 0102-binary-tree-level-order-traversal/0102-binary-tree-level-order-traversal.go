/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil{
        return [][]int{}
    }
    res := [][]int{}
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0{
        current := []int{}
        l := len(queue)
        for _,element := range queue{
            current = append(current,element.Val)
            if element.Left != nil{
                queue = append(queue,element.Left)
            }
            if element.Right != nil {
                queue = append(queue,element.Right)
            }
        }
        res = append(res,current)
        queue = queue[l:]
    }
    return res
}