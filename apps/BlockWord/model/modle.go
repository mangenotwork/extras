package model

import (
	"errors"
	"sort"
	"sync"
)

var Words = &wordMaps{
	words : make(map[string]struct{}),
	size : 0,
}

type wordMaps struct {
	words map[string]struct{}
	mtx sync.Mutex
	size int
}

func (wmp *wordMaps) Add(key string) {
	wmp.mtx.Lock()
	defer wmp.mtx.Unlock()
	wmp.words[key] = struct{}{}
	wmp.size++
}

func (wmp *wordMaps) Get() []string {
	wmp.mtx.Lock()
	defer wmp.mtx.Unlock()
	keyList := make([]string, 0, len(wmp.words))
	for k, _ := range wmp.words {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	return keyList
}

func (wmp *wordMaps) Del(key string) {
	wmp.mtx.Lock()
	defer wmp.mtx.Unlock()
	delete(wmp.words, key)
	wmp.size--
}

func (wmp *wordMaps) IsHave(key string) (ok bool) {
	_,ok = wmp.words[key]
	return
}

type Node struct {
	Children NodeChildren
	End      bool
}

func NewTrieNode() *Node {
	return &Node{
		Children: make(NodeChildren),
		End:      false,
	}
}

type NodeChildren map[rune]*Node

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

type Trie struct {
	Root *Node
}

func (t *Trie) AddWord(key string) error {
	if len(key) == 0{
		return errors.New("add nil word.")
	}
	runeChars := []rune(key)
	node := t.Root
	for i := 0; i < len(runeChars); i++ {
		runeChar := runeChars[i]
		if _, ok := node.Children[runeChar]; !ok {
			// 不存在，初始化子节点的map
			node.Children[runeChar] = NewTrieNode()
		}
		node = node.Children[runeChar] // 迭代
	}
	node.End = true // 叶子节点
	Words.Add(key)
	return nil
}

func (t *Trie) remove(pNode *Node, runeChars []rune, index int) {
	charsLen := len(runeChars)
	if index < charsLen {
		char := runeChars[index]
		if node, ok := pNode.Children[char]; ok {
			if index == charsLen-1 { // 达到词汇的最后一个开始进行删除
				// 判断是否是节点的最后一个
				if len(node.Children) > 0 {
					// 非最后一个的，删除该词汇的闭环
					node.End = false
				} else {
					// 如果非闭环，则为全词删除
					delete(pNode.Children, runeChars[index])
				}
			} else {
				t.remove(node, runeChars, index+1)
				if !node.End && len(node.Children) == 0 {
					// 不是叶子节点并且没有子节点了，删除
					delete(pNode.Children, runeChars[index])
				}
			}
		}
		index++
	}
}

// 删除词
func (t *Trie) Remove(key string) error {
	if len(key) == 0{
		return errors.New("add nil word.")
	}
	runeChars := []rune(key)
	t.remove(t.Root, runeChars, 0)
	Words.Del(key)
	return nil
}

// 替换屏蔽词
func (t *Trie) Replace(str string, PlaceHolder string) string {
	if len(str) == 0{
		return str
	}
	if len(PlaceHolder) == 0 {
		PlaceHolder = "*"
	}

	runes := []rune(str)
	length := len(runes)
	res := make([]rune, 0)
	i := 0
	for i < length {
		node := t.Root // 每次都节点跟开始查询
		if _, ok := node.Children[runes[i]]; !ok {
			// 在词库里没有，正常字符串要加入
			res = append(res, runes[i])
			i++
			continue
		}

		var ok bool

		j := i
	Loop:
		for (!node.End || len(node.Children) > 0) && j < length {
			// 从i开始接着进行node的查询，直到字符串结束或者node结束
			node, ok = node.Children[runes[j]]
			if !ok {
				res = append(res, runes[j])
				j++
				break Loop
			}
			res = append(res, []rune(PlaceHolder)[0])
			// delete则自动不添加
			j++
		}

		if j == i {
			i++
		} else {
			i = j
		}
	}

	return string(res)
}
