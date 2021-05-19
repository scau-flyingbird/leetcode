/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	isEnd bool
	children []*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{false, make([]*Trie, 26)}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	for _, ch := range word {
		c := byte(ch) - 'a'
		if node.children[c] == nil {
			node.children[c] = &Trie{false, make([]*Trie, 26)}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *Trie) SearchPrefix(word string) *Trie {
	node := this
	for _, ch := range word {
		c := byte(ch) - 'a'
		if node.children[c] == nil {
			return nil
		}else {
			node = node.children[c]
		}
	}
	return node
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.SearchPrefix(prefix)
	return node != nil
}
/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

