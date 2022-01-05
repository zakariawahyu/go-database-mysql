**Pengenalan Package Database**
- Bahasa pemograman GO-Lang secara default memiliki sebuah package bernama database
- Package database adalah package yang berisikan kumpulan standard interface yang menjadi standard untuk berkomunikasi ke database
- Hal ini menjadikan kode program yang kita buat untnuk mengakses jenis database apapun bisa menggunakan kode yang sama
- Yang berbeda hanya kode SQL yang perlu kita gunakan sesuai dengan database yang kita gunakan

**Database Driver**
- Sebelum kita membuat kode program menggunakan database di Go-Lang., terlebih dahulu kita wajib menambahkan driver databasenya
- Tanpa driver database, maka package database di Go-Lang tidak mengerti apapun karena hanya berisi interface saja
- https://golang.org/s/sqldrivers

**Repository Pattern**
- Dalam buku Domain-Driven Design, Eric Evans menjelaskan bahwa "repository is a mechanism for encapsulating storage ,retrieval and search behaviour, which emulates a collection of objects"
- Layer respositori adalah mekanisme untuk mengenkapsulasi storage (media penyimpanan), retrieval (pengaksesan data) dan kemampuan pencarian, dimana dia mengemaluasi kumpulan data
- Patterb Repository ini biasanya digunakan sebagai jembatan antar business logic aplikasi kita dengan semua perintah SQL ke database
- Jadi semua perintah SQL akan ditulis di Repository, sedangkan businees logic kode porgram kita hanya cukup menggunakan repository tersebut

**Alur Repository Pattern**
- Business Logic aplikasi jika akses ke database maka akan lewat repository
- Biasanya 1 tabel mempunyai 1 repository
- Repository biasanya berupa construct, kalau di golang pakai interface
- Membuat implementasi repository dengan struct
- Struct biasanya representasi dari data di database. entity/model

**Entity / Model**
- Dalam pemograman berorientasi obeject, biasanya sebuah tabel di database akan selalu dibuat representasinya sebagai class Entity atau Model, namun di Go-Lang karena tidak mengenal class, jadi jika kita akan representasikan datanya dalam bentuk struct
- Ini bisa mempermudah ketika membuat kode program, tidak perlu membuat banyak variable (lebih sederhana 1 struct berisi kolom-kolom database)
- Misal ketika kita query ke Repository, dibanding mengembalikan array alahkah baiknya Repository melakukan conversi terleih dahulu ke struct Entity/Model, sehingga kita tinggal menggunakan objectntya saja
 

