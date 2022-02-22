package main

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

var bucketFakeip = []byte("fakeip")

func main() {
	db, err := bolt.Open("C:\\Users\\Administrator\\.config\\clash-dev\\cache.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketFakeip)
		if bucket == nil {
			err := fmt.Errorf("bucket %q is not found", bucketFakeip)
			fmt.Println(err.Error())
			return err
		}

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%v\n", string(k), v)
		}

		return nil
	})
}
