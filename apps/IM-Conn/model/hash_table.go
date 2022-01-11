package model

import "fmt"

type node struct {
	value int
	next  *node
}

type hashTable struct {
	Table map[int]*node
	Size  int
}

// 哈希函数, 余数计算
func hashFunc(i, size int) int {
	return (i % size)
}

func insert(hash *hashTable, value int) int {
	index := hashFunc(value, hash.Size)
	element := node{
		value: value,
		next: hash.Table[index],
	}
	hash.Table[index] = &element
	return index
}

func traverse(hash *hashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.value)
				t = t.next
			}
		}
		fmt.Println()
	}
}


