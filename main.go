package main

import (
	"fmt"
	"html/template"
)

// type: defines a custom type
// struct: allows a set of related values to be grouped together
type Rsvp struct {
	Name, Email, Phone string
	WillAttend bool
}

/** function make: initialize a new slice (array with changeable size) 
	(0, 10) -> initial size and initial capacity
	[]*Rsvp => specifies the data type of the slice, in this case a slice of pointers to Rsvp structs
	Pointer is a memory address: here we don't create copies of the values when I add them to the slice
*/
var responses = make([]*Rsvp, 0, 10)

/** the key type is string
	the value type is *template.Template, which means
	a pointer to the Template struct which is inside the template package
*/
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	// syntax used within functions -> name := value (in this case, array of 5 strings)
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	// index of array, name receives the value at the current position
	for index, name := range templateNames {
		// ParseFiles returns a pointer and an error (t and err)
		t, err := template.ParseFiles("layout.html", name + ".html")
		// if there's no error, add a key-value pair to the map
		if (err == nil) {
			templates[name] = t
			fmt.Println("loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func main() {
	loadTemplates()
}