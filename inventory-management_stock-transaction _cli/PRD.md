Berikut adalah \*\*Product Requirement Document (PRD)\*\* atau spesifikasi detail kebutuhan untuk \*\*Project 1: Inventory Management & Stock Transaction CLI\*\*, disusun khusus untuk Anda jadikan panduan teknis saat membangun kodenya di kolom chat ini.

\---

\# Product Requirement Document (PRD)

\## Project 1: Inventory Management & Stock Transaction CLI (Golang)

\### 1. Deskripsi Proyek

Aplikasi berbasis CLI (Command Line Interface) yang dirancang untuk mengelola inventaris barang, memproses mutasi stok masuk (`IN`) dan keluar (`OUT`) secara aman menggunakan pointer dan custom error, serta menghasilkan laporan valuasi aset gudang secara otomatis.

\---

\### 2. Struktur Data (Struct & Collection)

Aplikasi harus mengimplementasikan tipe data terstruktur berikut:

\#### A. Struct `Item`

Merepresentasikan data master barang di gudang.

* `ID` (string): Kode unik barang (contoh: `"BRG01"`, `"BRG02"`).
* `Name` (string): Nama barang.
* `Stock` (int): Jumlah ketersediaan stok saat ini.
* `Price` (int / float64): Harga satuan barang.

\#### B. Struct `Transaction`

Merepresentasikan riwayat mutasi stok barang.

* `ItemID` (string): Merujuk pada ID barang yang ditransaksikan.
* `Type` (string): Jenis transaksi, bernilai `"IN"` (stok masuk) atau `"OUT"` (stok keluar).
* `Quantity` (int): Jumlah barang yang dimasukkan atau dikeluarkan.
* `Timestamp` (string / time.Time): Waktu transaksi dilakukan.

\#### C. Penyimpanan Data (Storage)

* Gunakan \*\*`map[string]Item`\*\* sebagai struktur data utama untuk menyimpan inventaris agar pencarian berdasarkan `ID` menjadi efisien ($O(1)$).
* Sediakan data awal (\*seed data\*) minimal \*\*5 barang\*\* di dalam map saat program pertama kali dijalankan.

\---

\### 3. Modul & Fungsi Utama (Modular Programming)

Kode tidak boleh menumpuk di dalam fungsi `main()`. Buat fungsi terpisah dengan spesifikasi tanda tangan (\*signature\*) sebagai berikut:

1. \*\*`AddStock(items map[string]Item, id string, qty int)`\*\*
* \*\*Deskripsi:\*\* Menambahkan jumlah stok barang berdasarkan `id` yang diberikan.
* \*\*Aturan:\*\* Harus menggunakan \*\*pointer\*\* pada manipulasi struct atau map agar perubahan nilai stok langsung tersimpan pada data aslinya.
* \*\*Validasi:\*\* Jika `id` barang tidak ditemukan di dalam map, tampilkan pesan error yang informatif.


1. \*\*`SellItem(items map[string]Item, id string, qty int) error`\*\*
* \*\*Deskripsi:\*\* Mengurangi stok barang karena penjualan/keluar (`OUT`).
* \*\*Validasi & Error Handling:\*\*
* Jika `id` barang tidak ditemukan, kembalikan error.
* Jika jumlah stok saat ini \*\*kurang\*\* dari `qty` yang ingin dikeluarkan, wajib mengembalikan custom error menggunakan \*\*`errors.New("stok tidak mencukupi")`\*\*.




1. \*\*`RecordTransaction(history \*[]Transaction, itemID string, tType string, qty int)`\*\* \*(Opsional/Anjuran)\*
* Mencatat setiap transaksi sukses (`IN`/`OUT`) ke dalam sebuah \*slice of Transaction\*.


1. \*\*`PrintWarehouseValuation(items map[string]Item)`\*\*
* Menghitung dan mencetak laporan akhir rekapitulasi gudang.



\---

\### 4. Alur Program & Fitur CLI (User Experience)

1. \*\*Menu Utama:\*\* Saat program berjalan, tampilkan daftar inventaris awal beserta informasi ID, Nama, Stok, dan Harga.
1. \*\*Interaksi Menu / Perintah:\*\* Berikan pilihan interaktif kepada pengguna menggunakan `bufio.NewReader` dan `strconv`:
* Opsi 1: Tambah Stok Masuk (`IN`) -> Memanggil `AddStock`.
* Opsi 2: Kurangi/Jual Stok (`OUT`) -> Memanggil `SellItem` (tangkap error jika stok kurang).
* Opsi 3: Keluar dan Cetak Laporan Akhir.


1. \*\*Laporan Akhir (Valuasi Aset Gudang):\*\*
* Ketika pengguna memilih keluar, program wajib melakukan perulangan `for ... range` pada map inventaris.
* Menghitung subtotal per barang: $\text{Stok Akhir} \times \text{Harga}$.
* Menjumlahkan seluruh subtotal menjadi \*\*Total Valuasi Aset Gudang\*\*.
* Menampilkan ringkasan transaksi yang telah dilakukan.



\---

\### 5. Kriteria Keberhasilan (Acceptance Criteria)

* [ ] Kode terstruktur secara modular (terpisah dari `main`).
* [ ] Penggunaan `map[string]Item` berhasil mengelola minimal 5 data awal.
* [ ] Fungsi `AddStock` sukses mengubah nilai asli menggunakan mekanisme referensi/pointer.
* [ ] Fungsi `SellItem` berhasil memicu dan menangani custom error (`errors.New`) ketika stok melebihi batas.
* [ ] Laporan akhir valuasi aset tercetak akurat menggunakan perulangan `range`.

\---

Silakan gunakan PRD di atas sebagai acuan. Anda bisa mulai menulis kodenya di sini, dan saya siap membantu mereview atau memperbaiki jika ada bagian yang error!
