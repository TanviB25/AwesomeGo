package main

import (
	"fmt"
)

type Style struct {
	body_size string
	price     int
	bname     []string
}

func main() {
	myCloset := Style{
		body_size: "XL",
		price:     300,
		bname: []string{
			"Zara",
			"Forever21",
			"Nykka",
		},
	}
	fmt.Println(myCloset)
	//fmt.Printf(myCloset.price)
	//fmt.Printf(myCloset.bname[])
	fmt.Printf(myCloset.bname[2])
}
