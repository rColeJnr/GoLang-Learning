package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	o      float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, or %s to quit"

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+z, Enter")
	} else { // Unix-likedd
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			o := polarCoord.o * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(o)
			y := polarCoord.radius * math.Sin(o)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

const result = "Polar radius=%.02f o=%.03f -> Cartesian x=%.02f y=%.03f\n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle:  ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, o float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &o); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, o}
		coord := <-answers
		fmt.Printf(result, radius, o, coord.x, coord.y)
	}
	fmt.Println()
}
