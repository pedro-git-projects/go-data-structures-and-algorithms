# Data Structures and Algorithms in Go

##### Table of Contents  
- [Node](#node)

- [Linked List](#linkedlist)

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
