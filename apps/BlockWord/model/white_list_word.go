package model

var WhiteWord = newWhiteWordTrie()

type whiteWordNode struct {
	children map[rune]*whiteWordNode
	end bool
}

func newWhiteWordNode() *whiteWordNode {
	return &whiteWordNode{
		children: make(map[rune]*whiteWordNode),
		end: false,
	}
}

type whiteWordTrie struct {
	root *whiteWordNode
}

func newWhiteWordTrie() *whiteWordTrie {
	return &whiteWordTrie{
		root: newWhiteWordNode(),
	}
}

func (t *whiteWordTrie) Insert(word string) {
	runeChars := []rune(word)
	node := t.root
	for i := 0; i < len(runeChars); i++ {
		_, ok := node.children[runeChars[i]]
		if !ok {
			node.children[runeChars[i]] = newWhiteWordNode()
		}
		node = node.children[runeChars[i]]
	}
	node.end = true
}

func (t *whiteWordTrie) Search(word []rune) bool {
	node := t.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			return false
		}
		node = node.children[word[i]]
	}
	return node.end
}

func (t *whiteWordTrie) StartsWith(prefix []rune) bool {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.children[prefix[i]]
		if !ok {
			return false
		}
		node = node.children[prefix[i]]
	}
	return true
}

func (t *whiteWordTrie) StartsWithRune(prefix rune) bool {
	node := t.root
	_, ok := node.children[prefix]
	if !ok {
		return false
	}
	node = node.children[prefix]
	return true
}

func (t *whiteWordTrie) remove(pNode *whiteWordNode, runeChars []rune, index int) {
	charsLen := len(runeChars)
	if index < charsLen {
		char := runeChars[index]
		if node, ok := pNode.children[char]; ok {
			if index == charsLen-1 {
				if len(node.children) > 0 {
					node.end = false
				} else {
					delete(pNode.children, runeChars[index])
				}
			} else {
				t.remove(node, runeChars, index+1)
				if !node.end && len(node.children) == 0 {
					delete(pNode.children, runeChars[index])
				}
			}
		}
		index++
	}
}

func (t *whiteWordTrie) Remove(word string) {
	runeChars := []rune(word)
	t.remove(t.root, runeChars, 0)
}

func (t *whiteWordTrie) WordList(){

}
