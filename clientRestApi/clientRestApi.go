package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func  main()  {
	resp, er := http.Get("https://jsonplaceholder.typicode.com/posts")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}