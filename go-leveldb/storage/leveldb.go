package storage

import (
	"encoding/json"
	"errors"
	"fmt"

	"go-leveldb/models"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var ErrLaunchNotFound = errors.New("launch not found")

type LaunchStore struct {
	db *leveldb.DB
}

func NewLaunchStore(dbPath string) (*LaunchStore, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return &LaunchStore{db: db}, nil
}

func (ls *LaunchStore) Close() error {
	return ls.db.Close()
}

func (ls *LaunchStore) StoreLaunches(launches []models.Launch) error {
	batch := new(leveldb.Batch)
	for _, launch := range launches {
		launchData, err := json.Marshal(launch)
		if err != nil {
			return fmt.Errorf("failed to marshal launch %s: %w", launch.ID, err)
		}
		batch.Put([]byte(launch.ID), launchData)
	}

	if err := ls.db.Write(batch, nil); err != nil {
		return fmt.Errorf("failed to write batch to database: %w", err)
	}
	return nil
}

func (ls *LaunchStore) GetLaunch(id string) (*models.Launch, error) {
	data, err := ls.db.Get([]byte(id), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, ErrLaunchNotFound
		}
		return nil, fmt.Errorf("failed to get launch %s: %w", id, err)
	}

	var launch models.Launch
	if err := json.Unmarshal(data, &launch); err != nil {
		return nil, fmt.Errorf("failed to unmarshal launch %s: %w", id, err)
	}

	return &launch, nil
}

func (ls *LaunchStore) GetAllLaunches() ([]models.Launch, error) {
	var launches []models.Launch
	iter := ls.db.NewIterator(util.BytesPrefix([]byte("")), nil)
	defer iter.Release()

	for iter.Next() {
		var launch models.Launch
		if err := json.Unmarshal(iter.Value(), &launch); err != nil {
			return nil, fmt.Errorf("failed to unmarshal launch: %w", err)
		}
		launches = append(launches, launch)
	}

	if err := iter.Error(); err != nil {
		return nil, fmt.Errorf("error iterating through launches: %w", err)
	}

	return launches, nil
}
