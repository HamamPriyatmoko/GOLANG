package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close() // wajib diclose setelah digunakan

	ctx := context.Background()

	// ExecContext ini tidak bisa mengembalikan result DATA dari database hanya result aja
	// jadi untuk exec ini lebih baik jika digunakan untuk query yang tidak mengembalikan data database
	// Jika ingin melakukan yang bukan query data disarankan pakai ini
	var query string = "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('joko', 'Joko Anwar', 'joko@gmail.com', 200000, 3.8, '1984-8-8', false)"
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		fmt.Println("There are errors")
		panic(err)
	}

	// Ini akan berfungsi kalau idnya pakai auto increment dan akan mengembalikan Id terakhir yang diinsert
	// rows, err := result.LastInsertId()

	// Ini akan mengembalikan hasil dari berapa row yang kena dampaknya
	rows, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
	fmt.Println(result)
	fmt.Println("Succes insert data to database")
}

// Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec,
// namun jika kita membutuhkan result, seperti SELECT SQL, kita bisa menggunakan function yang berbeda
// Function untuk melakukan query ke database, bisa menggunakan function
// (DB) QueryContext(context, sql, params)

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close() // wajib diclose setelah digunakan

	ctx := context.Background()

	var script string = "SELECT * FROM Customer"
	// rows adalah kumpulan dari baris data ditable tersebut dan bentuknya adalah iteration
	// jadi kita perlu melakukan iteration pada rows jika ingin mendapatkan semua hasil query
	// dan rows ini seperti cursor dia akan mengembalikan data pertama dulu dan jika ingin mendapatkan
	// data selanjutnya harus gunakan perintah Next
	// Ini akan mengembalikan sebuah rows data jadi cocok buat select
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		fmt.Println("There are errors")
		panic(err)
	}

	// rows perlu ditutup jika sudah tidak menggunakan lagi
	defer rows.Close()

	var result = make([]string, 0)

	// Ini akan membaca row selanjutnya, jika row punya data selanjutnya dia akan menghasilkan true dan akan
	// membaca data selanjutnya jika ternyata terdapat data lagi dan untuk data pertama itu tidak ada, jadi data
	// kedua adalah data pertama yang ada didatabase
	for rows.Next() {
		var (
			a    string
			name string
		)
		// Scan akan membaca/menyalin column dari row perulangan saat ini ke destination yang ada diparameter
		// jadi parameternya adalah tujuan column itu akan disalin maka dari itu dia bersifat pointer
		// dan tipe datanya apapun yang sesuai/merepresentasikan tipe data pada column tersebut
		if err := rows.Scan(&a, &name); err != nil {
			panic(err)
		}

		fmt.Println("Id:", a)
		fmt.Println("Name:", name)
		// Cara push kedalam slice gunakan append
		// result = append(result, a, name)
	}

	fmt.Println(result)

	// fmt.Println(rows)
	// fmt.Println(result)
	fmt.Println("Succes insert data to database")
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close() // wajib diclose setelah digunakan

	ctx := context.Background()

	// Best practicenya sebut semua nama column yang ingin diambil walaupun semua column diambil
	// karena yang ditakutkan adalah kita salah posisi dalam merepresentasikan columnya nanti
	var script string = "SELECT id, name, email, balance, rating, created_at, birth_date, married  FROM Customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		fmt.Println("There are errors")
		panic(err)
	}

	// rows perlu ditutup jika sudah tidak digunakan lagi/rows telah selesai mengembalikan semua datanya
	defer rows.Close()

	var result = make([]string, 0)

	// Ini akan membaca row selanjutnya, jika row punya data selanjutnya dia akan menghasilkan true dan akan
	// membaca data selanjutnya jika ternyata terdapat data lagi dan untuk data pertama itu tidak ada, jadi data
	// kedua adalah data pertama yang ada didatabase
	for rows.Next() {
		// Golang database tidak mengerti dengan tipe data NULL didatabase
		// Oleh karena itu, khusus untuk kolom yang bisa NULL didatabase, akan jadi masalah
		// jika kita melakukan Scan secara bulat-bulat menggunakan tipe data representasinya di Golang
		var (
			a          string
			name       string
			email      sql.NullString // Pastikan tipe datanya sama seperti yang ada dideklarasikan dicolumn database
			balance    int64
			rating     float32
			created_at time.Time
			birth_date sql.NullTime
			married    bool
		)
		// Scan akan membaca/menyalin column dari row perulangan saat ini ke destination yang ada diparameter
		// jadi parameternya adalah tujuan column itu akan disalin maka dari itu dia bersifat pointer
		// dan tipe datanya apapun yang sesuai/merepresentasikan tipe data pada column tersebut
		if err := rows.Scan(&a, &name, &email, &balance, &rating, &created_at, &birth_date, &married); err != nil {
			panic(err)
		}

		fmt.Println("====================")
		fmt.Println("Id:", a)
		fmt.Println("Name:", name)
		// Valid akan bernilai true jika String punya data atau bukan null
		if email.Valid {
			fmt.Println(email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)

		// Conversi dari []byte/[]uint8 ke string ada 3 cara
		// Pertama: Cara agar pake reference
		// fmt.Println("Created_at:", unsafe.String(&created_at[0], len(created_at)))
		// Kedua: Cara Copy data
		// fmt.Println("Created_at:", string(created_at))
		// Ketiga: Lewat Connection Didatabasenya langsung dengan menambahkan parameter parseTime=true
		fmt.Println("Created_at:", created_at)
		// fmt.Println("Birthday:", birth_date)
		if birth_date.Valid {
			fmt.Println(birth_date.Time)
		}

		fmt.Println("Married:", married)
		fmt.Println("====================")
		// Cara push kedalam slice gunakan append
		// result = append(result, a, name)
	}

	fmt.Println(result)

	// fmt.Println(rows)
	// fmt.Println(result)
	fmt.Println("Succes insert data to database")
}

// Ini simulasi sql injection
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close() // wajib diclose setelah digunakan

	ctx := context.Background()
	// Didalam sql # dianggap sebuah comment jadi query setelah tanda # akan dianggap comment
	username := "admin'; #"
	password := "salah"

	var script string = "SELECT username FROM user " + "where username = '" + username +
		"' And password = '" + password + "' Limit 1"

	fmt.Println(script)
	// rows adalah kumpulan dari baris data ditable tersebut dan bentuknya adalah iteration
	// jadi kita perlu melakukan iteration pada rows jika ingin mendapatkan semua hasil query
	// dan rows ini seperti cursor dia akan mengembalikan data pertama dulu dan jika ingin mendapatkan
	// data selanjutnya harus gunakan perintah Next ini akan mengembalikan sebuah
	// rows data jadi cocok buat select
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		fmt.Println("There are errors")
		panic(err)
	}

	// rows perlu ditutup jika sudah tidak menggunakan lagi
	defer rows.Close()

	// Ini akan membaca row selanjutnya, jika row punya data selanjutnya dia akan menghasilkan true dan akan
	// membaca data selanjutnya jika ternyata terdapat data lagi dan untuk data pertama itu tidak ada, jadi data
	// kedua adalah data pertama yang ada didatabase
	if rows.Next() {
		var (
			username string
		)
		// Scan akan membaca/menyalin column dari row perulangan saat ini ke destination yang ada diparameter
		// jadi parameternya adalah tujuan column itu akan disalin maka dari itu dia bersifat pointer
		// dan tipe datanya apapun yang sesuai/merepresentasikan tipe data pada column tersebut
		if err := rows.Scan(&username); err != nil {
			panic(err)
		}

		fmt.Println("Username:", username)
		fmt.Println("Succes login admin")
	} else {
		fmt.Println("Gagal login")
	}

}

func TestSqlParameter(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	defer db.Close()

	username := "admin; #"
	password := "salah"

	sqlQueries := "select username from user where username = ? AND password = ? LIMIT 1"
	// Parameter ke 3 dan seterusnya adalah variadic argument
	rows, err := db.QueryContext(ctx, sqlQueries, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		if err := rows.Scan(&username); err != nil {
			panic(err)
		}

		fmt.Println("Username:", username)
		fmt.Println("Succes login admin")
	} else {
		fmt.Println("Gagal login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close() // wajib diclose setelah digunakan

	ctx := context.Background()

	username := "Hamam'; DROP TABLE user; #"
	password := "Hamam123"

	var query string = "INSERT INTO user(username, password) VALUES(?,?)"
	result, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		fmt.Println("There are errors")
		panic(err)
	}

	// Ini akan berfungsi kalau idnya pakai auto increment dan akan mengembalika Id terakhir yang diinsert
	// rows, err := result.LastInsertId()

	// Ini akan mengembalikan hasil dari berapa row yang kena dampaknya
	rows, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
	fmt.Println(result)
	fmt.Println("Succes insert data to database")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close() // wajib diclose setelah digunakan

	ctx := context.Background()

	email := "hamam@gmail.com"
	comment := "Hallo This is my first comment"

	var query string = "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		fmt.Println("There are errors")
		panic(err)
	}

	// Ini akan berfungsi kalau idnya pakai auto increment dan akan mengembalika Id terakhir yang diinsert
	rows, err := result.LastInsertId()

	// Ini akan mengembalikan hasil dari berapa row yang kena dampaknya
	// rows, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
	fmt.Println(result)
	fmt.Println("Succes insert data to database", rows)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	// Gunakan PrepareStatement jika kamu menghadapi perulangan (loop)
	// yang melakukan perintah SQL yang sama berkali-kali

	// Tanpa Prepare, jika kamu menjalankan INSERT 10 kali dalam sebuah perulangan,
	// database akan melakukan 3 tahap ini sebanyak 10 kali:
	// 1. Parsing: Memeriksa sintaks SQL (apakah ada typo?).
	// 2. Compile: Mengubah SQL jadi instruksi mesin.
	// 3. Execute: Menjalankan perintah.
	var query string = "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	// Contoh Use Case di Real Life
	// Kapan kamu wajib pakai ini? Biasanya pada skenario Massive Data Processing.
	// Skenario: Import Data Excel ke Database
	// Bayangkan kamu diminta membuat fitur import data 5.000 karyawan dari file Excel.
	// Cara Buruk (Tanpa Prepare):
	// Kamu memanggil db.ExecContext di dalam loop 5.000 kali. Database akan lelah melakukan parsing 5.000 kali
	// untuk kalimat yang hampir sama. Ini memboroskan resource server.
	for i := 0; i < 10; i++ {
		email := "hamam" + strconv.Itoa(i) + "@gmail.com"
		comment := "commentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin() // Untuk memulai transaction
	if err != nil {
		panic(err)
	}

	var query string = "INSERT INTO comments(email, comment) VALUES(?,?)"
	// Do transaction here
	for i := 0; i < 10; i++ {
		email := "hamam" + strconv.Itoa(i) + "@gmail.com"
		comment := "commentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		// if i == 9 {
		// 	err := tx.Rollback()
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	break
		// }
		fmt.Println("Comment Id", id)
	}

	// err = tx.Commit()
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
