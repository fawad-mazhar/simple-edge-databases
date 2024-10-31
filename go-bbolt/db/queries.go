package db

import (
	"encoding/json"
	"fmt"
	"go-bbolt/internal/utils"
	"go-bbolt/models"

	"go.etcd.io/bbolt"
)

func (db *Database) GetAllLaunches(bucketName string) ([]models.Launch, error) {
	var launches []models.Launch

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var launch models.Launch
			err := json.Unmarshal(v, &launch)
			if err != nil {
				return err
			}
			launches = append(launches, launch)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return launches, nil
}

// GetLaunchByID retrieves a single launch by its ID from the specified bucket
func (db *Database) GetLaunchByID(bucketName, launchID string, dest interface{}) error {
	return db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		launchBytes := b.Get([]byte(launchID))
		if launchBytes == nil {
			return fmt.Errorf("launch not found: %s", launchID)
		}

		return json.Unmarshal(launchBytes, dest)
	})
}

func (db *Database) ListBuckets() []string {
	var buckets []string
	db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			buckets = append(buckets, string(name))
			return nil
		})
	})
	return buckets
}

func (db *Database) GetBucketKeys(bucketName string) []models.KeyValue {
	var pairs []models.KeyValue
	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return nil
		}

		return b.ForEach(func(k, v []byte) error {
			valueStr := string(v)
			isJSONValue := utils.IsJSON(valueStr)
			var jsonValue interface{}
			if isJSONValue {
				jsonValue = utils.FormatJSON(valueStr)
			}

			pairs = append(pairs, models.KeyValue{
				Key:       string(k),
				Value:     valueStr,
				IsJSON:    isJSONValue,
				JSONValue: jsonValue,
			})
			return nil
		})
	})
	return pairs
}

func (db *Database) GetKeyValueInJson(bucketName string, key string) map[string]interface{} {
	var value string

	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return nil
		}
		value = string(b.Get([]byte(key)))
		return nil
	})
	isJSONValue := utils.IsJSON(value)
	response := map[string]interface{}{
		"value":  value,
		"isJSON": isJSONValue,
	}
	if isJSONValue {
		response["jsonValue"] = utils.FormatJSON(value)
	}
	return response
}
