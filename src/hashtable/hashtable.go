package hashtable

import (
	"fmt"
	"hash/fnv"
	"reflect"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/errs"
	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/hnode"
)

// TODO: Reimplement with slice for variable hashtable size
// implement size

const defaultCapacity uint64 = 1 << 10

type HashTable[T any] struct {
	capacity uint64
	table    [defaultCapacity]*hnode.HNode[T]
}

func New[T any]() *HashTable[T] {
	return &HashTable[T]{
		capacity: defaultCapacity,
		table:    *new([defaultCapacity]*hnode.HNode[T]),
	}
}

func (hst HashTable[T]) Table() [defaultCapacity]*hnode.HNode[T] {
	return hst.table
}

func (hst HashTable[T]) String() string {
	str := ""
	for i := 0; i < int(hst.capacity); i++ {
		if hst.table[i] != nil {
			tmp := hst.table[i]
			for tmp != nil {
				str += fmt.Sprintf("%d: { %v, %v }\n", i, tmp.Key(), tmp.Value())
				tmp = tmp.Next()
			}
		}
	}
	return str
}

func (hst *HashTable[T]) hash(key any) (uint64, error) {
	h := fnv.New64()
	_, err := h.Write([]byte(fmt.Sprintf("%v", key)))
	if err != nil {
		return 0, err
	}

	hashValue := h.Sum64()
	hashed := (hst.capacity - 1) & (hashValue ^ (hashValue >> 16))
	return hashed, nil
}

func (hst *HashTable[T]) Insert(key any, value T) error {
	addr, err := hst.hash(key)

	if err != nil {
		return err
	}

	n := hnode.New(key, value)
	if hst.table[addr] == nil {
		hst.table[addr] = n
	} else {
		tmp := hst.table[addr]
		for tmp.Next() != nil {
			tmp = tmp.Next()
		}
		tmp.SetNext(n)
	}
	return nil
}

func (hst HashTable[T]) At(key any) (value T, err error) {
	addr, err := hst.hash(key)

	if err != nil {
		return *new(T), err
	}

	tmp := hst.table[addr]
	for tmp != nil {
		if reflect.DeepEqual(tmp.Key(), key) {
			return tmp.Value(), nil
		}
	}
	return *new(T), errs.OutOfBounds
}

func (hst HashTable[T]) Keys() []any {
	var keys []any
	for i := 0; i < int(hst.capacity); i++ {
		tmp := hst.table[i]
		for tmp != nil {
			keys = append(keys, tmp.Key())
			tmp = tmp.Next()
		}
	}
	return keys
}
