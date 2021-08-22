package main

import "fmt"

func main(){
	var months = [12]string{
		"Januari","Februari","Maret","April","Mei","Juni","Juli","Agustus","September","Oktober","November","Desember",
	}

	var slice = months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[0]="Mei lagi"
	fmt.Println(slice)

}