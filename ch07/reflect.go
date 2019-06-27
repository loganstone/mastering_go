package main

import (
	"fmt"
	"os"
	"reflect"
)

type book struct {
	Title string
	Auth  string
	Page  int
}

func (b book) compare(a book) bool {
	rb := reflect.ValueOf(&b).Elem()
	ra := reflect.ValueOf(&a).Elem()

	for i := 0; i < rb.NumField(); i++ {
		if rb.Field(i).Interface() != ra.Field(i).Interface() {
			return false
		}
	}
	return true
}

func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Printf("Type of examine: %s\n", t)

	for i := 0; i < r.NumMethod(); i++ {
		m := r.Method(i).Type()
		fmt.Println(t.Method(i).Name, "->", m)
	}
}

func main() {
	var b book
	r := reflect.New(reflect.ValueOf(&b).Type()).Elem()
	fmt.Printf("The Type of r is %s\n", reflect.TypeOf(r))

	book1 := book{"Book", "Logan", 1980}
	book2 := book{"Book", "Logan", 180}

	if book1.compare(book1) {
		fmt.Println("Equal")
	}

	if !book1.compare(book2) {
		fmt.Println("Not Equal")
	}

	var f *os.File
	printMethods(f)
}
