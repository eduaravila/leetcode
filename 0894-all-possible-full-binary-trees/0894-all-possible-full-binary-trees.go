/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {    
    return solution(n)
}

func solution(n int) []*TreeNode{
    if n == 0  {
        return []*TreeNode{}
    }
    if n == 1 {
        return []*TreeNode{&TreeNode{Val: 0}}
    }        
    res := []*TreeNode{}
    n--
    for i := 1 ; i < n ; i++{
        lt,rt := solution(i), solution(n-i)
        for _, t1 := range lt{  
            for _, t2 := range rt{
                res = append(res, &TreeNode{Val:0, Left: t1, Right: t2})
            }
        }
    }
    
    return res 
}