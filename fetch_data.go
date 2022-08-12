package main 

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(url string){
	resp,err := http.Get("http://www.google.com")

	if err!=nil{
		fmt.Fprintf(os.Stderr,"%s\n",err)
		os.Exit(1)
	}
	b,err := ioutil.ReadAll(resp.Body)

	resp.Body.Close() 

	if err!=nil{
		fmt.Fprintf(os.Stderr,"%s\n",err)
		os.Exit(1)	
	}

	fmt.Printf("%s\n",b)
}

func main(){
	for _,url:= range os.Args[1:] {
		fetch(url)
	}
}