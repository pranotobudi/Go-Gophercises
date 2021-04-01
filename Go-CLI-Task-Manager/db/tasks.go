package db

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func InitDB(path string) error {
	// db, err := bolt.Open("my.db", 0600, nil)
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close() --> it will cause error, should have one function and called when the program finished
	// defer db.Close()
	// idx := GetLastTaskIndex()
	// fmt.Printf("Index: %v", idx)
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

}
func DeleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
	return err
}
func RetriveOneTask(key int) (Task, error) {
	var ret Task
	tasks, err := RetrieveAllTask()
	for _, task := range tasks {
		if task.Key == key {
			ret = task
		}
	}
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func RetrieveAllTask() ([]Task, error) {
	var taskList []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			// fmt.Printf("%v %v \n", k, v)
			taskList = append(taskList, Task{btoi(k), string(v)})
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return taskList, nil
}

func CreateTask(task string) (int, error) {
	var idx uint64
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		i, _ := b.NextSequence()
		idx = i
		key := itob(int(i))
		value := []byte(task)
		return b.Put(key, value)
	})
	if err != nil {
		return -1, err
	}
	return int(idx), err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func btoi(b []byte) int {
	v := binary.BigEndian.Uint64(b)
	return int(v)
}
