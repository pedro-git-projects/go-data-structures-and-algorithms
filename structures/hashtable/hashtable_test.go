package hashtable

import (
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/structures/utils"
)

func TestAt(t *testing.T) {
	hst := New[int]()
	hst.Insert("p-bass", 1950)
	hst.Insert("j-bass", 1970)
	hst.Insert("stingray", 1987)
	hst.Insert("dingwall", 1987)

	got, err := hst.At("p-bass")
	expect := 1950

	if got != expect {
		t.Errorf("Expected %v but got %v", expect, got)
	}

	if err != nil {
		t.Error("Did not expect an error but got ", err)
	}

	got, err = hst.At("thunderbird")
	if err == nil {
		t.Error("Expected an error but got none")
	}
}

func TestKeys(t *testing.T) {
	hst := New[int]()
	hst.Insert("p-bass", 1950)
	hst.Insert("j-bass", 1970)
	hst.Insert("stingray", 1987)
	hst.Insert("dingwall", 1987)

	keys := hst.Keys()
	expect := []string{
		"p-bass", "j-bass", "stingray", "dingwall",
	}

	if !utils.CommutativeSliceComparison(keys, expect) {
		t.Errorf("Expected %v but got %v", expect, keys)
	}

}
