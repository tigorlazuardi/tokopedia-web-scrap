package cmd

import (
	"fmt"
	"os"
)

func Execute() int {
	fmt.Println(getArgs())
	return 0
}

func getArgs() []string {
	return os.Args[1:]
}
