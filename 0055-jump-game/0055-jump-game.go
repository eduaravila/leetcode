func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func canJump(nums []int) bool {
    n := len(nums)
    if n < 2{
        return true
    }
    var max int
    
    for i, num := range nums {
        if i > max{
            break
        }
        max = getMax(max, i + num)    
    }
    return max >= n-1
}


/*

[2,3,1,1,4]

2 . x x . .
3   . x x x
1    . x .
1      . x
       
       
       


3 . x x x .
2   . x x
      . x 
      
[1,0,0,1,0]


[1] - true


[0,1]

[1,0]


max = 1, 3

[3,2,1,0,4]

[1,1,1,1,0]


*/