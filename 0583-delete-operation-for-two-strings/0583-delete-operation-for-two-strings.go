func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func minDistance(word1 string, word2 string) int {
    l,r:=len(word1),len(word2)
    return r + l - 2 * solution(word1 ,word2 ,l-1 ,r-1 , make(map[key]int))    
}

type key struct{
    l,r int
}

func solution(word1,word2 string, l, r int, memo map[key]int)int{
    k := key{l,r}
    if val,e := memo[k];e{
        return val
    }
    if l < 0 || r < 0{
        return 0
    }
    if word1[l] == word2[r] {
        memo[k] = 1+ solution(word1,word2,l-1,r-1,memo)
        return memo[k]
    }
    memo[k] = getMax(solution(word1,word2, l-1,r,memo), solution(word1,word2, l,r-1,memo))
    return memo[k]
}
