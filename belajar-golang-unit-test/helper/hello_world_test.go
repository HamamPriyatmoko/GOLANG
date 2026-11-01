package helper

// Nama file untuk unit test pastikan menggunakan _test diakhir nama file
// Lalu untuk melakukan pengetesan pada sebuah function pastikan diawali nama functionnya dengan "Test"
// lalu diikuti dengan nama function yang akan ditest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// Bedanya assert dan require adalah klo require akan memanggil FailNow sedangkan assert Fail
)

// Method dari object testing.T:
// Fail() akan menggagalkan unit test namun tetap melanjutkan eksekusi unit test
// FailNow() akan menggagalkan unit test saat ini juga tanpa melanjutkan eksekusi unit test
// Error() akan memberikan catatan dengan diikuti oleh Fail setelahnya sehingga unit test dianggap gagal
// Fatal() akan memberikan catatan tapi diikuti oleh FailNow jdi unit test akan gagal dan langsung berhenti

func BenchmarkHelloWorldTable(b *testing.B)  {
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name: "HelloWorld(Hamam)",
			request: "Hamam",
		},
		{
			name: "HelloWorld(Hamam Priyatmoko)",
			request: "Hamam Priyatmoko",
		},
		{
			name: "HelloWorld(Irawan Rofi Hidayat)",
			request: "Irawan Rofi Hidayat",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// Untuk melakukan sub benchmark gunakan method Run
func BenchmarkSub(b *testing.B)  {
	b.Run("Hamam Priyatmoko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hamam Priyatmoko")
		}
	})

	b.Run("Irawan Rofi Hidayat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Irawan Rofi Hidayat")
		}
	})
}

// testing.B adalah object untuk melakukan testing benchmark pada sebuah code function
// dengan melakukan iteration yang dimana sudah ditentukan oleh testing.B menggunakan b.N
// Dan pastikan untuk menguji benchmark nama functionnya harus diawali dengan kata "Benchmark"
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hamam Priyatmoko")
	}
}

func TestHelloWorldTable(t *testing.T) {
	dataTests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Hamam)",
			request:  "Hamam",
			expected: "Hello Hamam",
		},
		{
			name:     "HelloWorld(Priyatmoko)",
			request:  "Priyatmoko",
			expected: "Hello Priyatmoko",
		},
		{
			name:     "HelloWorld(Silvya)",
			request:  "Silvya",
			expected: "Hello Silvya",
		},
	}

	for index, test := range dataTests {
		fmt.Println("Test", index)
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Must be "+test.expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	// Ini adalah cara membuat sub unit test didalam unit test
	t.Run("Hamam", func(t *testing.T) {
		result := HelloWorld("Mam")
		require.Equal(t, "Hello Mam", result, "Must be Hello Mam")
	})

	// t.Run(Nama Functionya Unit Testnya, func(t *testing.T) {

	// })

	// Jika ingin menjalankan hanya salah satu sub testnya saja pakai ini
	// go test -run TestSubTest/(Nama SubTestnya)

	// Jika ingin menjalankan semua unit Test/sub test yang pakai nama tertentu saja pakai ini
	// go test -run /(Nama SubTestnya)

	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Must be Hello Budi")
	})

}

// Mirip sebuah method Before dan After dijest, method/Function TestMain akan melakukan hal itu
// testing.M berguna untuk mengatur life cycle
func TestMain(m *testing.M) {
	fmt.Println("Code Before Unit Test Example Connect to database first before execute all unit test")

	m.Run() // Untuk running semua unit test

	fmt.Println("Code After all unit test finish")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on windows")
	}
	result := HelloWorld("Hamam")
	// Ini tidak akan dieksekusi jika ifnya skip dieksekusi tetapi tidak akan menggagalkan unit test
	require.Equal(t, "Hello Hamam", result, "Must be Hello Hamam")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Mam")
	assert.Equal(t, "Hello Mam", result, "Must be Hello Mam") // Kalo gagal akan memanggil Fail
}

func TestHelloWorldRequire(t *testing.T) {
	result := SayGoodBye("Hamam")
	require.Equal(t, "Good Bye Hamam", result, "Must be Good Bye Hamam")
	fmt.Println("Jika code diatas gagal maka ini tidak akan diprint atau dijalankan")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Mam")
	if result != "Hello Hamam" {
		// t.Error / t.Fail berfungsi untuk mengirim signal failure dan tidak akan mengehntikan proses
		t.Error("Gagal, nilai yang dikeluarkan tidak sesuai")
	}
}

// go test -v digunakan jika ingin melihat detail function apa saja yang ditest
// go test -v -run TestNamaFunction digunakan jika ingin memilih function yang ingin dirunning
// go test ./... digunakan ketika kita berada difolder paling atas dan untuk melakukan running semua test

func TestSayGoodBye(t *testing.T) {
	result := SayGoodBye("Budi")
	if result != "Good Bye Budi" {
		t.Error("Result Is Not 'Good Bye Budi'")
	}
}

// Note:
// Gunakan "go mod tidy" jika terjadi missing go sum entry, karna biasanya masalhnya adalah
// library yang coba kamu gunakan butuh library lain untuk dijalankan/perlu disamakan go sum dari librarynya
// jadi go mod tidy berfungsi membersihkan dan menambahkan jika ada yang dibutuhkan dan diperbaiki terkait
// package/library
