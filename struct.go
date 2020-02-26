package main

import (
	"fmt"
)

type Style struct {
	bsize []string
	price int
	bname []string
}

func main() {
	myCloset := Style{
		bsize: []string{
			"XL",
			"M",
			"XS",
		},
		price: 300,
		bname: []string{
			"Zara",
			"Forever21",
			"Nykka",
		},
	}
	fmt.Println(myCloset)
	fmt.Println(myCloset.bsize[1])
	fmt.Println(myCloset.price)
	//fmt.Printf(myCloset.price)
	//fmt.Printf(myCloset.bname[])
	fmt.Printf(myCloset.bname[2])
}
