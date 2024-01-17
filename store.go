package kvengines

type Store interface {
	Close() error
	Set(key, value []byte) error
	Get(key []byte) ([]byte, error)
	FlushDB() error
}
