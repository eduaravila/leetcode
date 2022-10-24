type values struct{
    children map[rune]*values
    word bool
}

type Trie struct {
    head *values
}


func Constructor() Trie {
    return Trie{
        head: &values{            
                children: make(map[rune]*values),
                    },
    }
}


func (this *Trie) Insert(word string)  {
    current := this.head
    for _, c := range word{
        if _, e:= current.children[c];!e{
            current.children[c] = &values{ children: make(map[rune]*values)}
        }
        current = current.children[c]
    }
    
    current.word= true
}


func (this *Trie) Search(word string) bool {
    current := this.head
    for _, c := range word{
        next,e := current.children[c]
        if !e{
            return false
        }
        current = next
    }
    return current.word
}


func (this *Trie) StartsWith(prefix string) bool {
    current := this.head
    for _, c := range prefix{
        next,e := current.children[c]
        if !e{
            return false
        }
        current = next
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