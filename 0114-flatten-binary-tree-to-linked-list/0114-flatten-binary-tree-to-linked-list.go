/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil{
        return
    }
    stack := []*TreeNode{}
    solution(root, &stack)
    fmt.Println(stack)
    
    var prev *TreeNode
    current:= stack[len(stack)-1]
    for len(stack) > 0{        
        current.Right = prev
        current.Left = nil
        prev = current
        stack = stack[:len(stack)-1]
        if len(stack) == 0{
            break
        }
        current = stack[len(stack)-1]
    }
        
}

func solution(root *TreeNode, stack *[]*TreeNode){
    if root == nil{
        return
    }
    
    *stack = append(*stack, root)
    solution(root.Left,stack)
    solution(root.Right,stack)
    
}