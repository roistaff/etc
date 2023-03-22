package main

import (
        "fmt"
        "flag"
	"os"
)

func main(){
	args := os.Args
	fmt.Println("count:",len(args))
	fmt.Println(args[1])
        username := flag.String("user", "UNKNOWN", "username")
        flag.Parse()
        fmt.Println("Hello", *username,"!")
}
