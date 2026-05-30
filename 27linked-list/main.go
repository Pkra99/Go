package linkedlist

import "fmt"

type Node struct {
	prev *Node 
	value int
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

type DoublyLinkedListInterface interface{
	insert(val int)
	insertValueAtHead(val int)
	insertValueAtPosition(val int, pos int)
	deleteValue(val int)
	deleteHead()
	deleteTail()
	length() int
	print()
	search(*Node)
}

func (dll *DoublyLinkedList) insert(val int) {
	newNode := &Node{value: val}

	if dll.head ==nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = dll.tail.next
}

func (dll *DoublyLinkedList) insertValueAtHead(val int) {
	newNode := &Node{value: val}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return 
	}

	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
}

func (dll *DoublyLinkedList) insertValueAtPosition(val int, pos int) {
	// handle inserting at the head
	if dll.head == nil || pos <= 1 {
		dll.insertValueAtHead(val)
		return
	}

	newNode := &Node{value: val}

	current := dll.head
	currentPos := 1

	// traverse till find the intersection
	for current.next != nil && currentPos < pos - 1  {
		current = current.next
		currentPos++
	}

	// handle inserting at the end
	if current.next == nil {
		newNode.prev = current
		current.next = newNode
		dll.tail = newNode
		return
	}

	// handle insert in the middle between current and current.next
	newNode.next = current.next
	newNode.prev = current
	current.next.prev = newNode
	current.next = newNode
}

func (dll *DoublyLinkedList) deleteHead() {
	
	if dll.head == nil {
		return
	}

	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return
	}
	
	dll.head = dll.head.next
	dll.head.prev = nil
}

func (dll *DoublyLinkedList) deleteValue(val int) {
	
	if dll.head == nil {
		return
	}

	if val == dll.head.value {
		dll.deleteHead()
		return
	}

	current := dll.head

	for current.next != nil {
		if current.next.value == val {
			target := current.next
			current.next = target.next

			if target.next != nil {
				target.next.prev = current
			} else {
				dll.tail = current
			}
			return
		}
		current = current.next
	}

}

func (dll *DoublyLinkedList) deleteTail() {
	if dll.tail == nil {
		return
	}

	if dll.tail == dll.head {
		dll.head = nil
		dll.tail  = nil
	}
	dll.tail = dll.tail.prev
	dll.tail.next = nil
}

func (dll *DoublyLinkedList) length() int {
	count := 0
	current := dll.head

	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (dll *DoublyLinkedList) print() {
	if dll.head == nil {
		fmt.Print("List is empty")
		return
	}

	current := dll.head

	fmt.Printf("nil")
	for current != nil {
		fmt.Printf( "%v <->", current.value)
		current = current.next
	}
	fmt.Printf("nil")
}

func (dll *DoublyLinkedList) search(val int) *Node{
	current := dll.head

	for current != nil {
		if current.value == val {
			return current
		}
		current = current.next
	}
	return nil
}
