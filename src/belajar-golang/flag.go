package main

import (
	"fmt"
	"flag"
)

func main()  {
	var host *string = flag.String("host","localhost", "put your host")
	var username *string = flag.String("username","root", "put your username")
	var password *string = flag.String("password","root", "put your password")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*username)
	fmt.Println(password)
}