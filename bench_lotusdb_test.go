package kvengines

import (
	"github.com/lotusdblabs/lotusdb/v2"
	"path/filepath"
)

type lotusStore struct {
	db   *lotusdb.DB
	sync bool
}

func NewLotusDbStore(path string, sync bool) Store {
	options := lotusdb.DefaultOptions
	options.Sync = sync
	options.DirPath = filepath.Join("benchmark", path)

	// open a database
	db, err := lotusdb.Open(options)
	if err != nil {
		panic(err)
	}

	return &lotusStore{
		db:   db,
		sync: sync,
	}
}

func (s *lotusStore) Close() error {
	s.db.Close()
	return nil
}
func (s *lotusStore) Set(key, value []byte) error {
	return s.db.Put(key, value, &lotusdb.WriteOptions{
		Sync: s.sync,
	})
}

func (s *lotusStore) Get(key []byte) ([]byte, error) {
	v, err := s.db.Get(key)
	return v, err
}

func (s *lotusStore) FlushDB() error {
	return s.db.Sync()
}
