/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
    if root == nil{
        return ""
    }
    
    l := tree2str(root.Left)
    r := tree2str(root.Right)
    
    if l=="" && r ==""{
        return fmt.Sprintf("%d",root.Val)
    }
    
    if l == "" {
        return fmt.Sprintf("%d%s%s",root.Val ,fmt.Sprintf("()")  ,fmt.Sprintf("(%s)",r) )
    }
    if r == ""{
        return fmt.Sprintf("%d%s",root.Val, fmt.Sprintf("(%s)",l) )
    }
    return fmt.Sprintf("%d%s%s", root.Val, fmt.Sprintf("(%s)",l), fmt.Sprintf("(%s)",r) )
}