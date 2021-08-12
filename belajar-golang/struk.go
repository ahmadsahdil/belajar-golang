package main

import(
	"fmt"
)

type User struct{
	ID int
	nama string
	TanggalLahir string
}

func main(){
	user := User{}
	user.ID = 1

	fmt.Println(user.ID)

}