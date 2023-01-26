# Data Structures and Algorithms in Go

### Table of Contents  

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

### Node

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/node)

```go
package node

type Node[T any] struct {
	value T
	next  *Node[T]
}
```



### Linked List

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/linkedlist)

```go
type LinkedList[T any] struct {
	head   *node.Node[T]
	tail   *node.Node[T]
	length int
}
```

### Two Pointer Node

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/dnode)

```go
type DNode[T any] struct {
	value T
	next  *DNode[T]
	prev  *DNode[T]
}
```

### Doubly Linked List

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/dlinkedlist)

```go
type DLinkedList[T any] struct {
	head   *dnode.DNode[T]
	tail   *dnode.DNode[T]
	length int
}
```

### Item

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/item)

```go
type Item[T any] struct {
	value T
	next  *Item[T]
}
```

### Stack

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/stack)

```go
type Stack[T any] struct {
	top    *item.Item[T]
	height int
}
```

### Queue 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/queue)

```go
type Queue[T any] struct {
	first  *item.Item[T]
	last   *item.Item[T]
	lenght int
}
```

### Binary Node 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/bnode)

```go
type BNode[T constraints.Ordered] struct {
	value T
	left  *BNode[T]
	right *BNode[T]
}
```
### Binary Search Tree 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/binsrchtree)

```go
type BST[T constraints.Ordered] struct {
	root *bnode.BNode[T]
}
```

### Hash Node 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/hnode)

```go
type HNode[T any] struct {
	key   any
	value T
	next  *HNode[T]
}
```
### Hash Table 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/hashtable)

```go
const defaultCapacity uint64 = 1 << 10

type HashTable[T any] struct {
	capacity uint64
	table    [defaultCapacity]*hnode.HNode[T]
}
```


