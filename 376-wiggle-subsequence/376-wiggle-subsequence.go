func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func wiggleMaxLength(nums []int) int {
    if len(nums) < 3 {        
        set := make(map[int]bool)
        for _,c := range nums{
            set[c]= true
        }
        return len(set)
    }
    return getMin(solution(nums, 1,1,1),len(nums))
}

func solution(nums []int,current,pos,neg int) int{   
    if current > len(nums)-1{
        return getMax(pos,neg)
    }
 
    if nums[current-1] > nums[current]{            
        return solution(nums,current+1,neg+1,neg)            
    }else if nums[current-1] < nums[current] {
        return solution(nums,current+1,pos,pos+1) 
    }else{
        return solution(nums,current+1,pos,neg) 
    }
    
    return 1
}

// Input: nums = [1,17,5,10,13,15,10,5,16,8]
// two options positive or negative
// one must be true 
// max value 
// end when both are false 
// update max Value every time we finish
// res = current - prev - nums[i] - nums[i-1]
// if neigther conditions are true, reset the count 

// 1,7,4,9,2,5

// res   0 1 2 3
// i     1 2 3 4
// count 0 1 2 3  
// current 6 -3 5 -7 3


//[1,17,5,10,13,15,10,5,16,8]
//[[1,0],[1,1],[2,1],[2,1]]
// prev= p,n,p, 
// [1,16,-12,5,3,2,-5,-5,11,-8] 
// [1, 2, 3, 4, 4, 4, 4, 5, 6, 8]

// if prev == pos then next < current