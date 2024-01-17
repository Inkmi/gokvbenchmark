package kvengines

import (
	"github.com/tidwall/buntdb"
	"path/filepath"
)

type buntdbStore struct {
	db *buntdb.DB
}

func NewBuntdbStore(path string, sync bool) Store {
	dir := filepath.Join("benchmark", path)
	opts := buntdb.Config{}
	if sync {
		opts.SyncPolicy = buntdb.Always
	}
	db, err := buntdb.Open(dir)
	if err != nil {
		return nil
	}
	db.SetConfig(opts)
	return &buntdbStore{
		db: db,
	}
}

func (s *buntdbStore) Close() error {
	s.db.Close()
	return nil
}

func (s *buntdbStore) Set(key, value []byte) error {
	return s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(string(key), string(value), nil)
		return err
	})
}

func (s *buntdbStore) Get(key []byte) ([]byte, error) {
	var v []byte

	err := s.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(string(key))
		if err == nil {
			v = []byte(val)
		}
		return err
	})

	return v, err
}

func (s *buntdbStore) FlushDB() error {
	return s.db.Update(func(tx *buntdb.Tx) error {
		return tx.DeleteAll()
	})
}
