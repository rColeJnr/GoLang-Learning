//hello.go

package main

import (
	"fmt" // provides funcst for text formatting and reading fortmatted text
	"os" //provides platform-independent os variables and fncs
	"strings" // provides funcs for manipulatins strs
)

func main() {
	who := "rCole!" // short var declaration. declares and inits a variable at the same time

	if len(os.Args) > 1 {
		/* so.Args[0] is "hello" or "hello.exe" */
		who = strings.Join(os.Args[1:], " ") // assignment op
	}

	fmt.Println("Hello", who)
}