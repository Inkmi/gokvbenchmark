package kvengines

import (
	"go.etcd.io/bbolt"
	"os"
	"path/filepath"
)

type boltStore struct {
	db *bbolt.DB
}

var boltBbucketName = []byte("test-bucket")

func NewBoltDb(path string, sync bool) Store {
	dir := filepath.Join("benchmark", path)
	opts := bbolt.DefaultOptions
	opts.NoSync = !sync
	var err error
	_ = os.MkdirAll(dir, os.ModePerm)

	dataFile := filepath.Join("benchmark", path, "bolt.data")
	boltDB, err := bbolt.Open(dataFile, 0644, opts)
	if err != nil {
		panic(err)
	}

	boltDB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucket(boltBbucketName)
		if err != nil {
			panic(err)
		}
		return nil
	})
	return &boltStore{
		db: boltDB,
	}
}

func (s *boltStore) Close() error {
	err := s.db.Close()
	return err
}

func (s *boltStore) FlushDB() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if err := tx.DeleteBucket(boltBbucketName); err != nil {
			return err
		}
		_, err := tx.CreateBucket(boltBbucketName)
		return err
	})
}

func (s *boltStore) Set(key, value []byte) error {
	s.db.Update(func(tx *bbolt.Tx) error {
		err := tx.Bucket(boltBbucketName).Put(key, value)
		if err != nil {
			panic(err)
		}
		return nil
	})
	return nil
}

func (s *boltStore) Get(key []byte) ([]byte, error) {
	var v []byte
	err := s.db.View(func(tx *bbolt.Tx) error {
		v = tx.Bucket(boltBbucketName).Get(key)
		return nil
	})
	return v, err
}
