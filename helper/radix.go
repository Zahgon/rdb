package helper

type radixNode struct {
	path      string
	end       bool
	children  []*radixNode
	totalSize int // total size of all key-value with this prefix
	keyCount  int
	fullpath  string
}

type radixTree struct {
	root *radixNode
}

func newRadixTree() *radixTree { _ = "STUB: not implemented"; return nil }

func commonPrefixLen(wordA, wordB string) int { _ = "STUB: not implemented"; return 0 }

func (tree *radixTree) insert(word string, size int) { _ = "STUB: not implemented"; return }

// assert: node == root || i > 0, because it is the first loop or from `for _, child := range node.children`
// split current node `rn`

// word must be a descendants of node

// assert node.fullpath == fullword

// word may have common prefix with a child, recurse until no common prefix

// now, word has no common prefix with child

func (tree *radixTree) traverse(cb func(node *radixNode, depth int) bool) {
	_ = "STUB: not implemented"
	return
}

func (n *radixNode) GetSize() int { _ = "STUB: not implemented"; return 0 }

func genKey(db int, key string) string { _ = "STUB: not implemented"; return "" }

func parseNodeKey(key string) (int, string) { _ = "STUB: not implemented"; return 0, "" }

// if key is db root, index+1 == len(key)
