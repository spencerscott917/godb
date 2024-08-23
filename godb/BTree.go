package main

const HEADER = 4
const BTREE_PAGE_SIZE = 16384
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 15000

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	if node1max > BTREE_PAGE_SIZE {
		panic("node1max greater than BTREE_PAGE_SIZE")
	}
}
