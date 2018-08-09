package trie

// Node : A Node in Trie
type Node struct {
	val    rune
	length int
	pass   int
	count  int
	next   map[rune]*Node
}

func (n *Node) allSuffix(arr *[]string, curr string) {
	curr += string(n.val)
	if len(n.next) == 0 {
		*arr = append(*arr, curr[1:])
	}
	for _, v := range n.next {
		v.allSuffix(arr, curr)
	}

}

// Prefix : Find all string in trie that have prefix s
func (t *Trie) Prefix(s string) []string {
	node := t.Find(s)
	arr := make([]string, 0)
	node.allSuffix(&arr, "")
	return arr
}

// Trie : Trie data structure
type Trie struct {
	root *Node
}

// Size : return the size of Trie
func (t *Trie) Size() int {
	return t.root.pass
}

// New : Construct a new Trie
func New() *Trie {
	return &Trie{
		root: &Node{length: 0, next: make(map[rune]*Node)},
	}
}

func (node *Node) newNode(val rune) *Node {
	new := &Node{
		val:    val,
		pass:   0,
		count:  0,
		length: node.length + 1,
		next:   make(map[rune]*Node),
	}
	node.next[val] = new
	return new
}

// Add : Add a string to Trie
func (t *Trie) Add(s string) {
	node := t.root

	for _, v := range s {
		node.pass++
		if next, ok := node.next[v]; ok {
			node = next
		} else {
			node = node.newNode(v)
		}
	}
	node.pass++
	node.count++
}

// Find : Check if Trie contain string s or not
func (t *Trie) Find(s string) *Node {
	node := t.root

	for _, v := range s {
		if next, ok := node.next[v]; ok {
			node = next
		} else {
			return nil
		}
	}
	return node
}

// FindWord :
func (t *Trie) FindWord(s string) bool {
	node := t.Find(s)
	return node.count > 0
}

// Remove : Delete string s in Trie
func (t *Trie) Remove(s string) (ok bool) {
	node := t.root

	if t.Find(s) == nil {
		return false
	}

	for _, v := range s {
		node.pass--
		node = node.next[v]
	}
	node.pass--
	node.count--
	return true
}
