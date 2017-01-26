package main

import (
	. "fmt"
	"net"
)

func main(){
	localHost := "10.22.75.132"
	Println("hei")	
	ln, err := net.Listen("tcp", localHost+":34933")
	if err != nil{
		Println("Error!", err.Error())
	}
	Println("ln = ", ln)
	conn, err := ln.Accept()
	if err!= nil{
		Println("Error!", err.Error())
	}
	Println("conn =", conn)
	
}
