func getIndex(arr []int, target int)int{
    l,r := 0,len(arr)-1
    
    for l <r{
        m := ((l+r) >>1)
        if arr[m] == target || l == m{
            return m
        }
        if arr[m] > target{
            r = m-1
        }else{
            l = m
        }
    }
    return l
}

func findClosestElements(arr []int, k int, x int) []int {
    index := getIndex(arr,x)
    res := []int{}    
    l,r := index,index+1
    for k > 0  {
        k--
        if l < 0{
            res = append(res,arr[r])
            r++
            continue
        }
        if r > len(arr)-1{
            res = append(res,arr[l])
            l--
            continue
        }
        a := x - arr[l] 
        b := arr[r] - x
        closest := arr[l]
        l--
        if a > b{
            closest = arr[r]            
            r++
            l++
        }        
        
        res =append(res,closest)

    }
    sort.Ints(res)
    return res
}

// a,b


// a-current  <= b - current

//a
//b
// 1,2,3,4,5
// 2,3,4,5,6