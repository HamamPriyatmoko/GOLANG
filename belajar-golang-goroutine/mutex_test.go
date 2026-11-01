package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	// Ini demo test kondisi terjadinya race condition
	x := 0
	// Untuk melakukan locking agar tidak race condition pake Mutex
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		// Ada kemungkinan gouroutine dieksekusi secara bebarengan/paralel
		go func() {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				// Variable ini akan terus bertambah dan dilakukan lock terhadap eksekusi goroutine yang lain
				x += 1
				mutex.Unlock()
			}
		}()
	}

	// Selain mutex, agar kita tidak manual melakukan Lock ataupun Unlock kita bisa gunakan
	// atomic untuk contohnya bisa dilihat difile *atomic_test.go*

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}

type BankAccount struct {
	RWMutext sync.RWMutex
	Balance  int
}

func (account *BankAccount) AddBalance(amount int) {
	// Ini Lock untuk proses Write
	account.RWMutext.Lock()
	account.Balance += amount
	// Ini Unlock untuk proses Write
	account.RWMutext.Unlock()
}

func (account *BankAccount) GetBalance() int {
	// Ini method untuk lock proses Read
	account.RWMutext.RLock()
	balance := account.Balance
	// Ini method untuk Unlock proses Read
	account.RWMutext.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// Go akan secara otomatis merubah account menjadi pointer account
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User1:", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User2:", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

// Ini simulasi deadlock
func TestDeadLock(t *testing.T) {
	// Ini simulasi deadlock
	user1 := UserBalance{
		Name:    "Hamam",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Iqbal",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000) // Goroutine 1  // Saat ingin melakukan unlock sudah keburu dilock sama
	go Transfer(&user2, &user1, 200000) // Goroutine 2  // goroutine kedua ini
	// Sehingga karna keduanya lock maka mereka kekunci dieksekusi user1 aja jadi eksekusi user 2 tidak berjalan
	// karena sudah kena deadlock

	time.Sleep(3 * time.Second)

	fmt.Println("Nama User:", user1.Name, ", Balance:", user1.Balance)
	fmt.Println("Nama User:", user2.Name, ", Balance:", user2.Balance)
}
