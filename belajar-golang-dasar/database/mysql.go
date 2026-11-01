package database

var connection string

// function init adalah keyword agar saat package digunakan maka secara otomatis 
// akan menjalankan function init
func init() { 
	connection = "mysql"
}

func GetDatabase() string {
	return connection
}