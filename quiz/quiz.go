package main

import (
	"flag"
	"fmt"
	//"strings"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "123"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}


func find_max_len(words []string,  word_tmp string) []string{
	//max_len := 0
	result := words
	for _, r := range words {
		c := string(r)
		fmt.Println(c)
	//for ( w : words ){
	//	if ( w.value <= S.value && size(w) > max_len ){
	//		result = w ; max_len = size(w)
	//	}
	}
	return result
}



var words arrayFlags


func main() {
	flag.Var(&words, "list", "list strings")
	word := flag.String("word", "", "string")
	flag.Parse()
	s:= find_max_len(words, *word)
	//s:=strings.Split(words[0], ",")
	fmt.Println(s)

	
}
