package main

import "encoding/binary"

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

type BNode []byte

// Header
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

// child pointers
func (node BNode) getPointer(idx uint16) uint16 {
	if idx >= node.nKeys() {
		panic("Attempted to get pointer outside of node")
	}
	pos := HEADER_SIZE + 8*idx
	return binary.LittleEndian.Uint16(node[pos:])
}

func (node BNode) setPointer(idx uint16, val uint64) {
	if idx >= node.nKeys() {
		panic("Attempted to get pointer outside of node")
	}
	pos := HEADER_SIZE + 8*idx
	binary.LittleEndian.PutUint16(node[pos:], val)
}

// kv store
func offsetPos(node BNode, idx uint16) uint16 {
	if idx < 1 || idx > node.nKeys() {
		panic("OffsetPos outside of node bounds")
	}
	return HEADER_SIZE + 8*node.nKeys() + 2*(idx-1)
}

func (node BNode) getOffset(idx) uint16 {
	if idx == 0 {
		return 0
	}
	return binary.LittleEndian.Uint16(node[offsetPos(node, idx):])
}
