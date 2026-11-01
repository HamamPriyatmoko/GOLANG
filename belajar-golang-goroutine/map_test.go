package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)
// Semua fitur yang ada difitur sync akan aman terhadap race condition

// Map memiliki sifat tahan dalam race condition/tidak akan terjadi race condition
// Map ini mirip Go-Lang map, namun yang membedakan, Map ini aman untuk penggunaan concurrent
// menggunakan goroutine

func addDataToMap(value int, data *sync.Map) {
	// Memiliki parameter Key and Value
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	var data sync.Map
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		// go addDataToMap(i, &data, &group)

		// Method Go disediakan oleh WaitGroup agar tidak usah secara manual
		// melakukan Add goroutine ke Wait dan Done, jadi Go ini akan otomatis melakukannya sendiri
		group.Go(func() {
			addDataToMap(i, &data)
		})
	}

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)

		// True jika ingin dilakukan iterasi semuanya
		return true
	})

	group.Wait()
}
