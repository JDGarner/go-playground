package searchwords

// Design a data structure that supports adding new words and searching for existing words.

// Implement the WordDictionary class:

// void addWord(word) Adds word to the data structure.

// bool search(word) Returns true if there is any string in the data
// structure that matches word or false otherwise.
// word may contain dots '.' where dots can be matched with any letter.

// Example 1:
// Input:
// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
// [[],["day"],["bay"],["may"],["say"],["day"],[".ay"],["b.."]]

// Output:
// [null, null, null, null, false, true, true, true]

// Explanation:
// WordDictionary wordDictionary = new WordDictionary();
// wordDictionary.addWord("day");
// wordDictionary.addWord("bay");
// wordDictionary.addWord("may");
// wordDictionary.search("say"); // return false
// wordDictionary.search("day"); // return true
// wordDictionary.search(".ay"); // return true
// wordDictionary.search("b.."); // return true
// Constraints:

// 1 <= word.length <= 20
// word in addWord consists of lowercase English letters.
// word in search consist of '.' or lowercase English letters.
// There will be at most 2 dots in word for search queries.
// At most 10,000 calls will be made to addWord and search.

type WordDictionary struct {
	tree *PrefixTree
}

func Constructor() WordDictionary {
	return WordDictionary{
		tree: NewPrefixTree(),
	}
}

func (w *WordDictionary) AddWord(word string) {
	w.tree.Insert(word)
}

func (w *WordDictionary) Search(word string) bool {
	return w.tree.Search(word)
}

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
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

func (t *PrefixTree) Insert(word string) {
	runes := []rune(word)
	currentNode := t.root

	for _, r := range runes {
		if currentNode.children[r] == nil {
			currentNode.children[r] = &TrieNode{
				children: map[rune]*TrieNode{},
			}
		}

		currentNode = currentNode.children[r]
	}

	currentNode.endOfWord = true
}

// Has to handle something like ".ay" or ".a."
func (t *PrefixTree) Search(word string) bool {
	// start iterating through the input
	// if a character is "." => need to search all children of current node
	// else just take the character

	runes := []rune(word)
	currentNode := t.root

	return searchHelper(0, runes, currentNode)
}

func searchHelper(currentIndex int, runes []rune, node *TrieNode) bool {
	// if we're at the last character and at endOfWord => true, else => false
	if currentIndex == len(runes) {
		return node.endOfWord
	}

	// if there are no children at all at this node then there's no possiblity for search
	// to be true
	if len(node.children) == 0 {
		return false
	}

	char := runes[currentIndex]

	if char == '.' {
		// call searchHelper with all the children of node
		// if we find anything that matches, return true
		for _, c := range node.children {
			if searchHelper(currentIndex+1, runes, c) {
				return true
			}
		}

		return false
	}

	newNode, exists := node.children[char]
	if !exists {
		return false
	}

	return searchHelper(currentIndex+1, runes, newNode)
}
