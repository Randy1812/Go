package main

import (
	"fmt"
	"path"
)

func main() {
	// var speed int
	// fmt.Println(speed)

	// var dir, file string
	// dir, file = path.Split("css/main.css")
	// fmt.Println("dir : ", dir)
	// fmt.Println("file : ", file)

	_, file := path.Split("css/main.css")
	fmt.Println("file : ", file)
}
