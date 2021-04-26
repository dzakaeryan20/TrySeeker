package main


import (
	"fmt"
	"io"
	"log"
	"os"
)

func main(){

	f, err := os.Open("text.txt")
	if(err != nil) {
		log.Fatal(err)
	}
	defer f.Close()

	// seek Start
	cut, err := f.Seek(1, io.SeekStart)
	if(err != nil) {
		log.Fatal(err)
	}
	fmt.Println("curren loc from start from", cut)
	buf := make([]byte, 2)
	n,err := f.Read(buf)
	if(err == io.EOF){
		log.Println("We Reached EOF")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))

	// SeekEnd
	cur , err := f.Seek(-4,io.SeekEnd)
	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println("current loc end fron", cur)
	buff := make([]byte, 4)
	i,err := f.Read(buff)

	if(err == io.EOF){
		log.Fatal("Eof")
	} else if(err != nil) {
		log.Fatal("Error")
	}

	fmt.Println(string(buff[:i]))


}
