

type Trie []Trie

func Constructor() Trie {
    return make([]Trie, ('z'+1) - 'a'+1)
}


func (this *Trie) Insert(word string)  {
    current := *this
    for _, c := range word{
        index :=   c- 'a'
        if current[index] == nil{
            children := Constructor()
            current[index] = children
        }
        current = current[index]
    }
    current[('z'+1)-'a'] = Trie{} 
}


func (this *Trie) Search(word string) bool {
    current := *this
    for _, c := range word{
        index :=   c- 'a'
        next := current[index]
        if next ==nil {
            return false
        }
        current = next
    }
    return current[('z'+1)-'a'] != nil
}


func (this *Trie) StartsWith(prefix string) bool {
    current := *this
    for _, c := range prefix{
        index :=   c- 'a'
        next := current[index]
        if next == nil {
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