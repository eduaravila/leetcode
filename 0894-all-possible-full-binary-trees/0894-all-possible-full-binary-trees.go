/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
    
    return solution(n, make(map[int][]*TreeNode))    
}

func solution(n int, memo map[int][]*TreeNode) []*TreeNode{ 
    if val,e := memo[n];e{
        return val
    }
    
    if n == 1 {
        return []*TreeNode{&TreeNode{
            Val:0,
        }}
    }
    
    res := []*TreeNode{}
    n--
    for l := 1 ; l < n ; l++{
        r := n - l
        lr,rr := solution(l,memo),solution(r,memo)
        memo[l] = lr
        memo[r] = rr
        for _,nodel := range lr {            
            for _,noder := range rr{            
                res = append(res,&TreeNode{Val:0, Left:nodel,Right:noder})
            }
            
        }
    }
    
    return res
} 