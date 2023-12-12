package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insertAtBeginning(data int) {
	newNode := &Node{data, ll.head}

	ll.head = newNode
}

func (ll *LinkedList) insertAtEnd(data int){
	newNode := &Node{data: data,next: nil}

	if ll.head == nil{
		ll.head = newNode
		return
	}

	current := ll.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (ll *LinkedList) deleteNode(value int){
	if ll.head.data == value {
		ll.head = ll.head.next
		return
	}

	current := ll.head

	for current.next != nil{
		if current.next.data == value{
			current.next = current.next.next
			// fmt.Println("Node deleted")
			return
		}
		current = current.next
	}

	fmt.Println("Node not found")

}

func (ll *LinkedList) deleteList(){
	ll.head = nil
}

func (ll *LinkedList) display() {
	current := ll.head

	for current != nil {
		fmt.Printf("%d->",current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) searchElement(value int){
	current := ll.head

	for current != nil{
		if current.data == value{
			fmt.Printf("Element found: '%v'",current.data)
			return
		}
		current = current.next
	}

	fmt.Println("Element not found")
}

func main() {
	list := LinkedList{}

	list.insertAtBeginning(699)
	list.insertAtBeginning(200)

	list.insertAtEnd(420)
	// list.display()

	// list.searchElement(420)

	list.deleteList()
	list.display()

	

}
