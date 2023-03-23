/*
IN this program, we are creating a flag called name. It is a string pointer.
*/
package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "stranger", "your wonderful name")
var age = flag.Int("age", 1, "your wonderful age")

func main() {
	flag.Parse()
	fmt.Printf("Hello sir %s who has wondered this lands for %d weeks, and certainly more to come, Welcome to the command line world\n", *name, *age)
}
