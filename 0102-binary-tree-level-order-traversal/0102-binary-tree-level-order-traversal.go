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
    levels := [][]int{}
    
    queue := []*TreeNode{root}
    level := []int{}
    for len(queue) > 0{
        l := len(queue)
        level = []int{}
        for _, node := range queue{
            level = append(level, node.Val)
            if node.Left != nil{
                queue =append(queue,node.Left)
            }
            if node.Right != nil{
                queue =append(queue,node.Right)
            }
        }
        levels =append(levels, level)
        queue = queue[l:]
    }
    return levels
}