package main

import "fmt"

func main() {
	var namakesatu, namakedua string

	fmt.Printf("Masukan nama depan anda : ")
	fmt.Scan(&namakesatu)

	fmt.Printf("Masukan nama belakang anda : ")
	fmt.Scan(&namakedua)

	fmt.Println("jadi nama anda adalah : ", namakesatu+"\t"+namakedua)

}
