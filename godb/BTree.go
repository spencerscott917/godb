package main

import (
	"encoding/binary"
)

const (
	HEADER = 4
 	BTREE_PAGE_SIZE = 4096
	BTREE_MAX_KEY_SIZE = 1000
 	BTREE_MAX_VAL_SIZE = 3000
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	if node1max > BTREE_PAGE_SIZE {
		panic("node1max greater than BTREE_PAGE_SIZE")
	}
}

type BNode []byte

type BTree struct {
	root uint64
	get func(uint64) []byte //dereference pointer
	new func([]byte) uint64 //allocate new page
	del func(uint64) //deallocate page
}

func (node BNode) nodeType() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node BNode) nKeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

