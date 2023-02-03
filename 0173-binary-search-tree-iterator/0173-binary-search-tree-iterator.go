/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    i int 
    values []int
}


func getValues(root *TreeNode,values *[]int){
    if root == nil{
        return
    }
    
    getValues(root.Left,values)
    *values = append(*values, root.Val)
    getValues(root.Right,values)
}

func Constructor(root *TreeNode) BSTIterator {
    values := []int{}
    getValues(root,&values)
    return BSTIterator{
        values: values,
        
    }
}


func (this *BSTIterator) Next() int {
    res := this.values[this.i]
    this.i++
    return res
}


func (this *BSTIterator) HasNext() bool {
    if this.i < len(this.values){
        return true
    }
    return false
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */