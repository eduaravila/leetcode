/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }   
    
    if n == 1 {
        return []*TreeNode{
            &TreeNode{Val:0},
        }
    }    
    n--
    res := []*TreeNode{}
    for l := 1 ; l < n ; l++{
        r := n - l
        
        lr,rr := allPossibleFBT(l),allPossibleFBT(r)
        
        for _, noder := range rr{
            for _, nodel := range lr{
                root := &TreeNode{Val:0,Left:nodel,Right:noder}
                res = append(res,root)
            }
        }        
    }
    return res
}