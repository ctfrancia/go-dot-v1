package main

import (
	"fmt"
	// "github.com/google/go-github/v32/github"
	// "flag"
	"bufio"
	// "io/ioutil"
	"os"
	// "reflect"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type dotFiles struct {
	initVimFile      string
	hasInitVimFile   bool
	zshrc            string
	hasZshrc         bool
	bashProfile      string
	hasBashProfile   bool
	githubAcct       string
	hasGitHubAcct    bool
	bitbucketAcct    string
	hasBitBucketAcct bool
}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("do you wish to track your zshrc file(y/n)?: ")
	text, _ := r.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if text == "y" || text == "yes" {
		fmt.Print("is it located in the default location(y/n)?")
		text, _ := r.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if text == "y" || text == "yes" {

		}
	}
	/*
		fmt.Println(text)
		fmt.Print("second text: ")
		text, _ = r.ReadString('\n')
	*/

	/*
		dat, err := ioutil.ReadFile("/certs/github")
		check(err)

		fmt.Println(dat)
	*/
}
