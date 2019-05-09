package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func handler(url string) {

	res, _ := http.Get(url)
	baseDirectory := "_"
	url = strings.Replace(url, "https://", "", -1)
	url = strings.Replace(url, "http://", "", -1)
	resInByteArr, _ := ioutil.ReadAll(res.Body)
	resInString := string(resInByteArr)
	fmt.Print(resInString)

	// creating local storage html files
	filePtr, _ := os.Create(baseDirectory + url + "_.html")
	_, e := filePtr.Write(resInByteArr)
	if e != nil {
		panic(e)
	}
	filePtr.Close()

}

func main() {

	arg := os.Args[1:]
	fmt.Println(arg)
	for _, ar := range arg {
		handler(ar)
	}

}
