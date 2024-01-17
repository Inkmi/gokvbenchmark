package kvengines

import (
	"github.com/ByteStorage/FlyDB/config"
	"github.com/ByteStorage/FlyDB/engine"
	"github.com/ByteStorage/FlyDB/flydb"
	"path/filepath"
)

type flyStore struct {
	db *engine.DB
}

func NewFlyDb(path string, sync bool) Store {
	opts := config.DefaultOptions
	opts.DirPath = filepath.Join("benchmark", path)
	opts.SyncWrite = sync
	FlyDB, err := flydb.NewFlyDB(opts)
	if err != nil {
		panic(err)
	}
	return &flyStore{
		db: FlyDB,
	}
}

func (s *flyStore) Close() error {
	err := s.db.Close()
	return err
}

func (s *flyStore) FlushDB() error {
	return nil
}

func (s *flyStore) Set(key, value []byte) error {
	err := s.db.Put(key, value)
	return err
}

func (s *flyStore) Get(key []byte) ([]byte, error) {
	v, err := s.db.Get(key)
	return v, err
}
