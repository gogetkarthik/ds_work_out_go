package main

import (
	"data_structure_work_out_go/pkg/user_input"
	"fmt"
)

type node struct {
	next  *node
	value int
}

func main() {

	var parent *node

	for {
		fmt.Println("inside for loop")
		option := user_input.GetLine()

		switch option {
		case "a":
			fmt.Println("inside add")
			parent = create(user_input.GetInt(), parent)
		case "d":

		case "i":
		case "p":
			print(parent)
		default:
			fmt.Println("option not defined ")

		}

	}

}

func delete(p *node, val int) bool {

	if p != nil {
		if p.next.value == val {
			p.next = p.next.next
		} else {
			delete(p.next, val)
		}
	}

	return false
}
func print(nd *node) {
	fmt.Println(nd.value)
	if nd.next != nil {
		print(nd.next)
	}

}

func create(val int, parent *node) *node {

	if parent == nil {
		fmt.Println("inside parent")
		parent = new(node)
		parent.value = val
	} else if parent.next == nil {
		fmt.Println("inside child")
		parent.next = new(node)
		parent.next.value = val
	} else {
		fmt.Println("inside loop")
		create(val, parent.next)
	}

	return parent

}
