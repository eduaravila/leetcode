func getMin(a,b int)int{
    if a < b {
        return a
    }
    return b
}

func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
    
    var i int
    res := [][]int{}
    
    for i < len(intervals){
        current := intervals[i]
        i++        
        
        for i < len(intervals) && current[1] >= intervals[i][0] && current[0] <= intervals[i][1]{
            current = []int{getMin(current[0], intervals[i][0]),getMax(current[1], intervals[i][1])} 
            i++
        }
        
        res = append(res,current)

    }
    return res
}


/*
 [[1,3],[2,6],[8,10],[15,18]]
  [1,6], [8,10],[15,18]
  
  
  
  [1,4],[4,5]
  [1,5] ->
   a,b
   
   [1,4], [1,9] -> [1,9]
   [1,4], [2,9] -> [1,9]
   
   [1,1], [1,2] -> [1,2]
   
   [1,10],[9,10] -> [1,10]
   
   [1,2], [3,9] ->  X
   
    two intervals overlap if 
    
   b1 >= a2 && a1 <= b2
    
    [1,4],[0,0]
    
    [0,0],[1,4]
    

        1,2,3,4
      0
    
   
   
    newinterval= [getMin(a1,a2), getMax(b1,b2)]
    current = newinterval
    
    
     [[1,3],[2,6],[8,10],[15,18]]
      
      current = [8,10]
      
      result = [1,6],
  
  
[[2,3],[4,5],[6,7],[8,9],[1,10]]

[1,10],[2,3],[4,5],[6,7],[8,9]


*/