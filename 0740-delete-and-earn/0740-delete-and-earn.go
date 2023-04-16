/*
    
    1 -> 1
    
    
    1,2,3 => 4 
    
    base case is len(input) < 2
    return input[0]
    
    
    prev := nums[i]-1
    next := nums[i]+1
    set[prev] = 0
    set[next] = 0
    set[i]--
    solution(nums, set, total + nums[i], )
    set[i]++
    set[prev] = tempPrev
    set[next] = tempNext




    [2,2,3,3,3,4]
    
    [
        2:2
        3:3
        4:1
    ]
    
    [2,3,4]
    
    
    2 -> 4 + 4 -> 8
    3 -> 9
    4 -> 4 + 4 -> 8
    
    [0,2,3,4]
    
    [0,4,9,4]
 */

func getMax(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func deleteAndEarn(nums []int) int {
    set := make(map[int]int)
    

    for _, num := range nums{
        set[num]++
    }
    
    unique := []int{}
    for num := range set{
        unique = append(unique, num)
    }

    sort.Ints(unique) 
    
    var a,b int   
    for i := 0 ; i < len(unique) ; i++{
        prev := unique[i] -1
        num := unique[i]
        if i > 0 && unique[i-1] == prev{            
            temp := b
            b = getMax(b, num * set[num] + a)
            a = temp
        }else{
            temp := b
            b += num * set[num]
            a = temp
        }
    }
    
    return getMax(a,b)
    
}
