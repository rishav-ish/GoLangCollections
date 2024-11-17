package collections

import (
	"fmt"
)

type Node struct {
	Val      any
	nextNode *Node
}

type LinkedList struct {
	HeadNode *Node
}

func (linkedList *LinkedList) AddToHead(Val any) {
	var node = Node{}
	node.Val = Val
	if linkedList.HeadNode != nil {
		node.nextNode = linkedList.HeadNode
	}
	linkedList.HeadNode = &node
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.HeadNode; node != nil; node = node.nextNode {
		fmt.Println(node.Val)
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.HeadNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) AddToEnd(val any) {
	var node *Node = &Node{}
	node.Val = val
	var lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	} else {
		linkedList.HeadNode = node
	}
}

func (linkedList *LinkedList) NodeWithValue(val any) *Node {
	var node *Node
	for node = linkedList.HeadNode; node != nil; node = node.nextNode {
		if node.Val == val {
			return node
		}
	}
	return node
}

// func (linkedList *LinkedList) AddAfter(val any, )
