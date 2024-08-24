package main

import "encoding/binary"

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

type BNode []byte

func (node BNode) nodeType() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node BNode) nKeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

func (node BNode) setHeader(nodeType uint16, nKeys uint16) {
	binary.LittleEndian.PutUint16(node[0:2], nodeType)
	binary.LittleEndian.PutUint16(node[2:4], nKeys)
}

func (node BNode) getPointer(idx uint16) uint16 {
	if idx >= node.nKeys() {
		panic("Attempted to get pointer outside of node")
	}
	pos := HEADER_SIZE + 8*idx
	return binary.LittleEndian.Uint16(node[pos:])
}

// not sure how this function is supposed to work, but it's how it is in the book
func (node BNode) setPointer(idx uint16, val uint64)
