/*
go set
implyment by binary search struct
*/

package main

import (
	"fmt"
)

//

type TreeNode struct {
	value  interface{}
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

//some thing strange
//to visit all binary search tree
//it will be visited start from left, then root, last right
//is weather need recursive
func (node *TreeNode) Next() *TreeNode {
	//if node's right is not nil  the next node is rightnode' leftmostnode
	if node.right != nil {
		tmpNode := node.right
		return tmpNode.leftMost()
	} else { //node right is nil, the next node is node's parent
		/*
			//	  .				//this node is node's next node
			//  .
			//	  .
			//  	. 			//tmpNode
			//		  .			//node
			//		.			//

		*/
		node1 := node
		tmpNode := node1.parent
		for node1 == tmpNode.right { //  if tmpNode node is rightnode,like next
			node1 = tmpNode
			tmpNode = node1.parent
		}
		return tmpNode
	}
}

//return the node in front of this node
func (node *TreeNode) ForHead() *TreeNode {
	//todo:
}

//if has no left node  return itsself
func (node *TreeNode) leftMost() *TreeNode {
	mleftNode := node
	for mleftNode.left != nil {
		mleftNode = mleftNode.left
	}
	return mleftNode
}

//if has no right node  return itselft
func (node *TreeNode) rightMost() *TreeNode {
	mrightNode := node
	for mrightNode.right != nil {
		mrightNode = mrightNode.right
	}
	return mrightNode
}

func (node TreeNode) Value() interface{} {
	return node.value
}

func compare(value1, value2 interface{}) bool {
	return value1 < value2
}

type Set struct {
	root *TreeNode
	size int
}

//todo: compare func is a key allow user define

//insert value to set, it be sort by itself
func (s *Set) Insert(value interface{}) {
	if s.root == nil {

	} else {
		//find where to insert
	}
}

/*
func (s *Set) getLocation(value interface{}) *TreeNode {

}
*/
//search a value isExists
func (s *Set) Find(value interface{}) bool {
	if s.root == nil {
		return false
	} else {
		node := s.root

		//node == value
		if node.value == value {
			return true
		}

		//node < value
		for compare(node.value, value) {
			node = node.Next()
			if node == nil {
				return false
			}

			if node.value == value {
				return true
			}
		}

		for !compare(node.value, value) {
			node = node.ForHead()
			if node == nil {
				return false
			}

			if node.value == value {
				return true
			}
		}
	}
}

func (s *Set) Erase(value interface{}) {

}

//do not change set, so use s
//for visiting a Set
func (s Set) Begin() *TreeNode {
	//return s.root
}

func (s Set) Size() int {
	return s.size
}

func (s Set) IsEmpty() bool {
	return s.root != nil
}

func (s *Set) Clear() {

}

func (s *Set) adjustSelf() {

}

//has to research recursive

func (s *Set) visitNode() *TreeNode {
	//return s.root.left

}

func (s Set) PrintSet() {
	if s.IsEmpty() {
		panic("set is empty")
	}

}
