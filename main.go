package main

import "example.com/godatastructures/collections"

func main() {
	// var linkedList collections.LinkedList
	// linkedList = collections.LinkedList{}
	// linkedList.AddToEnd("Ishita")
	// linkedList.AddToHead(1)
	// linkedList.AddToHead("Rishav")
	// linkedList.AddToEnd("Checking")
	// linkedList.IterateList()

	var dLinkedList collections.DoubleLinkedList
	dLinkedList = collections.DoubleLinkedList{}
	dLinkedList.AddToEnd("Check me out")
	dLinkedList.AddToHead(5)
	dLinkedList.AddToHead("Rishav")
	dLinkedList.AddToHead(2.66)
	dLinkedList.AddToEnd("Completed!")
	dLinkedList.AddToEnd("Test")
	dLinkedList.IterateList()
}
