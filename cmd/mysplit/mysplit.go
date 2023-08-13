package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	l int
	n int
	b string
	a int
)

func init() {
	flag.IntVar(&l, "l", 1000, "-l line_count")
	flag.IntVar(&n, "n", 2, "-n chank_count")
	flag.StringVar(&b, "b", "1MB", "-b byte_count")
	flag.IntVar(&a, "a", 0, "-a saffix_length(1~13)")
}

func validateArgs(input []string) error {
	// isConflictParam
	// IMO: 何が選ばれたか覚えておかないといけないのでは
	var count int = 0
	for _, v := range input {
		if v == "-l" || v == "-n" || v == "-b" {
			count++
		}
	}
	if count > 1 {
		return errors.New("error: too many parameters")
	}

	flag.Parse()
	if flag.NArg() != 1 && flag.NArg() != 2 {
		return errors.New("error: syntax error")
	}

	// is file exist
	if _, err := os.Stat(flag.Arg(0)); err != nil {
        return errors.New("Error: Doesn't exists the directory that you specified")
    }
	return nil
}

// func ParseParam(input []string) {

// 	//flag.perseOne()みたいにしたい
// 	flag.Parse()

// 	//ここでなんのパラメータが使用されているかを判定
// 	if l != 0 {
// 		fmt.Printf("param -l : %d\n", l)
// 	}else if n != 0 {
// 		fmt.Printf("param -n : %d\n", n)
// 	}else if b != "" {
// 		fmt.Printf("param -b : %s\n", b)
// 	}
// }

func main() {
	input := os.Args[1:]
	if err := validateArgs(input); err != nil {
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}
}
