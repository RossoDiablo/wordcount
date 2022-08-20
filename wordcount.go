package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		count++
	}
	fmt.Print(count)
}