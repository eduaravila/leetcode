func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(a,b int)bool{
        return  intervals[a][1] <= intervals[b][1]
    })
    var i int
    var total int
    
    for i < len(intervals){
        current := intervals[i]
        i++
        for i < len(intervals) && current[1] > intervals[i][0]{ 
            total++
            i++
        }
    }
    
    return total
}

/*
T: O(n log n)
S: O(1)

[[1,2],[2,3],[3,4],[1,3]]
[1,2], [2,3], [1,3],[3,4]

sort: [1,2] ,[1,3],[2,3],[3,4]

start1, end1 | start2, end2

collaps if:
end1 > start2 
    total++


[1,2] ,[1,3],[2,3],[3,4]

c = 1,3


- we dont convine overlaps just remove them


[1,2] -> 0

[1,1],[1,1] ->1
i

[[1,100],[11,22],[1,11],[2,12]]

[1,100],[1,11],[2,12],[11,22] -> 3 , [1,100]

[1,11], [1,100], [2,12], [11,22]
[1,11], [2,22], [11,22], [1,100]


*/