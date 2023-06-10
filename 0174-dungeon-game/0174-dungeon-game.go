func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func getMin(a,b int)int{
    if a < b {
        return a
    }
    return b
}


const inf = int(^uint(0)>>1)

func calculateMinimumHP(dungeon [][]int) int {
    

    
    m := len(dungeon)
    n := len(dungeon[0])
    dp := make([]int, n+1)
    prev :=[]int{}
    

    dp[n] = 1
    
    for i := 0 ; i <= n ; i++{
        prev = append(prev,inf)
    }
    
    prev[n-1] = 1
    
    for x := m -1; x >=0 ;x--{
        for y := n -1; y >=0 ;y--{
            steps := getMin(dp[y+1], prev[y]) -(dungeon[x][y])
            if steps <= 0 {
                dp[y] = 1
            }else{
                dp[y] = steps
            }
        }
        prev = dp
        dp = make([]int,n+1)
        dp[n] = inf
    }
    return prev[0]
}


/*
[[-2,-3,3],
[-5,-10,1],
[10,30,-5]]

[[-2,-3,1],
[-5,-10,4],
[10,30,-5]]



[[-2,-3,-3],
[-5,-10,6],
[10,30,-5]]

if space < 0 {
    healt+= -1 * space 
}


[-2,-5,-2],
[-7, -13,  -1],
[ 3,  33,  -6],

-5, -4, -1, -4, -6

starting = current. + min(dp(top), dp(left)) 

if space > 0 {
    
}

[[-2,-3,3],
[-5,-10,1],
[10,30,-5]]

healt = inf, current + 1




[[-3,5]]

[]


need to look ahead of the current positon 
go to max(current + col+1, row),current + col, row+1


[[-3,5]]


[[1,-2, 1],
[-5,-10,2],
[10,30,-3]]


min = 1, -2, -3

[
[1,-3,3],
[0,-2,0],
[-3,-3,-3]
]

min = 1,-2,-1


[
[7,5, 2],
[6, 11,5],
[1,  1, 6] 1
        1   
]


*/