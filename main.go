package main

import (
	"fmt"
	// "github.com/google/go-github/v32/github"
	// "flag"
	"bufio"
	// "io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const defaultZshPathUnix = "~/.zshrc"
const defaultInitVimPathUnix = "~/.config"
const goDotFilePermission = 0777

func check(e error) {
	if e != nil {
		log.Fatal(e)
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
	usr, err := user.Current()
	check(err)

	filedir := filepath.Join(usr.HomeDir, "/godot")
	check(err)

	if _, err := os.Stat(filedir); os.IsNotExist(err) {
		err = os.MkdirAll(filedir, 0777)
		check(err)
	}

	r := bufio.NewReader(os.Stdin)

	fmt.Print("do you wish to track your zshrc file(y/n)?: ")
	text, _ := r.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if text == "y" || text == "yes" {
		fmt.Print("is it located in the default location(y/n)?: ")
		text, _ := r.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if text == "y" || text == "yes" {

		}
	}
}
