package go_database_mysql

import (
	"context"
	"fmt"
	"testing"
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

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES('2', 'Wahyu')")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Berhasil insert data")
}

func TestExecSqlUpdate(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "UPDATE customer SET name='Nur' WHERE id='2'")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Berhasil update data")
}

func TestExecSqlDelete(t *testing.T) {
	db := GetConnection()
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "DELETE FROM customer WHERE id='2'")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Berhasil delete data")
}
