# Data Structures and Algorithms in Go

### Table of Contents  

---

#### Data Structures

- [Linked List](#linked-list)
- [Two Pointer Node](#two-pointer-node)
- [Doubly Linked List](#doubly-linked-list)
- [Item](#item)
- [Stack](#stack)
- [Queue](#queue)
- [Binary Node](#binary-node)
- [Binary Search Tree](#binary-search-tree)
- [Hash Node](#hash-node)
- [Hash Table](#hash-table)
- [Set](#set)
- [Graph](#graph)

#### Algorithms

- [Bubble Sort](#bubble-sort)
- [Selection Sort](#selection-sort)
- [Insertion Sort](#insertion-sort)
- [Merge Sort](#merge-sort)
- [Quick Sort](#quick-sort)

	---

### Node

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/node)

```go
package node

type Node[T any] struct {
	value T
	next  *Node[T]
}
```



### Linked List

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/linkedlist)

```go
type LinkedList[T any] struct {
	head   *node.Node[T]
	tail   *node.Node[T]
	length int
}
```

### Two Pointer Node

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/dnode)

```go
type DNode[T any] struct {
	value T
	next  *DNode[T]
	prev  *DNode[T]
}
```

### Doubly Linked List

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/dlinkedlist)

```go
type DLinkedList[T any] struct {
	head   *dnode.DNode[T]
	tail   *dnode.DNode[T]
	length int
}
```

### Item

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/item)

```go
type Item[T any] struct {
	value T
	next  *Item[T]
}
```

### Stack

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/stack)

```go
type Stack[T any] struct {
	top    *item.Item[T]
	height int
}
```

### Queue 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/queue)

```go
type Queue[T any] struct {
	first  *item.Item[T]
	last   *item.Item[T]
	lenght int
}
```

### Binary Node 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/bnode)

```go
type BNode[T constraints.Ordered] struct {
	value T
	left  *BNode[T]
	right *BNode[T]
}
```
### Binary Search Tree 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/binsrchtree)

```go
type BST[T constraints.Ordered] struct {
	root *bnode.BNode[T]
}
```

### Hash Node 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/hnode)

```go
type HNode[T any] struct {
	key   any
	value T
	next  *HNode[T]
}
```
### Hash Table 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/hashtable)

```go
const defaultCapacity uint64 = 1 << 10

type HashTable[T any] struct {
	capacity uint64
	table    [defaultCapacity]*hnode.HNode[T]
}
```
### Set 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/set)

```go
type Set[T comparable] interface {
	Insert(el T) bool
	Erase(el T)
	Cardinality() int
	Clear()
	Contains(el ...T) bool
	Difference(Set[T]) Set[T]
	Equals(Set[T]) bool
	Intersection(Set[T]) Set[T]
	IsSubset(Set[T]) bool
	IsProperSubset(Set[T]) bool
	IsSuperset(Set[T]) bool
	IsProperSuperset(Set[T]) bool
	Union(Set[T]) Set[T]
	String() string
	ToSlice() []T
}
```

### Graph 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/structures/graph)

```go
type Graph[T comparable] struct {
	adjacencyList map[T]set.Set[T]
}
```

### Bubble Sort 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/blob/master/sort/bubble.go)

```go
func Bubble[T constraints.Ordered](s []T) {
	for i := len(s) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				tmp := s[j]
				s[j] = s[j+1]
				s[j+1] = tmp
			}
		}
	}
}
```

### Selection Sort 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/blob/master/sort/selection.go)

```go
func Selection[T constraints.Ordered](s []T) {
	for i := range s {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		if i != min {
			tmp := s[i]
			s[i] = s[min]
			s[min] = tmp
		}
	}
}
```

### Insertion Sort 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/blob/master/sort/insertion.go)

```go
func Insertion[T constraints.Ordered](s []T) {
	for i := range s {
		tmp := s[i]
		j := i - 1

		for j >= 0 && tmp < s[j] {
			s[j+1] = s[j]
			s[j] = tmp
			j--
		}
	}
}
```

### Merge Sort 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/blob/master/sort/merge.go)

```go
func MergeSort[T constraints.Ordered](s *[]T) {
	ls := len(*s)
	mid := ls / 2

	if ls <= 1 {
		return
	}

	s1 := (*s)[mid:]
	s2 := (*s)[:mid]

	MergeSort(&s1)
	MergeSort(&s2)
	merge(s, &s1, &s2)
}
```

### Quick Sort 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/blob/master/sort/quick.go)

```go
func Quick[T constraints.Ordered](s []T) {
	if len(s) <= 1 {
		return
	}
	quickStep(s, 0, len(s)-1)
}
```
