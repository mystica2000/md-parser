package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"log"
	"flag"
	R "rules"
	
)

func main()  {

	srcPtr := flag.String("src","","Source File Path")
	descPtr := flag.String("desc","","Destination File Path")

	flag.Parse()
	
		
	fileName := *srcPtr
	
	if len(*descPtr) < 1 {
		*descPtr = *srcPtr
	}

	
	var writerString string 
	writerString = ""

	fileReader := ReadFile(fileName) // Read file and returns byte[]
    
	fileBytes := bytes.NewReader(fileReader) // reads byte and create byte reader
	
	scan := bufio.NewScanner(fileBytes) // to scan thru every line of the given file

	for scan.Scan() {
		str := scan.Text()
		var temp = R.DidMatch(str)
		writerString = writerString + "\n" + temp
	}


	WriteFile(*descPtr,writerString)

	if err:= scan.Err(); err != nil {
		fmt.Println("Error Reading ",err)
	}
}


func ReadFile(fileName string) []byte {
	fileReader, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return fileReader
}



func WriteFile(fileName string,textToWrite string) {
	
	f, err := os.Create(fileName)

	if err!= nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(textToWrite)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("WRITE COMPLETE")

}
