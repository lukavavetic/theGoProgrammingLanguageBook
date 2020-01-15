package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//b, err := ioutil.ReadAll(resp.Body)
		//
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//
		//fmt.Printf("%s", b)

		f, err := os.OpenFile("index.html", os.O_WRONLY, os.ModeAppend)

		written, err := io.Copy(f, resp.Body)

		if err != nil {
			fmt.Println("err", err)
			os.Exit(1)
		}

		println("written", written)
		println("http status", resp.Status)
	}
}