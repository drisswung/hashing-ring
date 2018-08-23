package hashingring

import (
	"testing"
	"errors"
	"strconv"
)

func TestNodes(t *testing.T) {
	nodes := nodes{
		&intNode{
			hash:  333,
			store: make(map[intKey]interface{}),
		},
		&intNode{
			hash:  0,
			store: make(map[intKey]interface{}),
		},
		&intNode{
			hash:  666,
			store: make(map[intKey]interface{}),
		},
	}
	ring := NewRing(nodes)

	for i := 1; i <= 1000; i++ {
		ring.Add(intKey(i), "lala8")

	}

	ring.AddNode(&intNode{
		hash:  667,
		store: make(map[intKey]interface{}),
	})
	notHit := 0
	for i := 1; i <= 1000; i++ {
		if _, ok := ring.Get(intKey(i)); !ok {
			notHit++
		}
	}
	if notHit != 1 {
		t.Error("not hit is not 1")
	}
}

type intKey int

func (i intKey) Hash() uint32 {
	return uint32(i)
}

type intNode struct {
	hash  uint32
	store map[intKey]interface{}
}

func (s *intNode) Get(key HashAble) (interface{}, bool) {

	v, ok := key.(intKey)
	if !ok {
		return nil, false
	}

	data, ok := s.store[v]
	return data, ok
}

func (s *intNode) Put(key HashAble, data interface{}) error {

	v, ok := key.(intKey)
	if !ok {
		return errors.New("wrong type")
	}

	s.store[v] = data
	return nil
}

func (s *intNode) Hash() uint32 {
	return s.hash
}

func (s *intNode) String() string {
	return strconv.Itoa(int(s.hash))
}
