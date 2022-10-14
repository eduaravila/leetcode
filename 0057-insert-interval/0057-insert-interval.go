func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
    a_new,b_new := newInterval[0],newInterval[1]
    var i int
    res := [][]int{}
    for i < len(intervals) && intervals[i][1] < a_new {
        res = append(res,intervals[i])
        i++
    }
    
    new_range := []int{a_new,b_new}
    for i < len(intervals) && intervals[i][0] <= b_new  {
        new_range = []int{getMin(new_range[0],intervals[i][0]),getMax(new_range[1],intervals[i][1])}
        i++
    }
    
    res = append(res,new_range)
    
    res = append(res,intervals[i:]...)
    return res
}

// intervals = [[1,3],[6,9]], newInterval = [2,5]

// a_new = 2, b_new = 5
// start: 0 
// current: [1,3]
// end: 0 
// current: [1,3]


// a_new = 4 b_new = 8
// start = 0,1
// current [1,2],[3,5]
// end =0,1,2,3
// current [1,2],[3,5]




// a,b = current[0],current[1]
// a_range,b_range := interval[i][0],interval[i][1]
// [1,2],[3,4] , new_interval = [5,6]
// [1,2],[3,4],[5,6]

// [1,2],[5,6] , new_interval = [3,4]
// [1,2],[3,4],[5,6]

// o(n)
// 