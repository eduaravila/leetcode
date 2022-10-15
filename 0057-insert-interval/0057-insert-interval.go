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

func insert(intervals [][]int, newInterval []int) [][]int {
    
    startNew,endNew := newInterval[0], newInterval[1]
    res := [][]int{}
    var i int
    
    // insert smaller values than interval
    for i < len(intervals) && intervals[i][1] < startNew {
        res = append(res, intervals[i])
        i++
    }
    
    mergeInterval := []int{startNew,endNew}
    for i <len(intervals) && mergeInterval[1] >= intervals[i][0]{
        mergeInterval = []int{getMin(mergeInterval[0],intervals[i][0]), getMax(mergeInterval[1],intervals[i][1])}
        i++
    }
    
    res= append(res, mergeInterval)
    // insert bigger values
    res = append(res, intervals[i:]...)
    return res
}

// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// i 0, 1, 2, 3, 4
// res [1,2], [3,10], [12,16]
// startNew 4
// endNew 8
// mergeInterval 3, 10