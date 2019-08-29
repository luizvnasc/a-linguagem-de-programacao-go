package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//echo3 original
	fmt.Println(strings.Join(os.Args[1:], " "))
	//exercicio 1.1
	fmt.Println(strings.Join(os.Args, " "))
	//exercicio 1.2
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}

}
