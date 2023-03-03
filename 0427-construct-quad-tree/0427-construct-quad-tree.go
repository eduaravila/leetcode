/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func Itob(a int)bool{
    if a == 0 {
        return false
    }
    return true
}

func allEqual(grid [][]int, x,y, size int)bool{
    last := grid[x][y]
    
    for _, col := range grid[x:x+size]{
        for _, val := range col[y:y+size]{
            if last != val{
                return false
            }        
        }
    }
    return true
}

func construct(grid [][]int) *Node {
    
    return solution(grid,0,0,len(grid))
    
}

func solution(grid [][]int, x,y, size int) *Node{    
    if allEqual(grid,x,y,size){
        return &Node{Val: Itob(grid[x][y]),IsLeaf:true}
    }
    
    node := &Node{}
    node.TopLeft = solution(grid, x ,y, size/2) 
    node.TopRight = solution(grid, x,y+size/2, size/2) 
    node.BottomRight = solution(grid, x+ size/2, y+size/2, size/2) 
    node.BottomLeft = solution(grid, x + size/2,y, size/2) 
    
    return node
}