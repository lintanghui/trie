package trie

import (
	"errors"
	"fmt"
)

var ws []string

type Trie struct {
	size int64
	root *Node
}

type Node struct {
	meta   rune
	level  int8 //
	parent *Node
	child  map[rune]*Node
	isEnd  bool
}

func New() *Trie {
	return &Trie{
		root: &Node{
			child: make(map[rune]*Node, 0),
		},
	}
}

// Add add word s into trie
func (t *Trie) Add(s string, level int8) {
	root := t.root
	rs := []rune(s)
	for _, r := range rs {
		if _, ok := root.child[r]; ok {
			root = root.child[r]
		} else {
			root.child[r] = &Node{
				meta:   r,
				parent: root,
				child:  make(map[rune]*Node, 0),
			}
			root = root.child[r]
		}
	}
	root.isEnd = true
	root.level = level
}

// Del del wors s from trie
func (t *Trie) Del(s string) error {
	root := t.root
	rs := []rune(s)
	for _, r := range rs {
		if _, ok := root.child[r]; ok {
			root = root.child[r]
		} else {
			return errors.New("word no exist")
		}
	}
	if !root.isEnd {
		return errors.New("word no exist")
	}
	t.delNode(root, false)
	return nil
}

func (t *Trie) delNode(n *Node, isParent bool) {
	if n.isEnd && isParent {
		return
	} else if n.isEnd && len(n.child) == 0 {
		parent := n.parent
		delete(parent.child, n.meta)
		//	fmt.Printf("del %s\n", string(n.meta))
		t.delNode(parent, true)
	} else if n.isEnd && len(n.child) != 0 {
		n.isEnd = false
	} else if !n.isEnd && len(n.child) == 0 {
		parent := n.parent
		delete(parent.child, n.meta)
		//		fmt.Printf("del %s\n", string(n.meta))
		t.delNode(parent, true)
	}
}

// Find return s if exist in trie
func (t *Trie) Find(s string) bool {
	root := t.root
	rs := []rune(s)
	for _, r := range rs {
		if _, ok := root.child[r]; ok {
			root = root.child[r]
		} else {
			return false
		}
	}
	return root.isEnd
}

// PrefixFind return all words with prefix 's'
func (t *Trie) PrefixFind(s string) (prs []string) {
	root := t.root
	rs := []rune(s)
	for _, r := range rs {
		if _, ok := root.child[r]; ok {
			root = root.child[r]
		} else {
			return
		}
	}
	ws = ws[:0]
	t.words(root, s)
	prs = make([]string, len(ws))
	copy(prs, ws)
	ws = ws[:0]
	return prs
}

// Words return all word in trie
func (t *Trie) Words() (w []string) {
	ws = ws[:0]
	t.words(t.root, "")
	w = make([]string, len(ws))
	copy(w, ws)
	ws = ws[:0]
	return
}

func (t *Trie) Filter(in string) (out string, level int8) {
	root := t.root
	rs := []rune(in)
	var offset int
	for i := 0; i < len(rs); {
		if r, ok := root.child[rs[offset+i]]; ok {
			root = r
			offset += 1
			if root.isEnd {
				for t := i; t < i+offset; t++ {
					rs[t] = rune('*')
				}
				if root.level > level {
					level = root.level
				}
				i = i + offset
				offset = 0
			}
		} else {
			root = t.root
			i++
			offset = 0
		}
	}
	out = string(rs)
	return
}

// String output all word in trie
func (t *Trie) words(n *Node, s string) {
	if n.isEnd == true {
		ws = append(ws, fmt.Sprintf("%s", s))
	}
	for r, ns := range n.child {
		rs := fmt.Sprintf("%s%s", s, string(r))
		t.words(ns, rs)
	}
}
