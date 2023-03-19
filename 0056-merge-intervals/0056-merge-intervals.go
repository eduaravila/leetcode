func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func merge(intervals [][]int) [][]int {    
    sort.Slice(intervals, func(a,b int)bool{
        return intervals[a][0] < intervals[b][0]
    })
    
    var start, end, i int
    res := [][]int{}
    a,b := 0,1
    max := intervals[0][b]
    for end < len(intervals){
        max = getMax(max,intervals[end][b])
        for i < len(intervals) && intervals[start][a] <= intervals[i][a] && max >= intervals[i][a] {            
            end = i
            max = getMax(max,intervals[end][b])
            i++            
        }
        
        res = append(res, []int{intervals[start][a], max})
        end++
        start = end
    }
    
    return res
}


/*
[[1,3],[2,6],[8,10],[15,18]]


start int
end int

a = 0
b = 1


for end < len(intervals) 
    merge more than 2 
    if arr[start][a] <= arr[i][a] && arr[end][b] >= arr[i][a] && arr[end][b] <= arr[i][b]
        end = i

    update the intervals from start to end, end should have the largest value
    
    start = end

*/