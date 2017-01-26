package main

import (
	. "fmt"
	"net"
)

func main(){
	// localHost:= "10.22.75.132"
	ln, err := net.Listen("udp", ":30000")
	if err != nil{
		Println("Error: ", err.Error())
	}
	Println("ln= ", ln)
	
	/* ln, err := net.Listen("tcp", localHost+":34933")
	if err != nil{
		Println("Error: ", err.Error())
	}
	Println("ln= ", ln)
	conn, err := ln.AcceptTCP()
	if err != nil{
		Println("Error: ", err.Error())
	}
	Println("conn = ", conn)*/
}