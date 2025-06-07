package utils

type ITrie interface {
	Insert()
	Suggest() []string
}
type TrieNode struct {
	children map[string]*TrieNode
	end      bool
}
type Trie struct {
	root *TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[string]*TrieNode),
		end:      false,
	}
}
func NewTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}
func (t *Trie) Insert(str string) {
	node := t.root
	for i := range str {
		char := string(str[i])
		_, found := node.children[char]
		if !found {
			node.children[char] = newTrieNode()
		}
		node = node.children[char]
	}
	node.end = true
}
func (t *Trie) collectSuggestions(node TrieNode, prefix string, suggestions *[]string) {
	if node.end {
		*suggestions = append(*suggestions, prefix)
	}
	for char, child := range node.children {
		t.collectSuggestions(*child, prefix+char, suggestions)
	}
}
func (t *Trie) Suggest(prefix string) (suggestions []string) {
	node := t.root
	for i := range prefix {
		char := string(prefix[i])
		childNode, found := node.children[char]
		if !found {
			return suggestions
		}
		node = childNode
	}
	t.collectSuggestions(*node, prefix, &suggestions)
	return suggestions
}
