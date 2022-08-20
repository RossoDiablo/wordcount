package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)
	str,_ := rdr.ReadString('\n')
	arr := strings.Split(str," ")
	fmt.Print(len(arr))
}