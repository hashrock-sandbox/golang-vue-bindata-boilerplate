package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/skratchdot/open-golang/open"
)

//go:generate go-assets-builder -s="/static" -o bindata.go static

type Input struct {
	In string
}

type Output struct {
	Out string
}

func jsonHandleFunc(rw http.ResponseWriter, req *http.Request) {
	output := Output{"返ってくる"}
	defer func() {
		outjson, e := json.Marshal(output)
		if e != nil {
			fmt.Println(e)
		}
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(outjson))
	}()
	if req.Method != "POST" {
		output.Out = "Not post..."
		return
	}
	body, e := ioutil.ReadAll(req.Body)
	if e != nil {
		output.Out = e.Error()
		fmt.Println(e.Error())
		return
	}
	input := Input{}
	e = json.Unmarshal(body, &input)
	if e != nil {
		output.Out = e.Error()
		fmt.Println(e.Error())
		return
	}
	fmt.Printf("%#v\n", input)
}

func main() {
	http.HandleFunc("/json", jsonHandleFunc)
	http.Handle("/", http.FileServer(Assets))
	fmt.Println("http://127.0.0.1:3000/")
	open.Start("http://127.0.0.1:3000/")
	http.ListenAndServe(":3000", nil)
}
