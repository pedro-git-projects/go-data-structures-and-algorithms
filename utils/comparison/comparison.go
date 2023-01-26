package comparison

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

func isEmpty(obj interface{}) bool {
	if obj == nil {
		return true
	}

	value := reflect.ValueOf(obj)

	switch value.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return value.Len() == 0
	case reflect.Ptr:
		if value.IsNil() {
			return true
		}
		deref := value.Elem().Interface()
		return isEmpty(deref)
	default:
		zero := reflect.Zero(value.Type())
		return reflect.DeepEqual(obj, zero.Interface())
	}
}

func CommutativeSliceComparison[T, U any](s0 []T, s1 []U) bool {
	if isEmpty(s0) && isEmpty(s1) {
		return true
	}

	s0Val := reflect.ValueOf(s0)
	s1Val := reflect.ValueOf(s1)

	s0Len := len(s0)
	s1Len := len(s1)

	if s0Len != s1Len {
		return false
	}

	visited := make([]bool, s1Len)
	for i := 0; i < s0Len; i++ {
		el := s0Val.Index(i).Interface()
		found := false
		for j := 0; j < s1Len; j++ {
			if visited[j] {
				continue
			}
			if cmp.Equal(s1Val.Index(j).Interface(), el) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
