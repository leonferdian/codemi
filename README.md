# codemi
Codemi Online Test

Aplikasi tes online ini merupakan aplikasi berbasis bahasa pemrograman go, output interface aplikasi menggunakan interaktif input di terminal.
Untuk menjalankan aplikasi clone repository ke go path, untuk mengakses informasi gopath inputkan command berikut "go env GOPATH" di terminal.

Aplikasi ini merupakan program sederhana yang digunakan untuk merekam data loker.
Setelah aplikasi berhasil dicloning ikuti beberapa langkah berikut untuk menjalankan aplikasi:

1. Buka terminal yang dapat digunakan untuk mengakses file path aplikasi
2. Untuk repository ini aplikasi terletak di folder codemi/interactive-loker/main.go
3. Setelah itu compile aplikasi dengan menjalankan perintah "go run main.go"
4. Saat program berjalan dengan baik maka akan muncul judul aplikasi kemudian input
5. Input pertama "init [jumlah loker]" akan muncul saat pertama program dijalankan, isi input dengan jumlah yang sesuai
6. Setelah menentukan jumlah loker maka user mengisi data loker dengan format "input [Tipe Identitas] [No Identitas]"
7. Setelah memasukkan input data dengan jumlah maksimal yang ditentukan user dapat mengecek data dengan perintah "status"
8. Untuk mencari data dengan nomer identitas gunakan perintah "find [No Identitas]"
9. Untuk melihat daftar dengan tipe identitas tertentu gunakan perintah "search [Tipe Identitas]"
10. Untuk menghapus data dengan nomor urutan yang spesifik maka gunakan perintah "leave [No urut data]"
11. Untuk keluar dari aplikasi gunakan perintah "exit"
12. Untuk informasi daftar perintah aplikasi gunakan perintah "help"

Semua perintah yang ada di dalam aplikasi sifatnya `insensitive`, tidak akan berpengaruh terhadap lowercase dan uppercase.
