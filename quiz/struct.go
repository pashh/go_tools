package main

import (
	"fmt"
)

type rect struct{
	width int
	height int
}

func (a *rect) area() int{
	k:=a.width*a.height
	return(k)
}


func main(){

	r:= rect{2, 3}
	fmt.Println(r.area())

	
}
