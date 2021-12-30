**Pengenalan Package Database**
- Bahasa pemograman GO-Lang secara default memiliki sebuah package bernama database
- Package database adalah package yang berisikan kumpulan standard interface yang menjadi standard untuk berkomunikasi ke database
- Hal ini menjadikan kode program yang kita buat untnuk mengakses jenis database apapun bisa menggunakan kode yang sama
- Yang berbeda hanya kode SQL yang perlu kita gunakan sesuai dengan database yang kita gunakan

**Database Driver**
- Sebelum kita membuat kode program menggunakan database di Go-Lang., terlebih dahulu kita wajib menambahkan driver databasenya
- Tanpa driver database, maka package database di Go-Lang tidak mengerti apapun karena hanya berisi interface saja
- https://golang.org/s/sqldrivers

