/**
*
* use the following extension to execute the file properly
* https://chrome.google.com/webstore/detail/history-export/lpmoaclacdaofhlijejogfldmgkdlglj
*
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"os/exec"
)

func handlerHistory(url, name string) {

	res, _ := http.Get(url)
	baseDirectory := "html_files_dump/_"
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

type chromeHistory struct {

	Id string `json:"id"`
	LastVisitTime string `json:"lastVisitTime"'`
	LastVisitTimeTimestamp float64 `json:"lastVisitTimeTimestamp"`
	Title string `json:"title"`
	TypedCount int `json:"typedCount"`
	Url string `json:"url"`
	VisitCount int `json:"visitCount"`

}

func makeDir(name string) bool {

	_, err := exec.Command("mkdir", name).Output()
	if err != nil {
		panic(err)
	}
	return true

}

func importChromeHistoryJSON(name string) (bool, []chromeHistory) {

	fmt.Println("loading the given History DumpBase ... ")
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
	fmt.Println("done!")
	return true, chromeHistoryArray

}

func main() {

	arg := os.Args[1]
	if makeDir(arg) {
		status, chis := importChromeHistoryJSON(arg)
		lenchis := len(chis)
		if status {
			for index, historyOject := range chis {
				fmt.Print("\tfetching ", index, " out of ", lenchis, " ... ")
				handlerHistory(historyOject.Url, arg)
				fmt.Println("done.!")
			}
		}
	}
}
