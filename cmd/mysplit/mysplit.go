package main

import (
	// "errors"
	"flag"
	"fmt"
)

var (
	l int
	n int
	b string
)

func init() {
	flag.IntVar(&l, "l", 1000, "-l line_count")
	flag.IntVar(&n, "n", 2, "-n chank_count")
	flag.StringVar(&b, "b", "1MB", "-b byte_count")
}

func JudgeParam(){
	//解析 
	//flag.perseOne()みたいにしたい
	flag.Parse()

	//ここでなんのパラメータが使用されているかを判定
	if l != 0 {
		fmt.Printf("param -l : %d\n", l)
	}else if n != 0 {
		fmt.Printf("param -n : %d\n", n)
	}else if b != "" {
		fmt.Printf("param -b : %s\n", b)
	}
}

func main() {
	JudgeParam()
}
