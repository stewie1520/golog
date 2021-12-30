package main

import (
	"encoding/binary"
	"log"
)

var (
	enc = binary.BigEndian
)

func main() {
	//defer func () {
	//	value := recover()
	//	fmt.Println(value)
	//}()
	//srv := server.NewHTTPServer(":8080")
	//log.Fatal(srv.ListenAndServe())
	//fmt.Println("Done")
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}