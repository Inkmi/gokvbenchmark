package kvengines

import (
	"github.com/rosedblabs/rosedb/v2"
	"path/filepath"
)

type roseDbStore struct {
	db *rosedb.DB
}

func NewRoseDbStore(path string, sync bool) Store {
	opts := rosedb.DefaultOptions
	opts.Sync = sync
	opts.DirPath = filepath.Join("benchmark", path)
	var err error
	roseDB, err := rosedb.Open(opts)
	if err != nil {
		panic(err)
	}
	return &roseDbStore{
		db: roseDB,
	}
}

func (s *roseDbStore) Close() error {
	s.db.Close()
	return nil
}
func (s *roseDbStore) Set(key, value []byte) error {
	return s.db.Put(key, value)
}

func (s *roseDbStore) Get(key []byte) ([]byte, error) {
	v, err := s.db.Get(key)
	return v, err
}

func (s *roseDbStore) FlushDB() error {
	return s.db.Sync()
}
