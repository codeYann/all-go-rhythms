package classic

import "fmt"

type vertice struct {
	Key   int
	left  *vertice
	right *vertice
}

type Tree struct {
	root   *vertice
	length int
}

func CreateTree() *Tree {
	return &Tree{root: nil, length: 0}
}

func (t *Tree) GetLength() int {
	return t.length
}

func (t *Tree) GetRoot() *vertice {
	return t.root
}

func inOrderRec(u *vertice) {
	if u.left != nil {
		inOrderRec(u.left)
	}
	fmt.Println(u.Key)
	if u.right != nil {
		inOrderRec(u.right)
	}
}

func preOrderRec(u *vertice) {
	fmt.Println(u.Key)
	if u.left != nil {
		preOrderRec(u.left)
	}
	if u.right != nil {
		preOrderRec(u.right)
	}
}

func postOrderRec(u *vertice) {
	if u.left != nil {
		postOrderRec(u.left)
	}
	if u.right != nil {
		postOrderRec(u.right)
	}
	fmt.Println(u.Key)
}

func (t *Tree) Inorder() {
	inOrderRec(t.root)
}

func (t *Tree) Preorder() {
	preOrderRec(t.root)
}
func (t *Tree) Postorder() {
	postOrderRec(t.root)
}

func heightRec(node *vertice) int {
	if node == nil {
		return 0
	} else {
		var left int = heightRec(node.left)
		var right int = heightRec(node.right)

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

func (t *Tree) Height() int {
	return heightRec(t.root)
}

func (t *Tree) insertRec(root, u *vertice) {
	if u.Key < root.Key {
		if root.left == nil {
			root.left = u
		} else {
			t.insertRec(root.left, u)
		}
	} else {
		if root.right == nil {
			root.right = u
		} else {
			t.insertRec(root.right, u)
		}
	}
}

func (t *Tree) Insert(key int) {
	u := &vertice{Key: key, left: nil, right: nil}
	if t.root == nil {
		t.root = u
	} else {
		t.insertRec(t.root, u)
	}
	t.length += 1
}

func removeRec(root *vertice, key int) *vertice {
	if root == nil {
		return nil
	}

	if key < root.Key {
		root.left = removeRec(root.left, key)
	} else if key > root.Key {
		root.right = removeRec(root.right, key)
	} else {
		if root.left == nil && root.right != nil {
			root = root.right
			root.right = nil
			return root
		} else if root.left != nil && root.right == nil {
			root = root.left
			root.left = nil
			return root
		} else {
			root = nil
			return root
		}

	}
	return root
}

func (t *Tree) Remove(key int) {


}
