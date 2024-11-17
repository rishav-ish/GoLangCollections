package collections

import "fmt"

type DoubleNode struct {
	Val          any
	nextNode     *DoubleNode
	previousNode *DoubleNode
}

type DoubleLinkedList struct {
	headNode *DoubleNode
}

func (linedList *DoubleLinkedList) AddToHead(val any) {
	dNode := &DoubleNode{Val: val}
	if linedList.headNode != nil {
		dNode.nextNode = linedList.headNode
		linedList.headNode.previousNode = dNode
	}
	linedList.headNode = dNode
}

func (linkedList *DoubleLinkedList) LastNode() *DoubleNode {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (linkedList *DoubleLinkedList) AddToEnd(val any) {
	newNode := &DoubleNode{Val: val}
	lNode := linkedList.LastNode()

	if lNode != nil {
		lNode.nextNode = newNode
		newNode.previousNode = lNode
	} else {
		linkedList.headNode = newNode
	}
}

func (linkedList *DoubleLinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.Val)
	}
}
