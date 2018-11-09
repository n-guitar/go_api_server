package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	// todo: 引数nil判定ができない
	if args == nil {
		fmt.Println("引数がありません")
		os.Exit(1)
	}
	fmt.Println("0番目", args[0])
}
