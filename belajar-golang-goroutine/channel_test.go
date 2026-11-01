package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Karena goroutine tidak bisa menrima return value dari sebuah function
// maka cara untuk mendaptkan return value dari function adalah menggunakan Channel
// Channel adalah tempat komunikasi secara synchronus(berjalan secara sync) yang bisa dilakukan oleh goroutine
// diChannel terdapat pengirim dan penerima, dan biasanya pengirim dan penerima adalah Goroutine yang berbeda
// Saat melakukan pengiriman data ke Channel, Goroutine akan ter-block, sampai ada yang menerima data tersebut
// Channel disebut juga synchronus blocking
// Channel bisa menjadi async await jika dipemrograman lain
// Untuk mengirim dari goroutine satu ke goroutine yang lain bisa pake channel

// Karakteristik Channel:
// Channel Hanya Bisa Menampung Satu Data, jika ingin menambah satu data harus menunggu data sebelumnya diambil
// Channel hanya bisa menrima satu jenis data
// Channel bisa diambil dari lebih satu goroutine
// Channel harus diclose jika tidak bisa memory leak

func TestCreateChannel(t *testing.T) {
	// Gunakan Make dan tambahakn diparameternya dengan keyword chan untuk membuat channel lalu diikuti
	// tipe data yang akan dikirim lewat channel dan sifat channel adalah pass by reference jadi ga perlu pointer
	channel := make(chan string)

	// Untuk mengirim data caranya seperti dibawah ini gunakan arrow kiri
	channel <- "Hamam"

	// Untuk menerima data dengan cara dibawah, "data" adalah sebuah variable
	// jika sebelumnya belum deklarasi variable data maka bisa pake cara dibawah
	data := <-channel
	fmt.Println(data)
	// jika sudah punya variablenya bisa pake cara ini:
	// namaVariable = <- channel

	// jika ingin kirim data keparameter begini caranya:
	fmt.Println(<-channel)

	// untuk menutup channel atau kalo pengen aman bisa pake defer
	// close(channel)
	// defer close(channel)
}

func TestCreateChannel2(t *testing.T) {
	// Gunakan make() dan tambahkan diparameternya dengan keyword chan untuk membuat channel lalu diikuti
	// tipe data yang akan dikirim lewat channel
	channel := make(chan string)
	defer close(channel)

	// tambahkan () untuk call functionnya/eksekusi functionya
	go func() {
		time.Sleep(2 * time.Second)
		// Jika dibawah channel datanya tidak ada yang mengambil maka dia akan block
		channel <- "Hamam Priyatmoko" //block
		// dan dibawah ini tidak akan dieksekusi karena diblock oleh channel
		fmt.Println("Selesai Kirim Data Ke Channel")
	}()

	// Jika channel tidak mengirim data apapun maka akan terjadi deadlock dibawah ini
	data := <-channel
	fmt.Println(data)

	// time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ini data chanel"
}

func GiveMeResponse2(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ini data chanel"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

// Gunakan chan<- jika ingin channelnya hanya mengirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Ngirim data aja nih"
	// JIKA DIPAKSA MENERIMA DATA MAKA AKAN ERROR
}

// Gunakan <-chan jika ingin channelnya hanya menerima data
func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	// buffered channel memiliki kapasitas 4
	channel := make(chan string, 4)
	defer close(channel)

	// Karena buffer maka tidak akan blocking tetapi jika data yang dikirim lebih dari
	// kapasitasnya maka akan blocking karena channel akan melakukan waiting
	go func() {
		channel <- "Hamam"
		channel <- "Priyatmoko"
		channel <- "Silvya"
		channel <- "Amanda"
	}()

	fmt.Println("Selesai") // ini bisa langsung eksekusi

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	// defer close(channel)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke- " + strconv.Itoa(i)
		}
	}()

	// Lebih baik pake range untuk menangani channel agar user lain bisa segera ditangani ketika
	// mengirim data, gunakan range jika data yang dikirim banyak
	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse2(channel2)

	// Ini adalah perulangan infinite
	counter := 0
	for {
		// Select berguna untuk mengontrol aliran data dari banyak channel secara bersamaan
		// Jika tidak menggunakan ini maka bisa jadi semisal channel1 belum mendapatkan data
		// maka akan memblocking program yang lain sampai channel1 dapat datanya, select mendengarkan
		// semua channel dan jika salah satu channel sudah ada yang memiliki data maka channel itu yang
		// akan dieksekusi duluan
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel ke-1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel ke-2", data)
			counter++
		// Ketika harus menunggu data disemua channel atau datanya tidak ada disemua channel,
		// kita bisa tambahkan default agar bisa melakukan eksekusi program ketika tidak ada data
		// yang diterima disemua channel/menunggu data
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}
