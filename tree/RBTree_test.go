package tree

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

type valueType int
type keyType int

func (n *valueType) SetValue(value interface{}) (interface{}, error) {
	if v, ok := value.(valueType); ok {
		temp := *n
		*n = v
		return temp, nil
	}
	return 0, errors.New("Fail to set value")
}

func (key keyType) CompareTo(obj IComparable) (int, error) {
	if k, ok := obj.(interface{}).(keyType); ok {
		re := key - k
		return int(re), nil
	}
	return 0, errors.New("the parameter can't be converted to keyType")
}

func TestInsert(t *testing.T) {
	tree := &RBTree{Size: 0}
	for i := 0; i < 10000; i++ {
		value := valueType(i + 100)
		tree.Insert(keyType(i), &value)
	}

	// use the preorder traversal to get all the nodes in the tree.
	li := list.New()
	totalCount := 0
	temp := tree.root
	for temp != nil || li.Len() != 0 {
		if temp != nil {
			li.PushBack(temp)
			totalCount++
			temp = temp.left
		} else {
			obj := li.Back()
			if obj != nil {
				li.Remove(obj)
				temp = obj.Value.(*RBTreeNode)
				temp = temp.right
			}
		}
	}
	fmt.Println("totalCount: ", totalCount)
	if totalCount != tree.Size {
		t.Error("The size is not right.")
	}
}

func TestRemove(t *testing.T) {
	tree := &RBTree{Size: 0}
	for i := 0; i < 10000; i++ {
		value := valueType(i + 100)
		tree.Insert(keyType(i), &value)
	}

	tree.Remove(keyType(109))
	if tree.GetNode(keyType(109)) != nil {
		t.Error("Remove fail")
	}
}
