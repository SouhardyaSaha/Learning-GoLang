package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type CustomWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)

	cs := CustomWriter{}
	io.Copy(cs, resp.Body)
}

func (CustomWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("The size of bs: ", len(bs))
	return len(bs), nil
}
