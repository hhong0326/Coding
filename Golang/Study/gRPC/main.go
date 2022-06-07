package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"grpc/list_people"
	"grpc/pb"

	"google.golang.org/protobuf/proto"
)

func main() {

	// fmt.Println("protobuf test")
	// p := pb.Person{
	// 	Id:    1234,
	// 	Name:  "John Doe",
	// 	Email: "jdoe@example.com",
	// 	Phones: []*pb.Person_PhoneNumber{
	// 		{Number: "555-4321", Type: pb.Person_HOME},
	// 	},
	// }

	// book := &pb.Person_AddressBook{}

	// out, err := proto.Marshal(book)
	// if err != nil {
	// 	log.Println("Failed to encode address book: ", err)
	// }
	// if err := ioutil.WriteFile("test", out, 0644); err != nil {
	// 	log.Println("Failed to write address book: ", err)
	// }

	// fmt.Println("Add Person Test")

	// if len(os.Args) != 2 {
	// 	log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	// }
	// fname := os.Args[1]

	// // Read the existing address book.
	// in, err := ioutil.ReadFile(fname)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Printf("%s: File not found.  Creating new file.\n", fname)
	// 	} else {
	// 		log.Fatalln("Error reading file:", err)
	// 	}
	// }

	// // [START marshal_proto]
	// book := &pb.Person_AddressBook{}
	// // [START_EXCLUDE]
	// if err := proto.Unmarshal(in, book); err != nil {
	// 	log.Fatalln("Failed to parse address book:", err)
	// }

	// // Add an address.
	// addr, err := add_person.PromtForAddress(os.Stdin)
	// if err != nil {
	// 	log.Fatalln("Error with address:", err)
	// }
	// book.People = append(book.People, addr)
	// // [END_EXCLUDE]

	// // Write the new address book back to disk.
	// out, err := proto.Marshal(book)
	// if err != nil {
	// 	log.Fatalln("Failed to encode address book:", err)
	// }
	// if err := ioutil.WriteFile(fname, out, 0644); err != nil {
	// 	log.Fatalln("Failed to write address book:", err)
	// }
	// // [END marshal_proto]

	fmt.Println("List People Test")

	// Main reads the entire address book from a file and prints all the
	// information inside.
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// [START unmarshal_proto]
	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.Person_AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	// [END unmarshal_proto]

	list_people.ListPeople(os.Stdout, book)
}
