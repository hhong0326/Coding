package list_people

import (
	"fmt"
	"grpc/pb"
	"io"
)

func writePerson(w io.Writer, p *pb.Person) {

	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "  Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "  Email:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.Person_MOBILE:
			fmt.Fprint(w, "  Mobile phone #: ")
		case pb.Person_HOME:
			fmt.Fprint(w, "  Home phone #: ")
		case pb.Person_WORK:
			fmt.Fprint(w, "  Work phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}

func ListPeople(w io.Writer, book *pb.Person_AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}
