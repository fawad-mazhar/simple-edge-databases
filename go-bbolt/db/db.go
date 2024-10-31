package db

import (
	"encoding/json"
	"fmt"
	"go-bbolt/models"

	"go.etcd.io/bbolt"
)

type Database struct {
	*bbolt.DB
}

// OpenDB opens the BoltDB file with the given filename
func OpenDB(filename string) (*Database, error) {
	db, err := bbolt.Open(filename, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

// CreateBucket creates a new bucket in the BoltDB
func (db *Database) CreateBucket(bucketName string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

// InsertLaunchData inserts the launch data into the specified bucket
func (db *Database) InsertLaunchData(bucketName string, launches []models.Launch) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		for _, launch := range launches {
			launchBytes, err := json.Marshal(launch)
			if err != nil {
				return err
			}
			err = b.Put([]byte(launch.ID), launchBytes)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
