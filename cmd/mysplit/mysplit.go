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
	splitType string = ""
)


func init() {
	flag.IntVar(&l, "l", 1000, "-l line_count")
	flag.IntVar(&n, "n", 2, "-n chank_count")
	flag.StringVar(&b, "b", "1MB", "-b byte_count")
	flag.IntVar(&a, "a", 2, "-a saffix_length(1~13)")
}

func validateArgs() error {
	flag.Parse()

	// check option conflict 　ほんとうこれでいいのか
	flag.Visit(func(f *flag.Flag) {
		if(f.Name=="l"){
			splitType+="l"
		}else if(f.Name=="n"){
			splitType+="n"
		}else if(f.Name=="b"){
			splitType+="b"
		}
	})
	if len(splitType) > 1 {
		return errors.New("error: "+ splitType[:1]+" cannot use with " + splitType[1:] )
	}else if len(splitType) == 0 {
		return errors.New("error: you must specify one of -l, -n, -b")
	}

	// check non-option argument
	if flag.NArg() != 1 && flag.NArg() != 2 {
		return errors.New("error: too many arguments")
	}

	// is file exist
	if _, err := os.Stat(flag.Arg(0)); err != nil {
        return errors.New("Error: Doesn't exists that you specified : " + flag.Arg(0))
    }
	return nil
}

func makeSpliter() any{  
	var fi FileInfo
	fi = FileInfo{flag.Arg(0), flag.Arg(1), a}
	switch splitType {
	case "l":
		s := Line{l, fi}
		return s
	case "n":
		s := Chunk{n, fi}
		return s
	case "b":
		s := Bynary{b, fi}
		return s
	}
	return nil
}

func main() {
	if err := validateArgs(); err != nil {
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}

	fi := FileInfo{flag.Arg(0), flag.Arg(1), a}
	switch splitType {
	case "l":
		s := Line{l, fi}
		s.Split()
	case "n":
		s := Chunk{n, fi}
		s.Split()
	case "b":
		s := Bynary{b, fi}
		s.Split()
	}
}
