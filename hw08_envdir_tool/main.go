package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Println("ошибка параметров")
		return
	}
	dir := args[1]
	environments, err := ReadDir(dir)
	if err != nil {
		log.Println(err)
	}

	cmd := args[2:]
	os.Exit(RunCmd(cmd, environments))
}
