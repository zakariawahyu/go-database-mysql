package go_database_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

/*
Membuat koneksi ke database
- Hal pertama yang akan kita lakukan ketika membuat aplikasi yang akan menggunakan database adalah melakukan koneksi ke databasenya
- UNtuk melakukan koneksi ke database di Go-Lang, kita bisa membuat object sql.DB menggunakan function sql.Open(driver,dataSourceName)
- Untuk menggunakan database mySQL, kita bisa menggunakan driver "mysql"
- Sedangkan untuk dataSourName, tiap database mempunyai cara pengulisan masing-masing. misal di MySQL kita bisa menggunakan datasource dengan format
username:password@tcp(host:port)/database_name
- Jika object sql.DB sudah tidak digunakan lagi, disarakankan untuk menutupnya menggunakan function Clode()
- Jika tidak ditutup maka akan terjadi koneksi yang leak (berulang tidak ada hentinya)
*/
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:masukdb@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err)
	}

	db.Close()
}
