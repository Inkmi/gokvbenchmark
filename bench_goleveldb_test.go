package kvengines

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"log"
	"path/filepath"
)

type leveldbStore struct {
	db   *leveldb.DB
	sync bool
}

func NewLevelDBStore(path string, sync bool) Store {
	dir := filepath.Join("benchmark", path)
	var err error
	levelDb, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &leveldbStore{
		db:   levelDb,
		sync: sync,
	}
}

func (s *leveldbStore) Close() error {
	return s.db.Close()
}

func (s *leveldbStore) Set(key, value []byte) error {
	return s.db.Put(key, value, &opt.WriteOptions{Sync: s.sync})
}

func (s *leveldbStore) Get(key []byte) ([]byte, error) {
	v, err := s.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *leveldbStore) FlushDB() error {
	return nil
}
