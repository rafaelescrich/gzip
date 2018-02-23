package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fi, err := os.Stat("file.txt")
	if err != nil {
		// Could not obtain stat, handle error
	}

	origSize := fi.Size()
	fmt.Printf("The original file is %d KB long\n", origSize/1000)

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(content); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	// fmt.Println(str)

	// TODO: Save str to file
	err = ioutil.WriteFile("compressed", []byte(str), 0644)
	check(err)

	fi, err = os.Stat("compressed")
	if err != nil {
		// Could not obtain stat, handle error
	}

	comprSize := fi.Size()
	fmt.Printf("The compressed file is %d KB long\n", comprSize/1000)

	fmt.Printf("Ratio of compression is %d percent\n", comprSize/origSize*100)

	// content, err := ioutil.ReadFile("compress.gz")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, _ := base64.StdEncoding.DecodeString(str)
	// // fmt.Println(data)
	// rdata := bytes.NewReader(data)
	// r, _ := gzip.NewReader(rdata)
	// s, _ := ioutil.ReadAll(r)
	// fmt.Println(string(s))
	fmt.Println("Finished compression")

}
