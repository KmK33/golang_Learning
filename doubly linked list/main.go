package main

import (
	"fmt"
)

type Node struct {
	prev *Node
	data int
	next *Node
}

type DoublyLinkedLIst struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedLIst) insertAtBeginning(data int) {
	newNode := Node{nil, data, nil}

	if dll.head == nil {
		dll.head = &newNode
		dll.tail = &newNode
	}

	dll.head.prev = &newNode
	newNode.next = dll.head
	dll.head = &newNode

}

func (dll *DoublyLinkedLIst) insertAtEnd(data int){
	newNode := Node{nil,data,nil}

	if dll.head == nil {
		dll.head = &newNode
		dll.tail = &newNode
	}

	newNode.prev = dll.tail
	dll.tail.next = &newNode
	dll.tail = &newNode
}

func (dll *DoublyLinkedLIst) delete(data int ){
	current := dll.head

	

	for current != nil {


		if current.data == data{
			
			if current == dll.head{
				dll.head = current.next
				if dll.head != nil {
					dll.head.prev = nil
				}
			}else if current == dll.tail{
				dll.tail = current.prev
				dll.tail.next = nil
				
			}else{
				current.prev.next = current.next
				current.next.prev = current.prev

			}
			// fmt.Println("Node delete")
			return
		}
		current = current.next
	}

	fmt.Println("Node not found ")
}

// func (dll *DoublyLinkedLIst) reverseList(){

// 	current := dll.head

// 	dll.tail.next = dll.head
// 	dll.tail.prev = nil

// 	dll.head.prev = dll.tail
// 	dll.head.next = nil

// 	dll.head = dll.tail

// 	dll.head = current


// }

func (dll *DoublyLinkedLIst) display(){
	current := dll.head

	if current == nil {
		fmt.Print("List is emptyyy")
		return
	}

	for current != nil {
		fmt.Printf("%d->",current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := DoublyLinkedLIst{}

	list.insertAtEnd(69)
	list.insertAtEnd(690)
	
	list.display()
	// list.reverseList()
	list.display()

	
}