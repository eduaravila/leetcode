
func numSubseq(nums []int, target int) int {
    sort.Ints(nums)
    var res,l int
    mod := int(math.Pow10(9) + 7) 
    r := len(nums)-1
    convinations:= make([]int,len(nums))
    convinations[0] = 1
    for i := 1 ;i < len(nums) ; i++{
        convinations[i] = convinations[i-1] * 2 % mod    
    }
    
    for l <= r{        
        if nums[r] + nums[l] > target{
            r--
        }else {
            res += convinations[r-l]
            l++
        }
        
    }
    return res % mod
}

// [3,3,3,8]

//k = 10
// [3] = 6
// [3,3] = 6
// [3,3,6] = 9
// [3,3,6,8]= 3
// l= 0 1
// r= 3 3
// m= 1 2
// 
