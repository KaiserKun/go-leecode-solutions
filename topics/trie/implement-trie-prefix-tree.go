package trie

// link: https://leetcode.cn/problems/implement-trie-prefix-tree/
// thought: 利用空间换时间，字典树的典型实现
type TrieNode struct {
	child [26]*TrieNode
	isEnd bool
}
type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.child[idx] == nil {
			node.child[idx] = &TrieNode{}
		}
		node = node.child[idx]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.child[idx] == nil {
			return false
		}
		node = node.child[idx]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.child[idx] == nil {
			return false
		}
		node = node.child[idx]
	}
	return true
}
