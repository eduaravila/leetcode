import "sort"
func threeSum(input []int) [][]int {
    sort.Ints(input)
    res := [][]int{}
    
    
    for i,val := range input {                 
        if i > 0 && val == input[i-1]{
            continue
        }
        l,r := i+1, len(input)-1    
        
        for l < r {
            sum := val + input[l] + input[r]
            if sum > 0 { 
                r--
            }else if sum < 0 {
                l++
            }else{
                res = append(res,[]int{val, input[l],input[r]})
                l++
                for l < r && input[l] == input[l-1] {
                    l++
                }
            }
            
        }        
                             
    }
    
    return res
}