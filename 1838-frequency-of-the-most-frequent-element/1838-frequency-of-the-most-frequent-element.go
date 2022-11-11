func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    n:= len(nums)-1
    l,r := n,n
    var max int
    for l >=0 {               
        k -= nums[r] - nums[l]        
        if k < 0{
            k += (nums[r] - nums[r-1]) * (r-l)
            r--
        }
        max = getMax(max,(r-l)+1)
        l-- 
    }
    return max
}

// [[3,6,9]] k=2

// output = 3

//l=4
//r=4
//