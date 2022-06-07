package add_person

import (
	"bufio"
	"fmt"
	"grpc/pb"
	"io"
	"strings"
)

func PromtForAddress(r io.Reader) (*pb.Person, error) {

	p := &pb.Person{}

	rd := bufio.NewReader(r)

	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}

	fmt.Print("Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}

	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter email (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}

	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}

		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}

		pn := &pb.Person_PhoneNumber{
			Number: phone,
		}

		fmt.Print("Is this a mobile, home, of work phone? ")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}

		ptype = strings.TrimSpace(ptype)

		switch ptype {
		case "mobile":
			pn.Type = pb.Person_MOBILE
		case "home":
			pn.Type = pb.Person_HOME
		case "work":
			pn.Type = pb.Person_WORK
		default:
			fmt.Printf("Unknown phone type %q.  Using default.\n", ptype)
		}

		p.Phones = append(p.Phones, pn)
	}

	return p, nil
}
