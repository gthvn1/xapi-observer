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

type Span struct {
	ID string `json:"id"`
}

func main() {
	// open ./trace-sample.json
	dat, err := os.ReadFile("./trace-sample.json")
	check(err)

	// Start by infering all spans
	var raw interface{}

	err = json.Unmarshal(dat, &raw)
	check(err)

	pretty, _ := json.MarshalIndent(raw, "", "  ")
	fmt.Println(string(pretty[:400]))

	rawSpans := raw.([]interface{})
	first := rawSpans[0].(map[string]interface{})

	fmt.Printf("%#v\n", first)

	// And now our real spans
	var spans []Span

	err = json.Unmarshal(dat, &spans)
	check(err)

	fmt.Printf("%v\n", spans)  // print only values
	fmt.Printf("%+v\n", spans) // print values and fields names
	fmt.Printf("%#v\n", spans) // print a complete Go structure

}

// Sum returns the sum of two integers.
func Sum(a, b int) int {
	return a + b
}
