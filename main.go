package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// open ./trace-sample.json
	dat, err := os.ReadFile("./trace-sample.json")
	check(err)

	var spans []Span

	err = json.Unmarshal(dat, &spans)
	check(err)

	// fmt.Printf("%v\n", spans)  // print only values
	// fmt.Printf("%+v\n", spans) // print values and fields names
	// fmt.Printf("%#v\n", spans) // print a complete Go structure
	for i, span := range spans {
		fmt.Printf("---------- Span %d ----------\n", i)
		span.Print()
	}
}
