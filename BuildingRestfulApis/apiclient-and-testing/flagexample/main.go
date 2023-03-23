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
var girlfriend string

func init() {
	flag.StringVar(&girlfriend, "girlfriend", "Jessica", "Your dream girl, bcs you are single")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello sir %s who has wondered this lands for %d weeks, and certainly more to come, Welcome to the command line world\n", *name, *age)
	fmt.Printf("Keep going, it will all workout one day, you will find the girl and you'll live the live you dream for you and your family and %s\n", girlfriend)
}
