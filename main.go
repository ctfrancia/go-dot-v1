package main

import (
	"fmt"
	// "github.com/google/go-github/v32/github"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/certs/github")
	check(err)

	fmt.Println(dat)
}
