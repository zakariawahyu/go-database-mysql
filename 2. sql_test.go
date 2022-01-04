package go_database_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

/**
Eksekusi Perintah SQL
- Saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL
- Di Go-Lang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan sebuah function
functionya adalah (DB)ExecContext(context,sql, params)
- Ketika mengirim perintah SQL, kita butuh mengirimkan context, dengan context kita bisa mengirim sinyal cancel jika kita ingin
membatalkan pengiriman perintah SQLnya
*/

func TestExecSqlInsert(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('1', 'Zakaria', 'zaka@gmail.com', 100000, 5.0, '1999-06-06', false )")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Berhasil insert data")
}

func TestExecSqlUpdate(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "UPDATE customer SET name='Nur', email = 'nur@gmail.com' WHERE id='1'")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Berhasil update data")
}

func TestExecSqlDelete(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "DELETE FROM customer WHERE id='1'")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Berhasil delete data")
}

/**
Query Select
- Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec, namun jika kita membutuhkan result seperti SEELCT SQL kita bisa menggunakan funcition berbeda
- Function untuk melakukan query ke database bisa menggunakan (DB).QueryCOntext(context, sql, params)
*/

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, name, email, balance, rating, birth_date, married, created_at  FROM customer")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
}

/**
Rows
- Hasil Query function adalah sebuah data structs sql.Rows
- Rows digunakan untuk melakukan iterasi terhadap hasil dari query
- Kita bisa menggunakan function (Rows).Next()()(Boolean) untuk melakukan iterasi terhadap data hasil query
- Boolean jika return true masih ada data, jika false sudah tidak ada lagi data dalam rows tersebut
- Untuk membaca data kita bisa menggunakan data (Rows).Scan(namaColumns)
- Dan jangan lupa setelah menggunakan rows kita bisa menutupnya dengan (Rows).Close()
*/

/**
Tipe Data Column
- Sebelumnya kita hanya membuat dengan tipe data varchar, untuk varchar dalam database kita gunalan string di Go-Lang
- Bagaimana dengan tipe data lain?
- Apa representasinya di Go-Lang misal tipe data timestamp, date dll

Mapping Tipe Data
- VARCHAR, CHAR								=> string
- INT, BIGINT								=> int32, int64 (sesuai kapasitasnya)
- FLOAT, DOUBLE								=> float32, float64
- BOOLEAN									=> bool
- DATE, DATETIME, TIMESTAMP, TIME			=> time.Time

Error Data Null / Nullable Type
- GO-Lang database tidak bisa membaca deengan tipe data NULL
- Oleh karena itu, khuus kolom yang bisa null/Nullable di database, akan jadi masalah jika kita melakukan scan secaar bulat-bulat
menggunakan tipe data representasinya gi GO-Lang
- Konversi secara otomatis NULL tidak di dukung oleh driver MySQL Go-Lang
- Oleh karena itu, khusus tipe kolom yang bisa NULL kita perlu menggunakan tipe data yang ada dalam package sql

Tipe Data Nullable
- string			=> database/sql.NullString
- bool				=> database/sql.NullBool
- float64			=> database/sql.NullFloat64
- int32				=> database/sql.NullInt32
- int64				=> database/sql.NullInt64
- time				=> database/sql.NullTime
*/

func TestRowsResult(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, name, email, balance, rating, birth_date, married, created_at  FROM customer")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		errr := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if errr != nil {
			panic(errr)
		}
		fmt.Println("================")
		fmt.Println("ID:", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date: ", birthDate.Time)
		}
		fmt.Println("Married: ", married)
		fmt.Println("Created At: ", createdAt)
		fmt.Println("================")
	}

	defer rows.Close()
}

/**
SQL Dengan Parameter
- Saat kita membuat aplikas, kita tidak mungkin akan melakukan hardcode perintah di SQL di kode Go-Lang
- Biasanya kita akan menerima input data dari user, lalu membuat perintah SQL dari input user dan mengirimnya menggunakan perintah SQL
*/

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	username := "admin'; # "
	password := "admin"

	rows, err := db.QueryContext(ctx, "SELECT username FROM user WHERE username='"+username+"' AND password='"+password+"' LIMIT 1")
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login: ", user)
	} else {
		fmt.Println("Gagal login")
	}
	defer rows.Close()
}

/**
SQL Injection
- SQL Injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi dalam lapisan basis data sebuah aplikasi
- Biasanya SQL Injection dilakukan dengan cara mengirim input dari user dengan perintah yang salah, sehingga menyebabkan hasil SQL yang kita input tidak valid
- SQL Injection sangat berbahaya, jika sampai salah membuat SQL, bisa jadi data kita tidak aman

Kode SQL Injection
- username := "admin'; #"
- password := salah

Hasil query diatas dimanipulasi dimana merubah dapat merubah query
SELECT username FROM user WHERE username='admin'; # AND password='"+password+"' LIMIT 1"

Solusi?
- Jangan membaut query SQL secara manual dengan menggabungkan string secara bulat-bulat
- Jika kita membutuhkan parameter ketika membuat SQL, kita bisa menggunakan function Execute atau Query dengan parameter
- Sebenarnya function EXec dan Query memiliki parameter tambahan yang bisa kita gunakan untuk mensubtitusi parameter dari
function tersebut ke SQL query yang kita buat
- Untuk menandai sebuah SQL memebutuhkan parameter, kita bisa menggunakan karakter ? (tanda tanya)
*/

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	sqlQuery := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login: ", user)
	} else {
		fmt.Println("Gagal login")
	}
	defer rows.Close()
}

func TestExcelSqlParameter(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	user := "zaka'; DROP TABLE user; #" // terbebas sql injection, data dimasukkan mentah-mentah
	passwd := "zaka"

	sqlQuery := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, sqlQuery, user, passwd)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Sukses insert data")
}

/**
Auto Increment
- Kadang kita membuat sebuah table dengan id auto increment
- Dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam MySQL
- Sebenarnya kita bisa melakukan query ulang ke database dengan menggunakan SELECT LAST_INSERT_ID()
- Tapi untungnya di golang ada cara yang lebih mudah
- Kita bisa menggunakan function (Result)LastInsertId() untuk mendapatkan ID terakhir yang dibuat secara auto increment
- Result adalah object yang dikembalikan ketika kita menggunakan function Exec
*/

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	email := "zakaria@gmail.com"
	comment := "test comment"
	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, sqlQuery, email, comment)
	if err != nil {
		panic(err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Last insert ID :", lastID)
}

/**
Query atau Exec dengan Parameter
- Saat kita menggunakan function Query atau Exec yang menggunakan parameter, sebenarnya implementasi dibahwahnya menggunakan feature Prepare Statement
- Jadi tahapan pertama nya disiapkan terlebih dahulu, setelah itu baru di isi dengan parameter
- Kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, hanya berbeda parameter. Misal insert dengan data banyak sekaligus
- Pembuatan Prepare Statement bisa dilakukan manual, tanpa harus menggunakan Query atau Exec dengan parameter

Prepare Statement
- Saat kita membuat Prepare Statement, secara otomatis akan mengenali koneksi database yang digunakan
- Sehingga ketika kita mengeksekusi Prepare Statement berkali-kali, maka akan menggunakan koneksi yang sama dan lebih efisien karena pembuatan prepare statementnya hanya sekali diawal saja
- Jika menggunakan Query dan Exec dengan parameter, kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama, oleh karena itu, bisa jadi prepate statement akan selalu dibuat
berkali-kali walaupun kita menggunakan SQL yang sama
- Untuk membuat Prepare Statement, kita bisa menggunakan function (DB)Prepare(context, sql)
- Prepare Statement direpresentasikan dalah struct database/sql.stmt
- Sama seperti resource sql lainya, Stmt harus di close(), jika tidak digunakan lagi

Jika insert banyak dengan parameter, disarankan menggunakan prepare statement
*/

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "zakaria" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini comment ke " + strconv.Itoa(i)

		statement, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		lastID, err := statement.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID :", lastID)
	}
	defer statement.Close()
	defer db.Close()
}

/**
Database Transaction di Go-Lang
- Secara default semua perintah SQL yang kita kirim menggunakan GO-Lang akan otomatis commit atau istilahnya auto commit
- Artinya setiap mengirim perintah SQL ke database akan di commit secara permanent di database
- Namun kita bisa menggunakan fitur transaksi sehingga SQL yang kita kirim tidak secara otomatis di commit ke database
- Untuk memulai transaksi, kita bisa menggunakan function (DB)Begin(), dimana akan menghasilkan struct TX yang merupakan representasi dari Transaction
- Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua  function di DB ada di TX SEPERTI Exec, Query atau Prepare
- Setelah selesai proses transaksi, kita bisa gunakan function (Tx)Commit() untuk melakukan commit atau Rollback() untuk membatalkan commit atau mematalkan proses database
*/

func TestDbTransaction(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		email := "zakaria" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini comment ke " + strconv.Itoa(i)

		statement, err := tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}

		lastID, err := statement.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID :", lastID)
	}

	//err = tx.Commit()
	err = tx.Rollback() // Membatalkan eksekusi
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
