package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/eareese/todo/cmd"
	"github.com/eareese/todo/db"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func main() {
	// looking for an env variable called TODO_TOKEN
	viper.SetEnvPrefix("todo")
	viper.BindEnv("token")
	token := viper.Get("token")
	if token == nil {
		must(errors.New("an API access token is required for Canvas functionality"))
	}

	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "todo.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
