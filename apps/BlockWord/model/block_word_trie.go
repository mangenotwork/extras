/*
	屏蔽词-前缀树实现
 */

package model

import (
	"errors"
	"github.com/mangenotwork/extras/common/logger"
	"strings"
	"time"
)

type Node struct {
	char   rune
	data   interface{} // 附加数据，为了扩展
	parent *Node
	depth  int
	child map[rune]*Node
	end   bool
}

func newNode() *Node {
	return &Node{
		child: make(map[rune]*Node),
	}
}

type Trie struct {
	root *Node
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) Add(word string) *Trie {
	word = strings.TrimSpace(word)
	node := t.root
	runes := []rune(word)
	for _, r := range runes {
		ret, ok := node.child[r]
		if !ok {
			ret = newNode()
			ret.depth = node.depth + 1
			ret.char = r
			node.child[r] = ret
		}
		node = ret
	}
	node.end = true
	return t
}

// data是word对应数据
func (t *Trie) AddWord(word string, data interface{}) *Trie {
	t.Add(word)
	t.root.data = data
	return t
}

func (t *Trie) findNode(key string) (result *Node) {
	node := t.root
	chars := []rune(key)
	for _, v := range chars {
		ret, ok := node.child[v]
		if !ok {
			return
		}
		node = ret
	}
	result = node
	return
}

func (t *Trie) collectNode(node *Node) (result []*Node) {
	if node == nil {
		return
	}

	if node.end {
		result = append(result, node)
		return
	}

	var queue []*Node
	queue = append(queue, node)
	for i := 0; i < len(queue); i++ {
		if queue[i].end {
			result = append(result, queue[i])
			continue
		}
		for _, v1 := range queue[i].child {
			queue = append(queue, v1)
		}
	}
	return
}

func (t *Trie) PrefixSearch(key string) (result []*Node) {
	node := t.findNode(key)
	if node == nil {
		return
	}
	result = t.collectNode(node)
	return
}

func (t *Trie) BlockWord(text, replace string) (result, runTime string) {
	startTime := time.Now()
	if len(replace) < 0 {
		replace = "***"
	}
	chars := []rune(text)
	if t.root == nil {
		return
	}

	var left []rune
	node := t.root
	start := 0
	i := 0
	isInvalidWord := false
	l := len(chars)

	for index:=0; index<l; index++ {
		v := chars[index]

		if WhiteWord.StartsWithRune(v) {
			i = index
			isInvalidWord = true
		}

		if isInvalidWord {
			if WhiteWord.Search(chars[i:index]) {
				if index+1 != l {
					continue
				}
				index++
			}
		}

		ret, ok := node.child[v]
		if !ok {
			left = append(left, chars[start:index+1]...)
			start = index + 1
			node = t.root
			continue
		}

		node = ret
		if ret.end {
			node = t.root
			left = append(left, []rune(replace)...)
			start = index + 1
			continue
		}

		if index+1 == l {
			left = append(left, v)
		}
	}
	result = string(left)
	runTime = time.Since(startTime).String()
	logger.Info("time : ", runTime)
	return
}

func (t *Trie) remove(pNode *Node, runeChars []rune, index int) {
	charsLen := len(runeChars)
	if index < charsLen {
		char := runeChars[index]
		if node, ok := pNode.child[char]; ok {
			if index == charsLen-1 {
				if len(node.child) > 0 {
					node.end = false
				} else {
					delete(pNode.child, runeChars[index])
				}
			} else {
				t.remove(node, runeChars, index+1)
				if !node.end && len(node.child) == 0 {
					delete(pNode.child, runeChars[index])
				}
			}
		}
		index++
	}
}

func (t *Trie) RemoveWord(key string) error {
	if len(key) == 0{
		return errors.New("add nil word.")
	}
	runeChars := []rune(key)
	t.remove(t.root, runeChars, 0)
	return nil
}

func (t *Trie) Search(txt string) bool {
	word := []rune(txt)
	node := t.root
	for i := 0; i < len(word); i++ {
		_, ok := node.child[word[i]]
		if !ok {
			return false
		}
		node = node.child[word[i]]
	}
	return node.end
}

func (t *Trie) IsHave(text string) bool {
	chars := []rune(text)
	if t.root == nil {
		return false
	}
	node := t.root
	i := 0
	isInvalidWord := false
	l := len(chars)
	for index:=0; index<l; index++ {
		v := chars[index]
		if WhiteWord.StartsWithRune(v) {
			i = index
			isInvalidWord = true
		}
		if isInvalidWord {
			if WhiteWord.Search(chars[i:index]) {
				if index+1 != l {
					continue
				}
				index++
			}
		}
		ret, ok := node.child[v]
		if !ok {
			node = t.root
			continue
		}
		node = ret
		if ret.end {
			node = t.root
			return true
		}
	}
	return false
}

func (t *Trie) BlockHaveList(text string) (rse []string) {
	rse = make([]string,0)
	chars := []rune(text)
	if t.root == nil {
		return
	}
	var left []rune
	node := t.root
	start := 0
	i := 0
	isInvalidWord := false
	l := len(chars)
	for index:=0; index<l; index++ {
		v := chars[index]
		if WhiteWord.StartsWithRune(v) {
			i = index
			isInvalidWord = true
		}
		if isInvalidWord {
			if WhiteWord.Search(chars[i:index]) {
				if index+1 != l {
					continue
				}
				index++
			}
		}
		ret, ok := node.child[v]
		if !ok {
			left = append(left, chars[start:index+1]...)
			start = index + 1
			node = t.root
			continue
		}
		node = ret
		if ret.end {
			node = t.root
			rse = append(rse, string(chars[start:index+1]))
			start = index + 1
			continue
		}
		if index+1 == l {
			left = append(left, v)
		}
	}
	return
}
