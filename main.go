package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	CountDown(os.Stdout)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func CountDown(out io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}
