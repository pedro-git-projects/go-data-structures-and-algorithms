# Data Structures and Algorithms in Go

##### Table of Contents  
[Node](#node)
[Linked List](#linkedlist)

<a name="node"/>
## Node
[Node](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/node)

```go
package node

type Node[T any] struct {
	value T
	next  *Node[T]
}
```


<a name="linkedlist"/>
## Linked List
[Linked List](https://github.com/pedro-git-projects/go-data-structures-and-algorithms/tree/master/src/linkedlist)

```go
type LinkedList[T any] struct {
	head   *node.Node[T]
	tail   *node.Node[T]
	length int
}
```
