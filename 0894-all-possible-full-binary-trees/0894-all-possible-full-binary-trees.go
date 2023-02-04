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
    if n == 1{
        return []*TreeNode{
            &TreeNode{
                Val:0,
            },
        }
    }

    n--
    res := []*TreeNode{}
    for l := 1 ; l < n ;l++{
        r := n - l
        lnodes, rnodes := solution(l),solution(r)
        
        for _, ln := range lnodes{            
            for _, rn := range rnodes{                        
                res = append(res,  &TreeNode{Val:0,Left:ln, Right: rn})
            }
            
        }
    }
    
    return res
}