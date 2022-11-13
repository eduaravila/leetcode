/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    q1 := []*TreeNode{p}
    q2 := []*TreeNode{q}
    
    for len(q1) > 0 && len(q2) > 0 {
        element1:= q1[0]
        element2:= q2[0]
        q1= q1[1:]
        q2= q2[1:]
        if element1 == nil && element2 == nil{
            continue
        }
        if element1 != nil && element2 == nil ||  element1 == nil && element2 != nil || element1.Val != element2.Val {
            return false
        }
        
        q1 = append(q1,element1.Right)
        q1 = append(q1,element1.Left)

        q2 = append(q2,element2.Right)
        q2 = append(q2,element2.Left)
        
        

    }
    return true
}