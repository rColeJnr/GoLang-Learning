package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	if len(os.Args) == 1 {
		// filepath.Base() returns the basename of the given path
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: bigdigits [-b|--bar] %s <whole-number>\n", filepath.Base(os.Args[0]))
		fmt.Println("-b --bar draw an underbar and on overbar")
		os.Exit(0)
	}
}

func main() {
	stringOfDigits := ""
	bar := false
	if os.Args[1] == "-b" || os.Args[1] == "--bar" {
		bar = true
		stringOfDigits = os.Args[2]
	} else {
		stringOfDigits = os.Args[1]
	}
	fmt.Println(stringOfDigits)
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "   "
			} else {
				log.Fatal("invaled whole number")
			}
		}
		if bar && row == 0 {
			fmt.Println(strings.Repeat("*", len(line)))
		}
		fmt.Println(line)
		if bar && row+1 == len(bigDigits[0]) {
			fmt.Println(strings.Repeat("*", len(line)))
		}
	}
}

// var declared outside of funcs or methos may not use the := op, but
// we can get the same  effect using the long declaration for (var name = value)
var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
		"   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
