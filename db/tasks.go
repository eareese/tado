package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		// ignoring this error because it will only error if the transaction is closed
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itobs(int(id64))
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// conversion helpers needed by BoltDB, since everything stored in it must
// be a byte slice, so our strings are fine but integers need converting.
func itobs(v int) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(v))
	return bs
}

func bstoi(bs []byte) int {
	return int(binary.BigEndian.Uint64(bs))
}
