package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func WriteLog(logdata string) {
	file, err := ioutil.TempFile("./log", "http.log")
	defer file.Close()
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}
	file.WriteString("Log:" + logdata)
	filedata, _ := ioutil.ReadFile(file.Name())
	fmt.Println(string(filedata))
}

func RawWriter(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK!"))
	fmt.Println("User has Connected!\n")
	fmt.Println("host:", req.Host)
	fmt.Println("path:", req.URL.Path)
	fmt.Println("rawpath:", req.URL.RawPath)
	fmt.Println("rawquery:", req.URL.RawQuery)
	fmt.Println(req.Method)
	body_byte, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	body := string(body_byte)
	fmt.Println(body, "\n")
	req.ParseForm()
	fmt.Println(req.PostFormValue)
	fmt.Println(req.Form, "\n")

	var logdata string
	logdata = "HOST:" + req.Host + "\nPath:" + req.URL.Path + "\nrawquery:" + req.URL.RawQuery + "\nMethod:" + req.Method + "\nBody:" + body + "\nForm:"
	WriteLog(logdata)
}

func main() {
	http.HandleFunc("/", RawWriter)
	http.ListenAndServe(":8001", nil)
}
