/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    values := []int{}
    queue := []*TreeNode{root}
    for len(queue) > 0{
        element := queue[0]
        values = append(values, element.Val)
        queue = queue[1:]
        if element.Left != nil{
            queue = append(queue,element.Left)
        }
        if element.Right!= nil{
            queue = append(queue,element.Right)
        }
        
    }
    sort.Ints(values)
    return values[k-1]
}