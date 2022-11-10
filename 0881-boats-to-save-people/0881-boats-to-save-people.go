func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    var res int
    l,r := 0,len(people)-1
    for r >=0 && people[r] >=limit{
        r--
        res++
    }
    
    for l < r{
        if people[r] + people[l] <=limit{            
            l++
            r--
        }else{
            r--            
        }
        res++
    }
    if l == r{
        res++
    }
    return res
    // the most people the better
}
// 1,2,2,3

// r = 2
// l = 0