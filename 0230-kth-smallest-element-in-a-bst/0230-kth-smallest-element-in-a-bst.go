/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// tree is already ordered
// inorder should sufice
// remove from k -1 every time we add a value
// when k == 0 we found the k smallest
// no need to keep all the values
func kthSmallest(root *TreeNode, k int) int {
    var value int    
    solution(root,&value, &k)
    return value
}

func solution(root *TreeNode, value *int, k *int){
    if root == nil || *k < 1 {
        return 
    }
    if root.Left!= nil{        
        solution(root.Left,value, k)    
    }
    if *k > 0 {
        *value = root.Val
    }
    *k-=1    
    if root.Right != nil{        
        solution(root.Right,value,k)
    }
}