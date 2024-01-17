package kvengines

import (
	"github.com/akrylysov/pogreb"
	"path/filepath"
)

type pogrebStore struct {
	db *pogreb.DB
}

func NewPogrebStore(path string, sync bool) Store {
	opts := &pogreb.Options{}

	if sync {
		// https://github.com/akrylysov/pogreb/blob/main/options.go
		opts.BackgroundSyncInterval = -1
	}
	dir := filepath.Join("benchmark", path)
	db, err := pogreb.Open(dir, opts)
	if err != nil {
		return nil
	}

	return &pogrebStore{
		db: db,
	}
}

func (s *pogrebStore) Close() error {
	s.db.Close()
	return nil
}

func (s *pogrebStore) Set(key, value []byte) error {
	return s.db.Put(key, value)
}

func (s *pogrebStore) Get(key []byte) ([]byte, error) {
	v, err := s.db.Get(key)
	return v, err
}

func (s *pogrebStore) FlushDB() error {
	return s.db.Close()
}
