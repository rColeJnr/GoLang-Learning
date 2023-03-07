//hello.go

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "rCole!"

	if len(os.Args) > 1 {
		/* so.Args[0] is "hello" or "hello.exe" */
		who = strings.Join(os.Args[1:], " ")
	}

	fmt.Println("Hello", who)
}