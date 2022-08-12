package main

import (
	"fmt"
	"os"
)

var temp string

func def(){
	fmt.Println(temp)
}

func main(){
	var err error
	temp,err= os.Getwd()

	fmt.Println(temp)
	fmt.Println(err)
	def()
}