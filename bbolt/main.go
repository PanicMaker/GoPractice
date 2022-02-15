package main

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
)

var bucketFakeip = []byte("fakeip")

func main() {
	db, err := bolt.Open("/Users/her/.config/clash/cache1.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketFakeip)
		if bucket == nil {
			return fmt.Errorf("bucket %q is not found", bucketFakeip)
		}

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%v\n", string(k), v)
		}

		return nil
	})
}
