
/*
    
    
   1,3,5,4,7
   
   1,3,2,2,1
   1,1,1,1,1
*/

func getMax(a,b int)int{
    if a>b {
        return a
    }
    return b
}

func findNumberOfLIS(nums []int) int {
    l := make([]int,len(nums))
    counter := make([]int,len(nums))
    
    for i := range nums{
        l[i] = 0
        counter[i] = 0
    }
    
    var res, lres int
    for i := len(nums)-1 ; i >=0 ; i--{
        ln,cnt := 1,1
        
        for x := i+1 ; x < len(nums) ; x++{
            a,b := nums[i], nums[x]
            
            if a < b{
                if  ln < l[x] +1 {
                    ln,cnt = l[x]+1, counter[x]
                }else if l[x]+1 == ln{
                    cnt += counter[x]
                }
            }
        }
        if ln > lres {
            res = cnt
            lres = ln
        }else if ln == lres{
            res += cnt
        }
        l[i] = ln
        counter[i] = cnt
    }

    return res
}