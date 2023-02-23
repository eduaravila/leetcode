func climbStairs(n int) int {
    if n == 1{
        return 1
    }
    if n == 2{
        return 2
    }
    
    one_step, two_steps, ways := 1, 2, 0
    
    for i := 2 ; i < n ; i++{
        ways = one_step + two_steps
        one_step = two_steps
        two_steps = ways                
    }
    
    return ways
}





/*

[1,2,3,5]

1,1,1,1
2,1,1,1
1,2,1,1
1,1,2,1
1,1,1,2
*/

