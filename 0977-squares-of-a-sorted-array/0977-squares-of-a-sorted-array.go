func mergeSort(input *[]int, l,r int)int{
    
    pivot :=(*input)[r]
    i := l-1
    for l <= r{        
        if (*input)[l] < pivot {
            i++
            (*input)[i],(*input)[l] = (*input)[l], (*input)[i]            
        }      
        l++
    }
    (*input)[r],(*input)[i+1] = (*input)[i+1], (*input)[r]
    return i+1
}


func sort(input *[]int, l,r int){
    if l >= r{
        return 
    }
    p := mergeSort(input,l,r)
    sort(input,p+1, r)
    sort(input,l, p-1)
}

func sortedSquares(nums []int) []int {    
    for i, num := range nums{
        nums[i] = int(math.Pow(float64(num),2))
    }    
    
    sort(&nums,0,len(nums)-1)
    return nums
}

// [16,1,0,9,100]
// l 0
// r 4
// p 16

// [9,1,0,16,100]

