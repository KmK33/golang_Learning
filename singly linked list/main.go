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
	fmt.Println("Node inserted")
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
	fmt.Println("Node inserted")
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

func (ll *LinkedList) reverseList(){
	current := ll.head 

	var prev *Node =  nil
	var next *Node =  nil

	for current != nil{
		next = current.next
		current.next = prev
		prev = current
		current = next	

		
	}

	ll.head = prev
}

func main() {
	var choice uint16;

	list := LinkedList{}

	for choice != 8  {
		fmt.Println("Select the operation you want to perform:")
		fmt.Println("1.Insert node at beginning\n2.Insert node at End")
		fmt.Println("3.Delete a Node\n4.Delete the list")
		fmt.Println("5.Reverse the list\n6.Search an element")
		fmt.Println("7.Display list\n8.Exit")

		_,err := fmt.Scanln(&choice)

		if err != nil{
			fmt.Print(err)
			

		}
		

		switch choice {
		case 1:
			var data int;
			fmt.Println("Enter data to be inserted : " )
			_,err := fmt.Scanln(&data)
			if err != nil {
				fmt.Println(err)
				break
			}
			list.insertAtBeginning(data)

		case 2:
			var data int;
			fmt.Println("Enter data to be inserted : " )
			_,err := fmt.Scanln(&data)
			if err != nil {
				fmt.Println(err)
				break
			}
			list.insertAtEnd(data)

		case 3:
			var data int;
			fmt.Println("Enter data to be deleted : " )
			_,err := fmt.Scanln(&data)
			if err != nil {
				fmt.Println(err)
				break
			}
			list.deleteNode(data)

		case 4:
			list.deleteList()
			fmt.Println("List deleted")

		case 5 :
			list.reverseList()
			fmt.Println("List reversed!!")

		case 6:
			var data int;
			fmt.Println("Enter data to be searced : " )
			_,err := fmt.Scanln(&data)
			if err != nil {
				fmt.Println(err)
				break
			}
			list.searchElement(data)

		case 7:
			list.display()

		case 8:
			fmt.Println("Exitingggg")

		default:
			fmt.Println("Invalid option!!!")

			

			
		}

		
	}


	

	

}
