package go_database_mysql

import (
	"database/sql"
	"time"
)

/*
Database Pooling
- sql.DB di Go-Lang sebenarnya bukanlah koneksi ke database, melainkan sebuah pooling.
- Database pooling adalah kumpulan dari beberapa koneksi
- Di dalam sql.DB, Go-Lang melakukan management koneksi ke database secara otomatis. Kegunakan database pooling
menjadikan kita tidak perlu melakukan management koneksi database secara manual
- Dennan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yang dibuat oleh Go-Lang,
sehingga tidak membanjiri koneksi ke datahase, karena biasanya ada batas maksimal koneksi yang bisa ditangani oleh database yang kita gunakan
- Kenapa butuh minimal koeksi? karena jika set minimal 10, by default ketika aplikasi pertama kali menyala akan langsung membuat 10 koneksi kedatabase
kegunaan ini jika tiba-tiba koneksi naik biar bisa ngejar dan tidak kelamaan untuk membuat koneksi.
- Jika trafictnya tinggi, bisa di set maksimal. jika maksimal 50 dan tiba-tiba ada 100 request, maka akan dieksekusi 50 koneksi dulu dan yang lainya menunggu

Pengaturan Database Pooling
(DB)SetMaxIdleConns(number)		= Pengaturan berapa jumlah koneksi minimal yang dibuat
(DB)SetMaxOpenCoons(number)		= Pengaturan berapa jumlah koneksi maksimal yang dibuat
(DB)SetConnMaxIdleTime(number)	= Pengaturan berapa lama koneksi yang sudah tidak digunakan akan langsung dihapus
(DB)SetConnMaxLifetime(number)	= Pengaturan berapa lama koneksi boleh digunakan
*/

/**
Error Tipe Data Date
- Secara default DRiver MySQL untuk Go-Lang akan melakukan query tipe data DATE, DATETIME, TIMESTAMP menjadi []bytes / []uint8
- Dimana ini bisa di konversi menjadi string lalu di parsing ke time.Time
- Nammnun hal ini merepotkan jika dijalankan manual, kita bisa meminta Driver MySQL untuk golang secara otomatis melakukan parsing
- Parsing otomatis tinggal menambahkan parseTime=true pada dataSourceName
*/

func GetConnection() *sql.DB {
	//db, err := sql.Open("mysql", "root:masukdb@tcp(localhost:3306)/golang_database")
	// konversi date otomatis
	db, err := sql.Open("mysql", "root:masukdb@tcp(localhost:3306)/golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(15 * time.Minute)

	return db
}
