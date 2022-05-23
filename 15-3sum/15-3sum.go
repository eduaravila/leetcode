import 
(
"fmt"
)

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    used := make(map[string]bool)
    current := 0 
    
    for current < len(nums){
        l, r := current+1,len(nums)-1
        
        for l < r {            
            remaining := nums[l] + nums[r] + nums[current]
            
            
            if remaining > 0 {
                r--
            }else if remaining < 0 {
                l++
            }else {
                key := fmt.Sprintf("%d-%d-%d",nums[current],nums[l],nums[r])
                if _,e := used[key];!e{
                    res = append(res,[]int{nums[current],nums[l],nums[r]})                                    
                    used[key]=true
                }
                l++
                r--                                
            }
        }
        current++
    }
    return res
}