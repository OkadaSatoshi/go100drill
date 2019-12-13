package main

import (
	"fmt"
	"github.com/Grape-from-kai/go100drill/practice/basic"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("ERORR: 引数を指定してください")
	} else {
		dir := os.Args[1]
		task := os.Args[2]
		if dir == "basic" {
			switch task {
			case "00":
				basic.HelloWorld()
			}
		}
	}
}