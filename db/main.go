package main

import (
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	// Connect to the database, with timeout option in case of locks
	db, err := bolt.Open("todo.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
