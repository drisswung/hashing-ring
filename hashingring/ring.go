package hashingring

import (
	"sync"
	"sort"
)

type nodes []Node

func (n nodes) Less(i, j int) bool {
	return n[i].Hash() < n[j].Hash()
}

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Swap(i, j int) {
	n[j], n[i] = n[i], n[j]
}

type Ring struct {
	sync.RWMutex
	nodes nodes
}

func NewRing(ns []Node) *Ring {
	var nodes nodes = ns
	sort.Sort(nodes)
	return &Ring{
		nodes: nodes,
	}
}

func (r *Ring) AddNode(n Node) {
	r.Lock()
	defer r.Unlock()
	r.nodes = append(r.nodes, n)
	sort.Sort(r.nodes)
}

func (r *Ring) Get(key HashAble) (interface{}, bool) {
	r.RLock()
	r.RUnlock()
	hash := key.Hash()
	i := sort.Search(r.nodes.Len(), func(i int) bool {
		return r.nodes[i].Hash() >= hash
	})
	if i >= r.nodes.Len() {
		i = 0
	}
	return r.nodes[i].Get(key)
}

func (r *Ring) Add(key HashAble, data interface{}) error {
	r.RLock()
	r.RUnlock()
	hash := key.Hash()
	i := sort.Search(r.nodes.Len(), func(i int) bool {
		return r.nodes[i].Hash() >= hash
	})
	if i >= r.nodes.Len() {
		i = 0
	}
	return r.nodes[i].Put(key, data)
}
