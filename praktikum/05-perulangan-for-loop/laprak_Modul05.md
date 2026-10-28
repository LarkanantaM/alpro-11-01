# <h1 align="center">Laporan Praktikum Modul 05 - Variabel dan Operator</h1>
<p align="center">LINGGA ARKANANTA MAHARDIKA - 109092600012</p>

## Dasar Teori

### A. Perulangan For di Go
Perulangan for menjalankan perintah yang sama secara berulang tanpa menulis kode berkali-kali, dan merupakan satu-satunya perulangan di Go. Bentuk tiga komponennya adalah for inisialisasi; kondisi; post statement { ... }: inisialisasi dijalankan sekali di awal, kondisi dicek sebelum tiap putaran, dan post statement dijalankan setelah tiap putaran

### B. For Hanya Kondisi dan Perulangan Bersarang
Karena Go tidak punya yang namanya while, perannya digantikan oleh for yang hanya berisi kondisi. for juga bisa diletakkan di dalam for lain (perulangan bersarang) untuk memproses data yang tiap itemnya punya beberapa nilai


## Guided

### 1. tabungan.go

```go
package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a, &b, &n)
	total := 0

	hariIni := a

	for i := 1; i <= n; i++ {
		total = total + hariIni
		hariIni = hariIni + b
	}
	fmt.Print(total)
}

```
#### Deskripsi
Menghitung tabungan awal dan melakukan perulangan pada uang tambahan. contoh: uang tambahannya 500, maka akan dilakukan perulangan sampai n

### 2. terbesar.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var angka, max int
	fmt.Scan(&angka)
	max = angka

	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Print(max)
}

```
#### Deskripsi
Program mencari bilangan bulat terbesar dari n yang diimput melalui for

### 2. bioskop.go

```go
package main

import "fmt"

func main() {
	var t int
	fmt.Print("Input t(jumlah transaksi yang akan dilakukan): ")
	fmt.Scan(&t)

	totalP := 0

	for i := 1; i <= t; i++ {
		var film, hari string
		var tTiket int

		fmt.Printf("Input transaksi ke-%d (film, hari, tiket)", i)
		fmt.Scan(&film, &hari, &tTiket)

		hargaS := 0

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if film == "Reguler" {
				hargaS = 40000
			} else if film == "IMAX" {
				hargaS = 50000
			} else if film == "3D" {
				hargaS = 60000
			}
		}

		if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if film == "Reguler" {
				hargaS = 45000
			} else if film == "IMAX" {
				hargaS = 55000
			} else if film == "3D" {
				hargaS = 65000
			}
		}
		totalH := hargaS * tTiket

		if tTiket >= 5 {
			diskon := (totalH * 10) / 100
			totalH = totalH - diskon
		}

		fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalH)

		totalP = totalP + totalH
	}

	fmt.Println("output total pendapatan:", totalP)
}

```
#### Deskripsi
Program yang membaca transaksi, menghitung harga sesuai tiket, jenis film dan hari pemutarannya.


## Unguided

### 1. prima.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ { 
		if n%i == 0 {
			fmt.Printf("%d ", i)
		} 
	}
	fmt.Println()

	if n%2 == 0 {
		fmt.Println("Tidak")
	} else {
		fmt.Println("Iya")
	}

	
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](unguided/hari/output.png)


#### Deskripsi
Program yang menampilkan seluruh faktor dari n. Lalu memberi nilai "iya" jika bilangan yang diinput bilangan prima jika yang diinput bukan bilangan prima akan diberi nilai "Tidak".

### 2. ecct.go

```go
package main

import "fmt"

func main() {
	var a float64
	var lulus int
	n := 6
	b := 2

	for u := 1; u <= b; u++ {
		var hasil float64

		for i := 1; i <= n; i++ {
			hasil = hasil + a
			fmt.Scan(&a)

			if a >= 5 {
				break
			}
		}

		hasilA := hasil / 6

		if hasilA >= 3.33 {
			fmt.Println("lulus")
			lulus++
		} else {
			fmt.Println("Tidak Lulus")
		}

	}
	fmt.Println(lulus)
}

```

##### Output
![Screenshot Output Unguided](unguided/konversi/output.png)

#### Deskripsi
Program yang menentukan mahasiswa lulus atau tidak dengan minimal skor 3.33. Program juga menghitung seberapa banyak mahasiswa yang lulus.


## Kesimpulan
Praktikum ini melatih penggunaan perulangan for untuk mengulang proses sesuai kebutuhan, seperti menjumlahkan data, mencari nilai terbesar, dan menghitung transaksi. Perulangan juga bisa digabung dengan percabangan dan perulangan bersarang untuk mengolah data yang lebih komplek

## Referensi
1. Google. ([2026]). *[The Go Programming Language Documentation]*. [Mountain View]: [Google]. Diakses pada [27 september 2026] melalui [[tautan/DOI](https://go.dev/doc/)]
2. [Donovan, A. A. A., & Kernighan, B. W.]. ([2015]). *[The Go Programming Language]*. [Boston]: [Donovan, A. A. A., & Kernighan, B. W.]. Diakses pada [27 september 2026] melalui [[tautan/DOI](https://www.gopl.io/)]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
