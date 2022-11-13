/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMax(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func isBalanced(root *TreeNode) bool {
    r, _ := solution(root)
    return r
}

func solution(root *TreeNode)(bool,int){
    if root == nil{
        return true,0
    }
    lb,ll := solution(root.Left)    
    rb,rl := solution(root.Right)
    
    balanced := int(math.Abs(float64(ll-rl))) <=1 && lb && rb
    return balanced, getMax(ll,rl)+1
}