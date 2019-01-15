package tree

import (
	"github.com/pkg/errors"
)

/* This is the structrue for red/black tree*/
/*
Priciples:
1, Every node must be red or black.
2, The root node must be black.
3, The red nodes can't be sequent(It means that if the node is red, its child node and parent node couldn't be red)
4, To every node, From this node to each leaf node has the same number of the black nodes.
*/

// RBTreeNode is the node of the RB tree.
// The Key of left child node of current node is less than the Key of the current node
// The Key of the current node is less than the Key of the right child node of the current node.
type RBTreeNode struct {
	Key    IComparable
	Value  IValueInterfaces
	parent *RBTreeNode // the parent of this node
	left   *RBTreeNode // the left child of this node
	right  *RBTreeNode // the right child of this node
	color  int
}

const (
	// RED means that the color of current node is red.
	RED = iota
	// BLACK means that the color of current node is black.
	BLACK
)

// RBTree is the tree for RB tree.
type RBTree struct {
	root *RBTreeNode
	Size int
}

// IComparable is for the Nodes will be sorted by the Key , the type of Key should implement this interface
// eg. obj1.CompareTo(obj2) , if return value greater than 0, the obj1 is greater than ob2.
// if return value equals 0, the obj1 equals obj2.
// if return value less than 0, the obj1 less than obj2.
type IComparable interface {
	CompareTo(obj IComparable) (int, error)
}

// IValueInterfaces : The Value of the RBTreeNode must implement this interface.
// This interface is to set the value of the RBTreeNode
type IValueInterfaces interface {
	SetValue(value interface{}) (interface{}, error)
}

// rotateLeft is to rotate the node left.
// make this node as the left child of its right child.
func (tree *RBTree) rotateLeft(node *RBTreeNode) {
	if node == nil {
		return
	}
	r := node.right
	node.right = r.left
	if r.left != nil {
		r.left.parent = node
	}

	r.parent = node.parent
	if node.parent == nil {
		tree.root = r
	} else {
		if node.parent.left == node {
			node.parent.left = r
		} else {
			node.parent.right = r
		}
	}
	node.parent = r
	r.left = node
}

// rotateRight is to rotate the node right.
// make this node as the right child of its left child.
func (tree *RBTree) rotateRight(node *RBTreeNode) {
	if node == nil {
		return
	}
	l := node.left
	node.left = l.right
	if l.right != nil {
		l.right.parent = node
	}

	l.parent = node.parent
	if node.parent == nil {
		tree.root = l
	} else {
		if node.parent.left == node {
			node.parent.left = l
		} else {
			node.parent.right = l
		}
	}
	node.parent = l
	l.right = node
}

// GetNode is to get the node that the Key of the Node is equal to the specified key.
// If not found, return nil.
func (tree *RBTree) GetNode(key IComparable) *RBTreeNode {
	if key == nil {
		panic(errors.New("The key is nil"))
	}

	target := tree.root
	for target != nil {
		re, err := key.CompareTo(target.Key)
		if err != nil {
			return nil
		}
		if re < 0 {
			target = target.left
		} else if re > 0 {
			target = target.right
		} else {
			return target
		}
	}

	return nil
}

// Insert : Insert the (key, value) to the RBTree.
func (tree *RBTree) Insert(key IComparable, value IValueInterfaces) {
	if key == nil {
		panic(errors.New("The key is nil"))
	}
	target := tree.root
	re := 0
	p := target
	for target != nil {
		p = target
		var err error = nil
		re, err = key.CompareTo(target.Key)
		if err != nil {
			panic(err)
		}
		if re < 0 {
			target = target.left
		} else if re > 0 {
			target = target.right
		} else {
			target.Value.SetValue(value)
			return
		}
	}
	insertNode := RBTreeNode{
		Key:    key,
		Value:  value,
		parent: p,
		left:   nil,
		right:  nil,
		color:  RED,
	}
	if re < 0 {
		p.left = &insertNode
	} else if re > 0 {
		p.right = &insertNode
	} else {
		tree.root = &insertNode
	}

	tree.fixAfterInsertion(&insertNode)
	tree.Size++
}

func (tree *RBTree) fixAfterInsertion(node *RBTreeNode) {
	node.color = RED
	for node != nil && node != tree.root && node.parent.color == RED {
		if parentOf(node) == leftOf(parentOf(parentOf(node))) {
			y := rightOf(parentOf(parentOf(node)))
			if colorOf(y) == RED {
				setColor(parentOf(node), BLACK)
				setColor(y, BLACK)
				setColor(parentOf(parentOf(node)), RED)
				node = parentOf(parentOf(node))
			} else if colorOf(y) == BLACK {
				if node == rightOf(parentOf(node)) {
					node = parentOf(node)
					tree.rotateLeft(node)
				}
				setColor(parentOf(node), BLACK)
				setColor(parentOf(parentOf(node)), RED)
				tree.rotateRight(parentOf(parentOf(node)))
			}
		} else if parentOf(node) == rightOf(parentOf(parentOf(node))) {
			y := leftOf(parentOf(parentOf(node)))
			if colorOf(y) == RED {
				setColor(parentOf(node), BLACK)
				setColor(y, BLACK)
				setColor(parentOf(parentOf(node)), RED)
				node = parentOf(parentOf(node))
			} else if colorOf(y) == BLACK {
				if node == leftOf(parentOf(node)) {
					node = parentOf(node)
					tree.rotateRight(node)
				}
				setColor(parentOf(node), BLACK)
				setColor(parentOf(parentOf(node)), RED)
				tree.rotateLeft(parentOf(parentOf(node)))
			}
		}
	}
	tree.root.color = BLACK
}

func parentOf(node *RBTreeNode) *RBTreeNode {
	if node == nil {
		return nil
	}
	return node.parent
}

func colorOf(node *RBTreeNode) int {
	if node == nil {
		return BLACK
	}
	return node.color
}

func setColor(node *RBTreeNode, color int) {
	if node == nil {
		return
	}
	node.color = color
}

func rightOf(node *RBTreeNode) *RBTreeNode {
	if node == nil {
		return nil
	}
	return node.right
}

func leftOf(node *RBTreeNode) *RBTreeNode {
	if node == nil {
		return nil
	}
	return node.left
}

// successor : find the least key of the right child node of the current node.
func (tree *RBTree) successor(node *RBTreeNode) *RBTreeNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		target := node.right
		for target.left != nil {
			target = target.left
		}
		return target
	}
	return node
}

// Remove is to delete the node of the tree , the key of the removed node is equal to the parameter - key.
func (tree *RBTree) Remove(key IComparable) {
	node := tree.GetNode(key)
	if node == nil {
		return
	}
	tree.Size--

	// make sure the node that will be deleted only has less than one child tree.
	if node.left != nil && node.right != nil {
		s := tree.successor(node)
		node.Key = s.Key
		node.Value = s.Value
		node = s
	}

	// the replaceNode will replace the node that will be deleted.
	replaceNode := func() *RBTreeNode {
		if node.left != nil {
			return node.left
		}
		return node.right
	}()

	if replaceNode != nil {
		replaceNode.parent = node.parent
		if node.parent == nil {
			tree.root = replaceNode
		} else if node == node.parent.left {
			node.parent.left = replaceNode
		} else if node == node.parent.right {
			node.parent.right = replaceNode
		}
		node.parent = nil
		node.left = nil
		node.right = nil
		if node.color == BLACK {
			tree.fixAfterDeletion(replaceNode)
		}
	} else if node.parent == nil {
		tree.root = nil
	} else {
		if node == node.parent.left {
			node.parent.left = nil
		} else if node == node.parent.right {
			node.parent.right = nil
		}
		target := node.parent
		node.parent = nil
		if node.color == BLACK {
			tree.fixAfterDeletion(target)
		}
	}
}

func (tree *RBTree) fixAfterDeletion(node *RBTreeNode) {
	for node != tree.root && colorOf(node) == BLACK {
		if node == leftOf(parentOf(node)) {
			sib := rightOf(parentOf(node))
			if colorOf(sib) == RED {
				setColor(sib, BLACK)
				setColor(parentOf(node), RED)
				tree.rotateLeft(parentOf(node))
				sib = rightOf(parentOf(node))
			}

			if colorOf(leftOf(sib)) == BLACK && colorOf(rightOf(sib)) == BLACK {
				setColor(sib, RED)
				node = parentOf(node)
			} else {
				if colorOf(rightOf(sib)) == BLACK {
					setColor(leftOf(sib), BLACK)
					setColor(sib, RED)
					tree.rotateRight(sib)
					sib = rightOf(parentOf(node))
				}
				setColor(sib, colorOf(parentOf(node)))
				setColor(parentOf(node), BLACK)
				setColor(rightOf(sib), BLACK)
				tree.rotateLeft(parentOf(node))
				node = tree.root
			}
		} else if node == rightOf(parentOf(node)) {
			sib := leftOf(parentOf(node))
			if colorOf(sib) == RED {
				setColor(sib, BLACK)
				setColor(parentOf(node), RED)
				tree.rotateRight(parentOf(node))
				sib = leftOf(parentOf(node))
			}

			if colorOf(rightOf(sib)) == BLACK && colorOf(leftOf(sib)) == BLACK {
				setColor(sib, RED)
				node = parentOf(node)
			} else {
				if colorOf(leftOf(sib)) == BLACK {
					setColor(rightOf(sib), BLACK)
					setColor(sib, RED)
					tree.rotateLeft(sib)
					sib = leftOf(parentOf(node))
				}
				setColor(sib, colorOf(parentOf(node)))
				setColor(parentOf(node), BLACK)
				setColor(leftOf(sib), BLACK)
				tree.rotateRight(parentOf(node))
				node = tree.root
			}
		}
	}
	setColor(node, BLACK)
}
