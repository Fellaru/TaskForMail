package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Потерянный параметр! Вы не указали имя файла.")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Can't read file:" + os.Args[1], err)
		panic(err)
	}
	// data is the file content, you can use it
	fmt.Print("Number of lines:")
	lines := strings.Count(string(data), "\n")
	fmt.Println(lines)
}
