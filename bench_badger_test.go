package kvengines

import (
	"github.com/dgraph-io/badger/v4"
	"path/filepath"
)

type badgerStore struct {
	db *badger.DB
}

func NewBadgerDb(path string, sync bool) Store {
	opts := badger.DefaultOptions(filepath.Join("benchmark", path))
	opts.SyncWrites = sync
	d, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return &badgerStore{
		db: d,
	}
}

func (s *badgerStore) FlushDB() error {
	return s.db.DropAll()
}

func (s *badgerStore) Set(key, value []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (s *badgerStore) Get(key []byte) ([]byte, error) {
	var v []byte
	err := s.db.View(func(txn *badger.Txn) error {
		if item, err := txn.Get(key); err == nil {
			v, _ = item.ValueCopy(v)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *badgerStore) Close() error {
	return s.db.Close()
}
