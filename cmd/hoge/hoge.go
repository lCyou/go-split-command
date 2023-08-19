package main

import (
	"flag"
	"fmt"
)

var (
	m string
	n int
)

func init() {
	flag.StringVar(&m, "m", "hoge", "-m write your comment")
	flag.IntVar(&n, "n", 2, "-n repeat your comment")
}

func main() {
	flag.Parse()

	for i := 0; i < n; i++ {
	fmt.Print(m)
	}
}
