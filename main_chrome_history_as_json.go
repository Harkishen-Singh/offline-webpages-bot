package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//func handlerHistory(url string) {
//
//	res, _ := http.Get(url)
//	baseDirectory := "html_files_dump/_"
//	url = strings.Replace(url, "https://", "", -1)
//	url = strings.Replace(url, "http://", "", -1)
//	resInByteArr, _ := ioutil.ReadAll(res.Body)
//	resInString := string(resInByteArr)
//	fmt.Print(resInString)
//
//	// creating local storage html files
//	filePtr, _ := os.Create(baseDirectory + url + "_.html")
//	_, e := filePtr.Write(resInByteArr)
//	if e != nil {
//		panic(e)
//	}
//	filePtr.Close()
//
//}

type chromeHistory struct {

	Id string `json:"id"`
	LastVisitTime string `json:"lastVisitTime"'`
	LastVisitTimeTimestamp float64 `json:"lastVisitTimeTimestamp"`
	Title string `json:"title"`
	TypedCount int `json:"typedCount"`
	Url string `json:"url"`
	VisitCount int `json:"visitCount"`

}

func importChromeHistoryJSON(name string) {

	var chromeHistoryArray []chromeHistory
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	fileHIS, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(fileHIS, &chromeHistoryArray)
	if err != nil {
		panic(err)
	}
	fmt.Println(chromeHistoryArray)

}

func main() {

	arg := os.Args[1]
	importChromeHistoryJSON(arg)

}
