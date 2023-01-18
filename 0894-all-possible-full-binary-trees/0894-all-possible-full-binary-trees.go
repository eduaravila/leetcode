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
    if n == 1 {
        return []*TreeNode{&TreeNode{
            Val:0,
        }}
    }
    
    res := []*TreeNode{}
    n--
    for l := 1 ; l < n ; l++{
        r := n - l
        lr,rr := solution(l),solution(r)
        for _,nodel := range lr {            
            for _,noder := range rr{
                root := &TreeNode{Val:0}
                root.Left = nodel
                root.Right = noder
                res = append(res,root)
            }
            
        }
    }
    return res
} 