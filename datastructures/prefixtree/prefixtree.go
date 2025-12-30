package prefixtree

// A prefix tree (also known as a trie) is a tree data structure used to
// efficiently store and retrieve keys in a set of strings.
// Some applications of this data structure include auto-complete
// and spell checker systems.

// A Trie is structured as a tree-like data structure where each node contains
// a hash map (or an array for fixed character sets) to store references to its
// child nodes, which represent characters.
// Each node also includes a boolean flag to indicate whether the current node
// marks the end of a valid word.
// The Trie starts with a root node that does not hold any character and serves
// as the entry point for all operations. The child nodes of the root and subsequent
// nodes represent unique characters from the words stored in the Trie,
// forming a hierarchical structure based on the prefixes of the words.

// Example 1:
// Input:
// ["Trie", "insert", "dog", "search", "dog", "search", "do", "startsWith", "do", "insert", "do", "search", "do"]

// Output:
// [null, null, true, false, true, null, true]

// Explanation:
// PrefixTree prefixTree = new PrefixTree();
// prefixTree.insert("dog");
// prefixTree.search("dog");    // return true
// prefixTree.search("do");     // return false
// prefixTree.startsWith("do"); // return true
// prefixTree.insert("do");
// prefixTree.search("do");     // return true

// Constraints:
// 1 <= word.length, prefix.length <= 1000
// word and prefix are made up of lowercase English letters.

// dog, do

//            root
//       do         hi
//  don     dog

// insert hi
//

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{
		root: &TrieNode{
			children: map[rune]*TrieNode{},
		},
	}
}

// Iterate through the characters of the word
// starting at the root of the Trie as the current node.
// If the current node already contains word[i],
// we continue to the next character and move to the node
// that word[i] points to.
// If word[i] is not present, we create a new node for word[i]
// and continue the process until we reach the end of the word.
// We mark the boolean variable as true as it is the end of the inserted word.

// Example inserting dog, then do:

// root
//   children:
//      d: &d{} endOfWord: false
//      g: &g{} endOfWord: true

// o:
// children:
//   &g{}
// endOfWord: false

// g:
// children:
// endOfWord: true

// after inserting 'do' o changed to:
//      o: endOfWord: true

// Inserts the string word into the prefix tree.
func (t *PrefixTree) Insert(word string) {
	current := t.root

	// If current node does not contain word[i], add it to the children of current
	// node, then move to that node
	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = &TrieNode{
				children: map[rune]*TrieNode{},
			}
		}

		current = current.children[char]
	}

	current.endOfWord = true
}

// returns true if the string word is in the prefix tree, false otherwise.
func (t *PrefixTree) Search(word string) bool {
	current := t.root

	for _, char := range word {
		node := current.children[char]
		if node == nil {
			return false
		}
		current = node
	}

	// if we reached here, all of the letters of word are in the tree, but
	// we have to make sure it's the end of a word
	return current.endOfWord
}

// Returns true if there is a previously
// inserted string word that has the prefix prefix, and false otherwise.
func (t *PrefixTree) StartsWith(prefix string) bool {
	current := t.root

	for _, char := range prefix {
		node := current.children[char]
		if node == nil {
			return false
		}
		current = node
	}

	return true
}
