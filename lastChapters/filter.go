package main

import (
	"filepath"
	"fmt"
	"strings"
)

func main() {

}

// Create a channel for passing filenames.
/*
	*Once the output channel has been created we create a goroutine which iterates over the files
	and sends each one to the channel. When all the files
	have been sent we close the channel.
*/
func source(files []string) <-chan string {
	var out chan string = make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename // send file name
		}
		close(out)
	}()
	return out
}

// The `in` channel parameter can be both bi and unidirectional, but this declaration ensures that it may only be received from.
func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in)) // create an output channel with the buffer size the same as the in channel
	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}
