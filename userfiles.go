package userfiles

import (
	"os"
	"os/user"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func InitUser() {
	usr, err := user.Current()
	check(err)

	godotFile := filepath.Join(usr.HomeDir, "/godot")
	check(err)

	if _, err := os.Stat(godotFile); os.IsNotExist(err) {
		err = os.MkdirAll(godotFile, 0777)
		check(err)

		f, err := os.Create(filepath.Join(godotFile, "config.json"))
		check(err)
		defer f.Close()
	}
}
