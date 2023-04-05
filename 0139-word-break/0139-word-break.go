/*

catsandog

cats
dog
sand
and
cat


leet code

["leet","code"]

apple pen apple

["apple","pen"]


var l, r int
words map[string]bool

if len(wordDict) < 1 {
    return true
}




*/

func solution(s string, dict map[string]bool,visited map[string]bool,i int) bool{
    if len(s) < 1 || dict[s[i:]] { 
        return true
    }
    
    if visited[s[i:]]{
        return false
    }
    
    for r := i+1; r < len(s); r++{
        if dict[s[i:r]] && solution(s,dict,visited, r) {
            return true
        }
    }
    
    visited[s[i:]] = true    
    
    return false
}

func wordBreak(s string, wordDict []string) bool {
    dictHash := make(map[string]bool)
    
    if len(s)<1 {
        return true
    }
    
    for _,word := range wordDict {
        dictHash[word] = true
    }
    
    return solution(s,dictHash,make(map[string]bool),0)    
}