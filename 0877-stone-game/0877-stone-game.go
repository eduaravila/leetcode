func stoneGame(piles []int) bool {
    return solution(piles, 0,len(piles)-1, 0,0, true,make(map[key]bool))
}

type key struct {
    a,b int
}

func solution(piles []int, l,r,alice,bob int,turn bool, memo map[key]bool)bool{        
    k := key{l,r}
    if val,e := memo[k];e{
        return val
    }
    
    if l > r && alice > bob {
        return true
    }    
    if l > r && bob > alice {
        return false
    }
    
    if turn {
        res := solution(piles,l+1,r,alice+piles[l],bob, !turn,memo) || solution(piles,l,r-1,alice+piles[r],bob, !turn,memo)
        memo[k] = res
        return res
    }
    res := solution(piles,l+1,r,alice,bob + piles[l], !turn,memo) || solution(piles,l,r-1,alice,bob + piles[r], !turn,memo)
    memo[k] = res
    return res
}
/*




l,r int

player int
    alice = 0
    bob = 0
*/