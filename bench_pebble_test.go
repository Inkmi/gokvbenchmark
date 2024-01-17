package kvengines

import (
	"github.com/cockroachdb/pebble"
	"math"
	"path/filepath"
)

type pebbleStore struct {
	db   *pebble.DB
	sync bool
}

func NewPebbleStore(path string, sync bool) Store {
	dir := filepath.Join("benchmark", path)
	opts := &pebble.Options{
		BytesPerSync: math.MaxInt,
	}
	var err error
	pebbledb, err := pebble.Open(dir, opts)
	if err != nil {
		panic(err)
	}
	return &pebbleStore{
		db:   pebbledb,
		sync: sync,
	}
}

func (s *pebbleStore) Close() error {
	s.db.Close()
	return nil
}
func (s *pebbleStore) Set(key, value []byte) error {
	return s.db.Set(key, value, &pebble.WriteOptions{
		Sync: s.sync,
	})
}

func (s *pebbleStore) Get(key []byte) ([]byte, error) {
	v, closer, err := s.db.Get(key)
	if err != nil && closer != nil {
		closer.Close()
	}
	return v, err
}

func (s *pebbleStore) FlushDB() error {
	return s.db.Flush()
}
