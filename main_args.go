package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func handler(url, name string) {

	res, _ := http.Get(url)
	baseDirectory := name + "/"
	url = strings.Replace(url, "https://", "", -1)
	url = strings.Replace(url, "http://", "", -1)
	url = strings.Replace(url, "/", "_", -1)
	url = strings.Replace(url, ".", "_", -1)
	resInByteArr, _ := ioutil.ReadAll(res.Body)

	// creating local storage html files
	filePtr, err := os.Create(baseDirectory + url + "_.html")
	if err != nil {
		panic(err)
	}
	defer filePtr.Close()
	_, e := filePtr.Write(resInByteArr)
	if e != nil {
		panic(e)
	}

}

func makeDir2(name string) bool {

	name = strings.Replace(name, ".", "_", -1)
	_, _ = exec.Command("mkdir", name).Output()
	return true

}

func main() {

	arg := os.Args[1:]
	fmt.Println(arg)
	dirName := "html_files_dump"
	makeDir2(dirName)
	for _, ar := range arg {
		handler(ar, dirName)
	}

}
