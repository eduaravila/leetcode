/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {    
    res := root
    solution(root,p,q, res)
    return res
}


func solution(root, p, q, res *TreeNode)bool {
    if root == nil{
        return false
    }
    if solution(root.Left,p,q, res) || solution(root.Right,p,q, res){
        return true
    }
    
    if hasChildren(root,p) && hasChildren(root,q){
        *res = *root
        return true
    }
    return false
    
}

func hasChildren(root, p *TreeNode)bool{
    if root == nil{
        return false
    }
    
    if root.Val == p.Val{return true}
    return hasChildren(root.Left, p) ||  hasChildren(root.Right, p)
}