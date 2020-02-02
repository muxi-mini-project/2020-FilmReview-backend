package main

import (
"fmt"
)

type name struct {
	Xing int
	Ming int
}

func main() {
	var I = []name{{5,6},{7,8}}
	var H *[]name = &I
	fmt.Println((*H)[1])
}
