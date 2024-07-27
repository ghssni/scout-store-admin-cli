# Scout Store Admin CLI

## Deskripsi Proyek
Aplikasi Command Line Interface (CLI) untuk manajemen toko yang menjual aksesoris pramuka menggunakan Golang dan database MySQL. Aplikasi ini akan memudahkan pengguna dalam mengelola aspek-aspek penting dari sebuah toko seperti stok barang, staff, dan laporan penjualan.

## Fitur Aplikasi
1. **Tambah Produk**
   - Memungkinkan staff untuk menambah produk baru ke dalam database.
   - Staff akan diminta untuk memasukkan nama produk, harga, dan jumlah stok awal.

2. **Tambah Penjualan**
   - Mengizinkan staff untuk memasukkan produk yang terjual dan sistem akan mengurangi stok awal produk.
   - Staff dapat memasukkan produk yang terjual dan sistem akan mengurangi stok awal produk.

3. **Tambah Staff**
   - Memungkinkan pengguna untuk menambah staff baru.
   - Pengguna akan diminta untuk memasukkan nama, email, dan posisi dari staff yang akan ditambahkan.

4. **Rekap Penjualan**
   - Menampilkan laporan penjualan berdasarkan periode tertentu.
   - Pengguna dapat memilih periode waktu untuk melihat total penjualan, jumlah produk terjual, dan total pendapatan.

5. **Exit**
   - Keluar dari aplikasi. Mengakhiri sesi penggunaan aplikasi CLI.

## Teknologi yang Digunakan
- **Golang:** Untuk pengembangan aplikasi CLI.
- **MySQL:** Untuk manajemen database.

## Cara Menjalankan
- Clone repository ini
- Jalankan `go run main.go`

## Skema Database
- **Products:** `ProductID`, `ProductName`, `Price`, `Stock`
- **Staff:** `StaffID`, `StaffName`, `RoleID`
- **Roles:** `RoleID`, `RoleName`
- **Sales:** `SaleID`, `Date`, `ProductID`, `Total`,`TotalAmount`, `StaffID`

### ERD
![ERD](./images/ERD.png)

## Preview Scout Store Admin CLI
![Menu Utama](./images/)
![Tambah Produk]()
![Tambah Staff]()
![Tambah Penjualan]()
![Rekap Penjualan]()

## Pembagian Tugas
- **Driver([Ahmad Luhur Pakerti](https://github.com/00shiki)):** Menulis kode fitur tampilan menu
- **Navigator([Mahattir Onassis]()):** Meninjau kode dan memberikan feedback
- **Observer ([Ghassani Tyas](https://github.com/ghssni)):** Mencatat dan mendokumentasikan fitur
