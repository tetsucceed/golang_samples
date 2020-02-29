package utils

import (
	"fmt"
	"sort"
)

type Contact struct {
	phone []int
	name  string
}

type Contacts []Contact

type Sort interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (p Contacts) Len() int {
	return len(p)
}

func (p Contacts) Less(i, j int) bool {
	if p[i].name < p[j].name {
		return true
	}
	return false
}

func (p Contacts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (b Contacts) ViewList() {
	for i, contact := range b {
		fmt.Printf("\t %v) %v \n", i, contact.name)
		for _, phone := range b[i].phone {
			fmt.Printf("\t -) %v \n", phone)
		}
	}
}

func NewContact(name string, ph []int) Contacts {
	var pb = make(Contacts, 1)
	pb[0] = Contact{name: name, phone: ph}
	return pb
}

func (p *Contacts) AddContact(name string, s []int) {
	*p = append(*p, Contact{name: name, phone: s})
}

func AddressBook() {

	var pb = NewContact("Никита", []int{12383, 129327})
	pb.AddContact("Вова", []int{12323483, 123129327})
	pb.ViewList()
	fmt.Printf("-------------\n")
	sort.Sort(pb)
	pb.ViewList()

}
