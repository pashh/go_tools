package main

import (
	"fmt"
	"strings"
	"errors"
)

type PyString string

func (py PyString) Split(str string) ( string, string , error ) {
	s := strings.Split(string(py), str)
	if len(s) < 2 {
		return "" , "", errors.New("Minimum match not found")
	}
	return s[0] , s[1] , nil
}


func main() {
	var py PyString
	py = "127.0.0.1:5432"
	ip, port , err := py.Split(":")       // Python Style
	fmt.Println(ip, port, err)
}

