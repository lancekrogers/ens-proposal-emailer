package store

import (
	"encoding/json"
	"errors"

	bolt "go.etcd.io/bbolt"
)

var ErrNoLastProcessedBlock = errors.New("No last processed block found")

// Store is an interface for interacting with the datastore
type Store interface {
	SetLastProcessedBlock(blockNumber int) error
	GetLastProcessedBlock() (int, error)
}

// BoltDBStore is an implementation of the Store interface using BoltDB
type BoltDBStore struct {
	DB *bolt.DB
}

// NewBoltDBStore initializes and returns a new BoltDBStore
func NewBoltDBStore(dbPath string) (*BoltDBStore, error) {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	// Initialize your buckets here if they don't exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ProcessedBlockBucket"))
		return err
	})

	if err != nil {
		return nil, err
	}

	return &BoltDBStore{DB: db}, nil
}

// SetLastProcessedBlock stores the last processed block number
func (s *BoltDBStore) SetLastProcessedBlock(blockNumber int) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ProcessedBlockBucket"))
		data, err := json.Marshal(blockNumber)
		if err != nil {
			return err
		}
		return b.Put([]byte("LastProcessedBlock"), data)
	})
}

// GetLastProcessedBlock retrieves the last processed block number
func (s *BoltDBStore) GetLastProcessedBlock() (int, error) {
	var blockNumber int

	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ProcessedBlockBucket"))
		data := b.Get([]byte("LastProcessedBlock"))

		if data == nil {
			return ErrNoLastProcessedBlock
		}

		return json.Unmarshal(data, &blockNumber)
	})

	if err != nil {
		return 0, err
	}

	return blockNumber, nil
}
