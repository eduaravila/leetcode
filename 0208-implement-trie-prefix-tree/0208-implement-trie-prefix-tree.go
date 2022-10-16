type Node map[rune]Node

type Trie struct {
    root Node
}


func Constructor() Trie {
    return Trie{root: make(map[rune]Node)}
}


func (this *Trie) Insert(word string)  {
    current := this.root
    for _, c := range word{
        if next,e := current[c]; e{
            current = next
        }else{
            current[c] = make(map[rune]Node)
            current = current[c]
        }
    }
    current['#'] = Node{} // finish 
    
}


func (this *Trie) Search(word string) bool {
    current := this.root
    for _, c := range word{
        if _,e := current[c]; !e{
            return false
        }
        current = current[c]        
    }
    
    return current['#'] != nil
}


func (this *Trie) StartsWith(prefix string) bool {
    current := this.root
    for _, c := range prefix{
        if _,e := current[c]; !e{
            return false
        }
        current = current[c]        
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */