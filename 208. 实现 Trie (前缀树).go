package main

// 记得传指针

type Node struct {
	Children map[byte]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			Children: make(map[byte]*Node),
			isEnd:    false,
		},
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for i := range len(word) {
		if _, ok := curr.Children[word[i]]; !ok {
			curr.Children[word[i]] = &Node{Children: make(map[byte]*Node)}
		}
		curr = curr.Children[word[i]]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := range len(word) {
		if _, ok := curr.Children[word[i]]; !ok {
			return false
		}
		curr = curr.Children[word[i]]
	}
	if curr.isEnd != true {
		return false
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for i := range len(prefix) {
		if _, ok := curr.Children[prefix[i]]; !ok {
			return false
		}
		curr = curr.Children[prefix[i]]
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
