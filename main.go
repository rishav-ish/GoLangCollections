package main

import (
	"fmt"

	"example.com/godatastructures/collections"
)

func main() {
	// var linkedList collections.LinkedList
	// linkedList = collections.LinkedList{}
	// linkedList.AddToEnd("Ishita")
	// linkedList.AddToHead(1)
	// linkedList.AddToHead("Rishav")
	// linkedList.AddToEnd("Checking")
	// linkedList.IterateList()

	// fmt.Println(linkedList.HeadNode)
	// _ = moveAhead(linkedList.HeadNode)
	// fmt.Println(linkedList.HeadNode)

	// var dLinkedList collections.DoubleLinkedList
	// dLinkedList = collections.DoubleLinkedList{}
	// dLinkedList.AddToEnd("Check me out")
	// dLinkedList.AddToHead(5)
	// dLinkedList.AddToHead("Rishav")
	// dLinkedList.AddToHead(2.66)
	// dLinkedList.AddToEnd("Completed!")
	// dLinkedList.AddToEnd("Test")
	// dLinkedList.IterateList()

	var set *collections.Set
	set = &collections.Set{}
	set.New()
	set.AddElement("Rishav")
	set.AddElement(1)
	set.AddElement("1")

	// fmt.Println(set)
	// fmt.Println(set.ContainsElement(1))
	// fmt.Println(set.ContainsElement("1"))
	// fmt.Println(set.ContainsElement("Rishav"))

	anotherSet := &collections.Set{}
	anotherSet.New()
	anotherSet.AddElement(1)
	anotherSet.AddElement("Kumar")

	fmt.Println(set)
	fmt.Println(anotherSet)

	fmt.Println(set.Intersect(anotherSet))
	fmt.Println(set.Union(anotherSet))
	fmt.Println(set.Difference(anotherSet))

}
