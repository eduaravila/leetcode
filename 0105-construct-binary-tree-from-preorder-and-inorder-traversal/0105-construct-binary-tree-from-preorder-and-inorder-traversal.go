/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    
    indexes := make(map[int]int)
    for i,v := range inorder {
        indexes[v] = i
    }
    current := 0
    return solution(0,len(inorder)-1,preorder, indexes,&current)    
}
func solution( l,r int, preorder []int,indexes map[int]int,current *int) *TreeNode {
    if l > r{
        return nil
    }    
    root := &TreeNode{Val:preorder[*current]}
    *current = *current+1
    root.Left  = solution(l,indexes[root.Val]-1,preorder, indexes,current)    
    root.Right  =solution(indexes[root.Val]+1,r,preorder,indexes, current)
    return root
}