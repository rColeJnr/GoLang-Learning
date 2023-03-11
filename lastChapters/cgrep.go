package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

type Job struct {
	filename string
	results  chan<- Result
}

type Result struct {
	filename string
	lino     int
	line     string
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <regexp> <files>\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invalid regexp: %s\n", err)
	} else {
		grep(lineRx, commandLineFiles(os.Args[2:]))
	}
}

var workers = runtime.NumCPU()

func grep(lineRx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(filenames)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, filenames, results) // Executes in tis own goroutine
	for i := 0; i < workers; i++ {
		go doJobs(done, lineRx, jobs) // Each exeutes in its own coroutine
	}
	go awaitCompletion(done, results) //its own coroutine
	processResults(results)
}

func addJobs(jobs chan<- Job, filenames []string, results chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename, results}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, lineRx *regexp.Regexp, jobs <-chan Job) {
	for job := range jobs {
		job.Do(lineRx)
	}
	done <- struct{}{} /* When an invocation runs
	out of jobs it signiﬁes that it has ﬁnished by sending an empty struct to the done
	channel (which is declared as a send-only channel).*/

}

// this func ensures that the main coroutine waits until all the processing is done before terminating.
func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

func processResults(results <-chan Result) {
	for result := range results {
		fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
	}
}

func waitAndProcessResultsOne(done <-chan struct{}, results <-chan Result) {
	for working := workers; working > 0; {
		select { // Blocking
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
		case <-done:
			working--
		}
	}
DONE:
	for {
		select { //Nonblocking
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
		default:
			break DONE
			/*(A bare break is not sufﬁcient since that would only break out
			of the select statement.)*/
		}
	}
}

func waitAndProcessResults(timeout int, done <-chan struct{}, results <-chan Result) {
	finish := time.After(time.Duration(timeout))
	for working := workers; working > 0; {
		select {
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
		case <-finish:
			fmt.Println("timed out")
			return
		case <-done:
			working--
		}
	}
	for {
		select { // Nonblocking
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino,
				result.line)
		case <-finish:
			fmt.Println("timed out")
			return // Time's up so finish with what results there were
		default:
			return
		}
	}

}

func (job Job) Do(lineRx *regexp.Regexp) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n\r")
		if lineRx.Match(line) {
			job.results <- Result{job.filename, lino, string(line)}
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("error:%d:%s\n", lino, err)
			}
			break
		}
	}
}

func commandLineFiles(files []string) []string { return files }

func minimum(x int, ys ...int) int {
	for _, y := range ys {
		if y < x {
			x = y
		}
	}
	return x
}
