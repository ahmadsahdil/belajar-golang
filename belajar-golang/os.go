package main

import (
	"os"
	"fmt"
)

func main()  {
	fmt.Println(os.Args)
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname: ",name)
	}else{
		fmt.Println("error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
	

}