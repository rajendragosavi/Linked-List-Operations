package main

import "fmt"

type Employee struct {
	Name   string
	Design string
	next   *Employee // struct pointer pointing  to the next Employee struct
}

func main() {
	raj := &Employee{Name: "raj", Design: "engineer"}
	empList := raj
	meghana := &Employee{Name: "meghana", Design: "support"}
	arunpurus := &Employee{Name: "purus", Design: "lead"}

	neha := &Employee{Name: "neha", Design: "developer"}
	adithya := &Employee{Name: "aditya", Design: "Operations"}

	// Add a particular node at the end of the list.
	empList = addNodeatEnd(meghana, empList)
	empList = addNodeatEnd(arunpurus, empList)

	// Add a node before the specified node.
	empList = insertBeforenode("meghana", neha, empList)

	//Add a node after the specified node.
	empList = insertAfterNode("meghana", adithya, empList)

	PrintList(empList)

}

// add a node at end
func addNodeatEnd(newEmp, emplist *Employee) *Employee {
	if emplist == nil {
		return emplist
	}
	for e := emplist; e != nil; e = e.next {
		if e.next == nil {
			e.next = newEmp
			return emplist
		}
	}
	return emplist
}

// Finding mid of list *********************
func findMid(emplist *Employee) *Employee {
	var fastpr, slowptr *Employee
	fmt.Println("received list is ", emplist)
	if emplist == nil {
		fmt.Println("List is empty")
	}
	fastpr, slowptr = emplist, emplist
	for fastpr.next != nil {
		fastpr = fastpr.next
		if fastpr.next != nil && slowptr.next != nil {
			slowptr = slowptr.next
			fastpr = fastpr.next
		}
	}
	return slowptr
}
// Reverse the existing linked-list
func reverseList(emplist *Employee) {
	if emplist == nil {

	}

}

// insert a new-employee , after an employee whose name has been given.
func insertBeforenode(name string, newemployee, emplist *Employee) *Employee {
	if emplist == nil {
		return emplist
	}

	e := emplist
	for e != nil {
		fmt.Println("value of e before loop is ", e)
		if e.next.Name == name {
			fmt.Println("value of e is ", *e)
			fmt.Println("value of e.next is ", *e.next)
			newemployee.next = e.next
			e.next = newemployee
			break
		}
	}
	return emplist
}

func insertAfterNode(name string, newemployee, emplist *Employee) *Employee {
	if emplist == nil {
		return emplist
	}
	e := emplist
	for e != nil {
		nextNode := e.next
		if e.Name == name {
			newemployee.next = nextNode
			e.next = newemployee
			break
		}
		e = e.next
	}
	return emplist
}

//print the linked-list
func PrintList(emplist *Employee) {
	for e := emplist; e != nil; e = e.next {
		fmt.Println(*e)
	}
}
