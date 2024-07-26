package main

import (
	"fmt"
	"os"
)

// fungsi untuk uuseryang memilih nomor 1
func barang() {
	// fmt.Println("anda masuk menu 1")
	var barangbarang = [4]string{"beng-beng : ", "hpsamsung", "kacamata", "bantal"}
	var hargabarang = [4]int{2000, 100000, 250000, 15000}
	fmt.Print(barangbarang[0], "\t", hargabarang[0])
	fmt.Print("\n", barangbarang[1], "\t", hargabarang[1])
	fmt.Print("\n", barangbarang[2], "\t", hargabarang[2])
	fmt.Println("\n", barangbarang[3], "\t", hargabarang[3])
}

func Transaksi() {
	fmt.Println("anda masuk menu 2")
}
func Pembelian() {
	fmt.Println("selamat belanja")
}

func main() {
	// tentuin variabelnya dulu
	var langkah string = "1-2-3"

	var menu int

	// menu utama aplikasi
	fmt.Println("SELAMAT DATANG DI TOKO ONLINE 13")
	fmt.Println("================================")
	fmt.Println("Langkah Belanja Menu :", langkah)

	fmt.Println("1.Data Barang")
	fmt.Println("2.Transaksi")
	fmt.Println("3.Keluar")
	fmt.Println("4.Pembelian")

	//masukin inputan brayy

	fmt.Printf("Pilih Angka Menu sesuai dengan menu diatas : ")
	fmt.Scan(&menu)
	// tes tangkap input duluu
	//fmt.Println(menu) bisa

	if menu == 1 {
		barang()
	} else if menu == 2 {
		Transaksi()
	} else if menu == 3 {
		fmt.Println("terima kasih sudah mencoba aplikasi kami")
		os.Exit(0)
	} else if menu == 4 {
		Pembelian()
	} else {
		fmt.Println("maaf menu yang kamu input tidak termasuk ya")
	}
}
