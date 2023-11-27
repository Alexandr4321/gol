package main

import "fmt"

func main()  {
	Bob:=Cats{"Oleg",20,true}
	fmt.Println(Bob)
}

type Cats struct{
	name string
	age int
	isHappy bool
}

