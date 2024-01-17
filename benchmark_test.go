package kvengines

import (
	"fmt"
	"testing"
)

type BenchmarkStore struct {
	Factory func(string, bool) Store
	Path    string
	Name    string
}

func Benchmark(b *testing.B) {
	stores := []BenchmarkStore{
		{NewPogrebStore,
			"pogreb",
			"Pogreb",
		},
		{NewBoltDb,
			"bbolt",
			"Bbolt",
		},
		{NewBadgerDb,
			"badger",
			"BadgerDB",
		},
		{NewBuntdbStore,
			"buntdb",
			"BuntDB",
		},
		{NewFlyDb,
			"flydb",
			"FlyDB",
		},
		{NewLevelDBStore,
			"leveldb",
			"GoLevelDB",
		},
		{NewLotusDbStore,
			"lotusdb",
			"LotusDB",
		},
		{NewNutsdbStore,
			"nutsdb",
			"NutsDB",
		},
		{NewPebbleStore,
			"pebble",
			"Pebble",
		},
		{NewRoseDbStore,
			"rosedb",
			"ROSEDB",
		},
	}

	type entry struct {
		key   []byte
		value []byte
	}

	valueSize := 100
	size := 1000000
	entries := make([]entry, size)
	for i := 0; i < size; i++ {
		entries[i] = entry{
			key:   GetValue(16),
			value: GetValue(valueSize),
		}
	}

	fsync := false
	fsyncS := "nofsync"

	for _, storeData := range stores {
		store := storeData.Factory(storeData.Path, fsync)

		benchName := fmt.Sprintf("%s %dx %db %s Put", storeData.Name, size, valueSize, fsyncS)
		b.Run(benchName, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				err := store.Set(entries[i].key, entries[i].value)
				if err != nil {
					fmt.Println(fmt.Sprintf("Can't set: %s %d", benchName, i))
					panic(err)
				}
			}
		})

		benchName = fmt.Sprintf("%s %dx %db %s Get", storeData.Name, size, valueSize, fsyncS)
		b.Run(benchName, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				v, err := store.Get(entries[i].key)
				if err != nil || len(v) != valueSize || v[50] != entries[i].value[50] {
					if err != nil {
						fmt.Println(fmt.Sprintf("Error get incorrect: %s i=%d", benchName, i))
						panic(err)
					} else {
						panic(fmt.Sprintf("Value get incorrect: %s i=%d k=%+v", benchName, i, entries[i].key))
					}
				}
			}
		})
	}

}
