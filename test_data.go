package kvengines

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/gofrs/uuid"
	"hash/fnv"
)

func GetKey(n int) []byte {

	return []byte("test_key_" + fmt.Sprintf("%09d", n))
}

func intToRandomBytes(integer int) []byte {
	hasher := fnv.New128()
	intBytes := make([]byte, 16)
	binary.BigEndian.PutUint64(intBytes, uint64(integer))
	_, err := hasher.Write(intBytes)
	if err != nil {
		panic(err)
	}
	hash := hasher.Sum(nil)
	return hash[:16]
}

func GetRandomKey(n int) []byte {
	return intToRandomBytes(n)
}

func GetUuidKey() []byte {
	return uuid.Must(uuid.NewV7()).Bytes()
}

func GetValue(length int) []byte {
	token := make([]byte, length)
	rand.Read(token)
	return token
}
