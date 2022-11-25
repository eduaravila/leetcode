func splitArray(nums []int, k int) int {
    return solution(nums,0,k,make(map[string]int))  
}

func getMin(a, b int)int{
    if a < b{
        return a
    }
    return b
}

func getMax(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func solution(nums []int , i,k int, memo map[string]int) int{
    key := fmt.Sprintf("%d-%d",i,k)
    if val,e:= memo[key];e{
        return val
    }
    
    if k == 1{
        var sum int
        for i < len(nums){
            sum += nums[i]
            i++
        }
        return sum
    }
    var max, sum int
    res := int(^uint(0)>>1)
    for i <= len(nums)-k{ // we substract k to give other subarrays at least have 1 element
        sum += nums[i]
        max = getMax(sum,solution(nums,i+1,k-1,memo))
        res = getMin(max,res)
        if sum > res {
            break
        }
        i++
    }
    memo[key] = res
    return res
}