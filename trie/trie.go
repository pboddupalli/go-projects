package trie

type Trie struct {
  children [26]*Trie
  end bool
}

func TrieConstructor() Trie {
  return Trie{}
}

func (this *Trie) Insert(word string)  {
  root := this
  for _, c := range word {
    if root.children[byte(c)-'a'] == nil {
      root.children[byte(c)-'a'] = &Trie{}
    }
    root = root.children[byte(c)-'a']
  }
  root.end = true
}

func (this *Trie) Search(word string) bool {
  root := this
  for i := 0; i < len(word) && root != nil; i++ {
    root = root.children[word[i]-'a']
  }
  return root != nil && root.end
}

func (this *Trie) StartsWith(prefix string) bool {
  root := this
  for i := 0; i < len(prefix) && root != nil; i++ {
    root = root.children[prefix[i]-'a']
  }
  return root != nil
}
