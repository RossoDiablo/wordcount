package main

import (
	"flag"
	"fmt"
	"strings"
)

func Wordcount() {
	flag.Parse()
	str := flag.Arg(0)
	arrStr := strings.Fields(str)
	fmt.Println(len(arrStr))
}
