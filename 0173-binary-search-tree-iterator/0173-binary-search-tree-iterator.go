/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {    
    levels []*TreeNode
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
    
    return BSTIterator{
        levels: []*TreeNode{root},        
    }
}



func (this *BSTIterator) Next() int {
    var res *TreeNode
    for this.HasNext() && res == nil {        
        n := len(this.levels)-1
        node := this.levels[n]
        if node.Left != nil{
            this.levels = append(this.levels, node.Left)
            node.Left = nil
            continue
        }
        res = node
        this.levels = this.levels[:n]

        if node.Right != nil{
            this.levels = append(this.levels, node.Right)
            node.Right = nil
            continue            
        }
        
    }
    return res.Val
       
}


func (this *BSTIterator) HasNext() bool {
    return len(this.levels) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */