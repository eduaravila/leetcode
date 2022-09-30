func isNStraightHand(nums []int, groupSize int) bool {
    
    if len(nums) % groupSize != 0 {
        return false
    }
    
	var min, currentGroupValue, currentGroupSize int
    sort.Ints(nums)
    counter := make(map[int]int)
    
    for _,num := range nums {
        counter[num]++
    }
    
	for min < len(nums) { 

        for currentGroupSize < 1 && min < len(nums) {

            if counter[nums[min]] > 0 {
                currentGroupSize++
                currentGroupValue = nums[min]
                counter[nums[min]]--
            }else{
                min++
            }
        }
        if min >= len(nums) -1 {
            break
        }
        if currentGroupSize >= groupSize {
            currentGroupValue = 0
            currentGroupSize = 0
            continue
        }
                                
        if counter[currentGroupValue + 1 ] < 1 {
            return false
        }
        
        currentGroupValue++
        currentGroupSize++
        counter[currentGroupValue]--
                
        
    }
    
	
	return true  

}