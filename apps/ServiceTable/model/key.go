package model

import (
	"github.com/mangenotwork/extras/common/logger"
)

var Key = NewKeyTrie()

type KeyNode struct {
	children map[rune]*KeyNode
	end bool
	keyType string
}

func newKeyNode() *KeyNode {
	return &KeyNode{
		children: make(map[rune]*KeyNode),
		end: false,
	}
}

type KeyTrie struct {
	root *KeyNode
	count int
}

func NewKeyTrie() *KeyTrie {
	return &KeyTrie{
		root: newKeyNode(),
		count: 0,
	}
}

// Insert 添加key
// keyType: key对应的数据类型
func (t *KeyTrie) Insert(key, keyType string) {
	runeChars := []rune(key)
	node := t.root
	for i := 0; i < len(runeChars); i++ {
		_, ok := node.children[runeChars[i]]
		if !ok {
			node.children[runeChars[i]] = newKeyNode()
		}
		node = node.children[runeChars[i]]
	}
	node.end = true
	node.keyType = keyType
	t.count++
}

// Has 是否存在
func (t *KeyTrie) Has(key string) bool {
	runeChars := []rune(key)
	node := t.root
	for i := 0; i < len(runeChars); i++ {
		_, ok := node.children[runeChars[i]]
		if !ok {
			return false
		}
		node = node.children[runeChars[i]]
	}
	return node.end
}

// Get 获取key
func (t *KeyTrie) Get(key string) (bool, *keyData) {
	runeChars := []rune(key)
	node := t.root
	for i := 0; i < len(runeChars); i++ {
		_, ok := node.children[runeChars[i]]
		if !ok {
			return false, nil
		}
		node = node.children[runeChars[i]]
	}
	if node.end {
		return true, &keyData{
			Key: key,
			KeyType: node.keyType,
		}
	}
	return false, nil
}

// Like 模糊查询key
func (t *KeyTrie) Like(txt string) []*keyData {
	runeChars := []rune(txt)
	rse := make([]*keyData, 0)
	if ok,v := t.Get(txt); ok {
		rse = append(rse, v)
	}

	node := t.root
	for i := 0; i < len(runeChars); i++ {
		_, ok := node.children[runeChars[i]]
		if !ok {
			return rse
		}
		node = node.children[runeChars[i]]
	}
	logger.Info("输入值: ", string(runeChars))
	likeGet(runeChars, node, &rse)
	logger.Info("rse = ", rse)
	return rse
}

func likeGet(runeChars []rune, node *KeyNode, rse *[]*keyData) {
	if node != nil {
		for k, v := range node.children {
			temp := runeChars
			temp = append(temp, k)
			if v.end {
				logger.Info("keyname = ", string(temp), "  | type = ", v.keyType)
				*rse = append(*rse, &keyData{
					Key: string(temp),
					KeyType: v.keyType,
				})
			}
			likeGet(temp, v, rse)
		}
	}
}

// GetAll 获取所有的key
func (t *KeyTrie) GetAll() []*keyData {
	node := t.root
	rse := make([]*keyData, 0)
	likeGet([]rune(""), node, &rse)
	logger.Info("all = ", rse)
	return rse
}

// StartsWith
func (t *KeyTrie) StartsWith(prefix []rune) bool {
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

// StartsWithRune
func (t *KeyTrie) StartsWithRune(prefix rune) bool {
	node := t.root
	_, ok := node.children[prefix]
	if !ok {
		return false
	}
	node = node.children[prefix]
	return true
}

func remove(pNode *KeyNode, runeChars []rune, index int) {
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
				remove(node, runeChars, index+1)
				if !node.end && len(node.children) == 0 {
					delete(pNode.children, runeChars[index])
				}
			}
		}
		index++
	}
}

// Remove
func (t *KeyTrie) Remove(key string) {
	runeChars := []rune(key)
	remove(t.root, runeChars, 0)
	t.count--
}

// key的数据属性
type keyData struct {
	Key string `json:"key"`
	KeyType string `json:"type"`
}

