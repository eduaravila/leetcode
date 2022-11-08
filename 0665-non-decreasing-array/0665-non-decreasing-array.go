func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}


func checkPossibility(nums []int) bool {    
    l := 0
    var changes int
    inf := int(^uint(0)>>1)
    max := -inf
    for i := 1; i < len(nums) ;i++{        
        if nums[i] < nums[l] && nums[i] < max{
            nums[i] = nums[l]
            changes++
        }else if nums[i] < nums[l] {
            nums[l] = nums[i]
            changes++
        }
        max = getMax(max,nums[l])
        l++
    }
   
    return changes < 2
}


//[1,4,2,3]
//[1,2,3,4,3]
//[2,3,1,3] = true
//[2,4,0,2], lm 4 rm 2
// r to l
// max = len(nums)-1
// if we find another max then that number is not in the correct place

// [1,3,4,5,6,0]

// [1,2,0,4,5,6]