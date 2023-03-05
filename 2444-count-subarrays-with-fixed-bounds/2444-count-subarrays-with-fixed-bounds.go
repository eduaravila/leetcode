func getMax(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func getMin(a, b int)int{
    if a < b{
        return a
    }
    return b
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
    l, min ,max := -1,-1,-1
    var res int64
    for i, num := range nums{
        if num > maxK || num < minK{
            l = i
        }
        if num == minK{
            min = i
        }
        if num == maxK{
            max = i
        }
        
        res += int64(getMax(0, getMin(min,max) - (l)))
    }
    return res
}

/*

[1,3,5,2,7,5]

out = -1
min = 0
max = -1, 2

0-1


3 pointer 

min 
max 
and left bound is > than max and < min if out of bounds


*/
