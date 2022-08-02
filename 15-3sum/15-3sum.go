import "sort"
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    if len(nums) < 3 {
        return [][]int{}
    }
    m,r := 1,len(nums)-1
    res := [][]int{}
    for l := range nums {    
        if l > 0 && nums[l] == nums[l-1]{
            continue
        }
        m, r = l+1, len(nums)-1
        for m < r {
           cSum := nums[l] + nums[m] + nums[r]                                                         
            if cSum > 0 { 
                r--            
            }else if cSum < 0 {
                m++   
            }else{
                res = append(res,[]int{nums[l], nums[m] , nums[r]})
                m++                
                for m < r && nums[m] == nums[m-1]{
                    m++
                }
            }
        }
    }
    return res
}

// sort [-1,0,1,2,-1,-4]
// [-4, -1, -1, 0, 1, 2]
// l = 0 
// r = len(nums) - 1
// -4 + 2 = -2 
// middle = l + 1 
// -2 + -1 = -3
// -2 + -1 = -3
// -2 + 0 = -2
// -2 + 1 = -1 less than 0 move l 
// l = 1 
// r = len(nums) -1 
// -1 + 2 = 1
// 1 + -1 = 0 // found match 

// check if l == l-1
// to omit repeated patters
// if we found a result r++
// if result < 0 l++
// check for repeated l 
// middle is  = l + 1 
