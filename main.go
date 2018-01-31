package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/eareese/todo/canvas"
	"github.com/eareese/todo/cmd"
	"github.com/eareese/todo/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "todo.db")

	// set up database connection
	must(db.Init(dbPath))

	// set up API checker
	must(canvas.Init())

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
