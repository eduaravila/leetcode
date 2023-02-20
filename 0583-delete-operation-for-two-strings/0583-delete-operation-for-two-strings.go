func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func minDistance(word1 string, word2 string) int {
    similar := [][]int{}
    
    for i:=0;i <= len(word1) ;i++{        
        similar =append(similar, make([]int, len(word2)+1))
    }
    
    
    for l:= 1 ; l <= len(word1) ; l++{        
        for r:= 1 ; r <= len(word2) ; r++{
            if word1[l-1] == word2[r-1]{
                similar[l][r] = similar[l-1][r-1] +1 
            }else{
                similar[l][r] = getMax(similar[l][r-1], similar[l-1][r])
            }
        }        
    }
    
    l, r := len(word1), len(word2)
    return l + r  - 2 * similar[l][r]
}


// sea 
// eat
// 0

// ea
//eat
// 1

// a
//eat
// 1

