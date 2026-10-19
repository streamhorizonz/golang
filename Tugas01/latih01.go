package main

import (
	"fmt"
)

func main() {
	fmt.Println("Wello")

	// Deklarasi Variable
	// 1. Formal / Eksplisit
	var namaVariable string = "Nama_Saya"
	var namaInteger int
	var namaBoolean bool
	// 2. Infferensi/Implisit
	namaVariabel2 := "Nama_Saya"
	namaInteger2 := 10
	namaBoolean2 := true
	var (
		namaVariabel3 string = "Nama_Saya"
		namaInteger3  int    = 10
		namaBoolean3  bool   = true
	)
	fmt.Println(namaVariable)

	fmt.Println(namaInteger)
	fmt.Println(namaInteger2)
	fmt.Println(namaBoolean2)
	fmt.Println(namaVariabel3)
	fmt.Println(namaInteger3)
	fmt.Println(namaBoolean3)
}
