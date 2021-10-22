package classic

import (
	"errors"
	"fmt"
)

type Node_Tree struct {
	Key   int
	left  *Node_Tree
	right *Node_Tree
}

type Tree struct {
	Root   *Node_Tree
	length int
}

func CreateTree() *Tree {
	return &Tree{Root: nil, length: 0}
}

func (tree *Tree) InOrder(node *Node_Tree) {
	if node != nil {
		tree.InOrder(node.left)
		fmt.Println(node.Key)
		tree.InOrder(node.right)
	}

}

func (tree *Tree) findRecTree(node *Node_Tree, key int) *Node_Tree {
	if node == nil || node.Key == key {
		return node
	}

	if node.Key < key {
		return tree.findRecTree(node.right, key)
	}
	return tree.findRecTree(node.left, key)

}

func (tree *Tree) Find(key int) (*Node_Tree, error) {
	if tree.Root == nil {
		return nil, errors.New("tree is empty")
	}

	response := tree.findRecTree(tree.Root, key)
	if response == nil {
		return nil, errors.New(fmt.Sprint("Impossible to find", key))
	} else {
		return response, nil
	}
}

func (tree *Tree) insertRecTree(root, node *Node_Tree) {
	if node.Key < root.Key {
		if root.left == nil {
			root.left = node
		} else {
			tree.insertRecTree(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			tree.insertRecTree(root.right, node)
		}
	}
}

func (tree *Tree) Insert(key int) {
	node := &Node_Tree{Key: key, left: nil, right: nil}
	if tree.Root == nil {
		tree.Root = node
	} else {
		tree.insertRecTree(tree.Root, node)
	}
	tree.length += 1
}
