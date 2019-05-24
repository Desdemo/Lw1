package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := strconv.QuoteToASCII("苏苏")
	fmt.Println(s1)
	fmt.Println(strconv.Unquote(s1))
}
