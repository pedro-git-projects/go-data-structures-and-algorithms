package binsrchtree_test

import (
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/binsrchtree"
)

func TestInsert(t *testing.T) {
	/*
			  47
			/   \
		   21    76
		   / \   / \
		  18 27 52  82
	*/
	bst := binsrchtree.New(47)
	bst.Insert(21)
	bst.Insert(76)
	bst.Insert(18)
	bst.Insert(52)
	bst.Insert(82)
	bst.Insert(27)

	got := bst.Root().Left().Right().Value()
	expect := 27

	if got != expect {
		t.Errorf("Expected %d but bot %d", expect, got)
	}

	err := bst.Insert(21)
	if err == nil {
		t.Error("Expected an error but got none")
	}
}

func TestContains(t *testing.T) {
	bst := binsrchtree.New(47)
	bst.Insert(21)
	bst.Insert(76)
	bst.Insert(18)
	bst.Insert(52)
	bst.Insert(82)
	bst.Insert(27)

	got := bst.Contains(52)
	expect := true
	if got != expect {
		t.Errorf("Expected %t but bot %t", expect, got)
	}

	got = bst.Contains(29)
	expect = false
	if got != expect {
		t.Errorf("Expected %t but bot %t", expect, got)
	}
}
