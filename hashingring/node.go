package hashingring

type HashAble interface {
	Hash() uint32
}

// must be thread safe
type Node interface {
	HashAble
	Put(key HashAble, data interface{}) error
	Get(key HashAble) (interface{}, bool)
}
