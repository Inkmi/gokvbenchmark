package kvengines

import (
	"github.com/nutsdb/nutsdb"
	"path/filepath"
)

type nutsdbStore struct {
	db *nutsdb.DB
}

var nutsBucketName = "test-bucket"

func NewNutsdbStore(path string, sync bool) Store {
	opts := nutsdb.DefaultOptions
	opts.Dir = filepath.Join("benchmark", path)
	opts.SyncEnable = sync
	opts.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	var err error
	nutsDB, err := nutsdb.Open(opts)
	if err != nil {
		panic(err)
	}
	nutsDB.Update(func(tx *nutsdb.Tx) error {
		return tx.NewBucket(nutsdb.DataStructureBTree, nutsBucketName)
	})

	return &nutsdbStore{
		db: nutsDB,
	}
}

func (s *nutsdbStore) Close() error {
	s.db.Close()
	return nil
}

func (s *nutsdbStore) Set(key, value []byte) error {
	return s.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put(nutsBucketName, key, value, 0)
	})
}

func (s *nutsdbStore) Get(key []byte) ([]byte, error) {
	var v []byte
	var err error

	s.db.View(func(tx *nutsdb.Tx) error {
		v, err = tx.Get(nutsBucketName, key)
		return err
	})

	return v, err
}

func (s *nutsdbStore) FlushDB() error {
	return s.db.Close()
}
