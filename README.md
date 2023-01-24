# Data Structures and Algorithms in Go

##### Table of Contents  
- [Node](#node)

- [Linked List](#linkedlist)

- [Two Pointer Node](#dnode)

- [Doubly Linked List](#dlinkedlist)

- [Item](#item)

- [Stack](#stack)

- [Queue](#queue)

## Node

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/node)

```go
package node

type Node[T any] struct {
	value T
	next  *Node[T]
}
```



## Linked List

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/linkedlist)

```go
type LinkedList[T any] struct {
	head   *node.Node[T]
	tail   *node.Node[T]
	length int
}
```

## Two Pointer Node

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/dnode)

```go
type DNode[T any] struct {
	value T
	next  *DNode[T]
	prev  *DNode[T]
}
```



## Doubly Linked List

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/dlinkedlist)

```go
type DLinkedList[T any] struct {
	head   *dnode.DNode[T]
	tail   *dnode.DNode[T]
	length int
}
```

## Item

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/item)

```go
type Item[T any] struct {
	value T
	next  *Item[T]
}
```

## Stack

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/stack)

```go
type Stack[T any] struct {
	top    *item.Item[T]
	height int
}
```

## Queue 

[Code](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/queue)

```go
type Queue[T any] struct {
	first  *item.Item[T]
	last   *item.Item[T]
	lenght int
}
```

