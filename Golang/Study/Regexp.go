package main

import (
	"bytes"
	f "fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	f.Println(match)
	r, _ := regexp.Compile("p([a-z]+)ch")

	f.Println(r.MatchString("peach"))
	f.Println(r.FindString("peach punch"))
	f.Println(r.FindStringIndex("peach punch"))
	f.Println(r.FindStringSubmatch("peach punch"))
	f.Println(r.FindStringSubmatchIndex("peach punch"))
	f.Println(r.FindAllString("peach punch", -1))

	f.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	f.Println(r.FindAllString("peach punch pinch", 2))
	f.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	f.Println(r)

	f.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	f.Println(string(out))

}
