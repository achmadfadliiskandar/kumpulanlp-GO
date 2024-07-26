package main

import (
	"fmt"
	"os"
	"strings"
)

type Barang struct {
	Nama  string
	Harga int
	Stok  int
}

var barangbarang = []Barang{
	{"Beng-Beng", 2000, 50},
	{"HP Samsung", 100000, 10},
	{"Kacamata", 250000, 20},
	{"Bantal", 15000, 30},
}

func tampilkanDataBarang() {
	fmt.Println("Daftar Barang:")
	for _, barang := range barangbarang {
		fmt.Printf("%s\t Harga: %d\t Stok: %d\n", barang.Nama, barang.Harga, barang.Stok)
	}
}

func transaksi() {
	tampilkanDataBarang()

	var namaBarang string
	var jumlahBeli int

	fmt.Print("Masukkan nama barang yang akan dibeli: ")
	fmt.Scan(&namaBarang)

	namaBarang = strings.ToLower(namaBarang)

	var index int
	var found bool

	// Cari indeks barang yang akan dibeli
	for i, barang := range barangbarang {
		if strings.ToLower(barang.Nama) == namaBarang {
			index = i
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Barang tidak ditemukan")
		return
	}

	fmt.Print("Masukkan jumlah barang yang akan dibeli: ")
	fmt.Scan(&jumlahBeli)

	if jumlahBeli > barangbarang[index].Stok {
		fmt.Println("Stok barang tidak mencukupi")
		return
	}

	totalHarga := jumlahBeli * barangbarang[index].Harga
	fmt.Printf("Total Harga: %d\n", totalHarga)

	// Kurangi stok barang
	barangbarang[index].Stok -= jumlahBeli
}

func pembelian() {

	fmt.Println("Selamat berbelanja")
	for {
		var menu int
		var name, alamat string

		fmt.Print("Masukan Nama : ")
		fmt.Scan(&name)
		fmt.Print("Masukan alamat : ")
		fmt.Scan(&alamat)
		fmt.Println("Pilih barang yang akan dibeli:")
		tampilkanDataBarang()
		fmt.Print("Masukkan nomor barang: ")
		fmt.Scan(&menu)

		if menu < 1 || menu > len(barangbarang) {
			fmt.Println("Nomor barang tidak valid")
			continue
		}

		var jumlah int
		fmt.Print("Masukkan jumlah barang yang akan dibeli: ")
		fmt.Scan(&jumlah)

		if jumlah > barangbarang[menu-1].Stok {
			fmt.Println("Stok tidak mencukupi")
			continue
		}
		fmt.Println("Barang ini akan diantarkan kepada", name)
		fmt.Println(alamat)

		totalHarga := jumlah * barangbarang[menu-1].Harga
		fmt.Printf("Total Harga: %d\n", totalHarga)

		// Kurangi stok barang
		barangbarang[menu-1].Stok -= jumlah

		fmt.Print("Ingin beli barang lagi? (y/n): ")
		var jawaban string
		fmt.Scan(&jawaban)

		if jawaban != "y" {
			break
		}
	}
}

func menu() {
	var menu int

	fmt.Println("SELAMAT DATANG DI TOKO ONLINE 13")
	fmt.Println("================================")
	fmt.Println("1. Data Barang")
	fmt.Println("2. Pembelian")
	fmt.Println("3. Keluar")

	fmt.Printf("Pilih angka menu sesuai dengan menu di atas: ")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		tampilkanDataBarang()
	case 2:
		transaksi()
	case 3:
		fmt.Println("Terima kasih sudah mencoba aplikasi kami")
		os.Exit(0)
	default:
		fmt.Println("Maaf, menu yang Anda input tidak termasuk ya")
	}
}

func main() {
	for {
		menu()
	}
}
