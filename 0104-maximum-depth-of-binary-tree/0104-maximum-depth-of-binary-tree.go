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

func maxDepth(root *TreeNode) int {
    var max int
    
    solution(root,1, &max)
    return max
}

func solution(root *TreeNode,current int, max *int){
    if root == nil{
        return 
    }
    *max = getMax(current,*max)
    solution(root.Left,current+1,max)
    solution(root.Right,current+1,max)    
}