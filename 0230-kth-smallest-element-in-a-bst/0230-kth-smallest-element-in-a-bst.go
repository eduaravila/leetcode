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
    current := k
    solution(root,&values, &current)    
    return values[k-1]
}

func solution(root *TreeNode, values *[]int, k *int){
    if root == nil || *k < 1{
        return
    }
    if root.Left!= nil{        
        solution(root.Left,values, k)    
    }
    *k-=1
    *values = append(*values,root.Val)
    if root.Right != nil{        
        solution(root.Right,values,k)
    }
}