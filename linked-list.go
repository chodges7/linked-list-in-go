package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	next *node
	name string
	data int
}

type linkedList struct {
	head   *node
	length int
}

func main() {
	l := linkedList{}
	scanner := bufio.NewScanner(os.Stdin)
out:
	for {
		fmt.Printf("----- Menu -----\n")
		fmt.Printf("1) Print list\n2) Prepend node\n")
		fmt.Printf("3) Append node\n4) Delete node\n")
		fmt.Printf("5) Exit\n")
		fmt.Printf("\n#) ")
		scanner.Scan()
		intInput, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Printf("\n")
		switch intInput {
		case 1:
			l.print_list()
		case 2:
			fmt.Printf("Name of node(string): ")
			scanner.Scan()
			strInput := scanner.Text()

			fmt.Printf("Data of node(int): ")
			scanner.Scan()
			intInput, _ = strconv.ParseInt(scanner.Text(), 10, 64)

			l.prepend_node(strInput, int(intInput))

		case 3:
			fmt.Printf("Name of node(string): ")
			scanner.Scan()
			strInput := scanner.Text()

			fmt.Printf("Data of node(int): ")
			scanner.Scan()
			intInput, _ = strconv.ParseInt(scanner.Text(), 10, 64)

			l.append_node(strInput, int(intInput))
		case 4:
			fmt.Printf("Name of node(string): ")
			scanner.Scan()
			strInput := scanner.Text()

			output := l.delete_node(strInput)

			fmt.Printf("Data from node: %d\n -- press enter --", output)
			scanner.Scan()
			fmt.Printf("\n")

		case 5:
			break out

		default:
			fmt.Println("Please retry")
		}
	}
}

func (l *linkedList) append_node(nameIn string, dataIn int, current_list ...*node) {
	var cur = new(node)
	length := len(current_list)
	if length == 0 {
		cur = l.head
	} else {
		cur = current_list[0]
	}

	if cur.next == nil {
		cur.next = &node{
			next: nil,
			name: nameIn,
			data: dataIn,
		}
		l.length++
		return
	} else {
		l.append_node(nameIn, dataIn, cur.next)
	}
}

func (l *linkedList) delete_node(nameIn string, current_list ...*node) int {
	var cur = new(node)
	length := len(current_list)
	if length == 0 {
		cur = l.head
	} else {
		cur = current_list[0]
	}

	if cur.next == nil && cur.name == nameIn {
		// deleting the only node in list
		dataOut := cur.data
		l.head = nil
		l.length--
		return dataOut
	} else if cur.next == nil {
		// at end of list ∴ return error
		return -1
	} else if cur.next.name == nameIn {
		// next's mame matches ∴ delete
		dataOut := cur.next.data
		cur.next = cur.next.next
		l.length--
		return dataOut
	} else if cur.name == nameIn { // edge case, where first node is being deleted
		l.head = cur.next
		l.length--
		return cur.data
	} else {
		// go to next node
		return l.delete_node(nameIn, cur.next)
	}
}

func (l *linkedList) prepend_node(nameIn string, dataIn int) {
	temp := l.head
	l.head = &node{
		next: temp,
		name: nameIn,
		data: dataIn,
	}
	l.length++
}

func (l *linkedList) print_list(current_list ...*node) {
	scanner := bufio.NewScanner(os.Stdin)
	length := len(current_list)
	var cur = new(node)
	if length == 0 {
		cur = l.head
	} else {
		cur = current_list[0]
	}

	// Check if the list is empty
	if l.length > 0 {
		fmt.Printf("node: %s, %d\n", cur.name, cur.data)
		if cur.next != nil {
			l.print_list(cur.next)
		} else {
			fmt.Printf("List has %d nodes\n -- press enter --", l.length)
			scanner.Scan()
			return
		}
	} else {
		fmt.Printf("List has no nodes\n -- press enter --")
		scanner.Scan()
		fmt.Printf("\n")
	}
}
