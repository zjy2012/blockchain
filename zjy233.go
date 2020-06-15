package main

import (
	"fmt"
	"strconv"
)
func  main(){
	s1:="c"
	s3:=xtob(s1)
	fmt.Println(s3)
}
func  xtob (x string)string{
	base,_:=strconv.ParseInt(x,16,20)
	return strconv.FormatInt(base,2)
}
