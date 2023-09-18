package db

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

var defaultBucket = []byte("default")

// Database is a open source bolt database
type Database struct {
	db *bolt.DB
}

// NewDatabase returns instance of database
func NewDatabase(dbPath string) (db *Database, closeFunc func() error, err error) {
	boltDb, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, nil, err
	}

	db = &Database{db: boltDb}
	closeFunc = boltDb.Close

	if err := db.createDefaultBucket(); err != nil {
		closeFunc()
		return nil, nil, fmt.Errorf("creating default bucket: %w", err)
	}
	return db, closeFunc, nil
}

func (d *Database) createDefaultBucket() error {
	return d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(defaultBucket)
		return err
	})
}

// set key to value in the database
func (d *Database) SetKey(key string, value []byte) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(defaultBucket)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(key), value)
	})
}

// GetKey returns the value of the key from the database
func (d *Database) GetKey(key string) ([]byte, error) {
	var result []byte
	err := d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(defaultBucket)
		result = bucket.Get([]byte(key))
		return nil
	})

	if err == nil {
		return result, nil
	}
	return nil, err
}
