/*

numbers = [1,2,3....n]

n == 1 -> 1

n == 2

k = 1


n == 3 

k = {1,2} ,{ 1,1,1}


1 => n

max int


if n == 1 {
    max = max(productof(numbers), max)
}

3.2


3.3 9
2,2,2 8

3.3.2

2.2.2.2


4.3

3.3.1
2.2.2.1

2.2.2

3.3

3.2.1

9

3.3.3
5 2 2 



from 1 to n
    n= n -i

n = 3 -> 2.1 = > 2

n = 10 -> 36

[0,1,1,2,4,6,9, 12, 18, 24, 36]

pair = prev-1 + prev-2
odd = prev 2 * 2


n 6 
n 7 
both are 98



[1,1]

if n == 1{

    return 1
}


*/

func getMax(a,b int)int{
    if a > b{
        return a
    }
    
    return b
}

func integerBreak(n int) int {
    memo := make(map[int]int)
    memo[1] = 1
    
    return solution(n,n,memo)
}

func solution(target, n int, memo map[int]int)int{
    if val,e := memo[n]; e{
        return val        
    }
    
    res := 0
    if n != target {
        res = n
    }
    for i := 1 ; i < n ; i++{
        newRes := i * solution(target,n-i,memo)
        res = getMax(res, newRes)
    }
    memo[n] = res
    return res
    
}

